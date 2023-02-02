package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"world/thrift/gen-go/hellothrift"

	"github.com/apache/thrift/lib/go/thrift"
)

const TRACEID = "__TRACEID__"

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

//定义服务
type Greeter struct {
}

//实现IDL里定义的接口
//SayHello
func (this *Greeter) SayHello(ctx context.Context, req *hellothrift.HelloReq) (resp *hellothrift.HelloResp, err error) {
	traceId := req.Header.GetTraceID()
	respHeader := hellothrift.NewRespHeader()
	respHeader.TraceID = traceId
	fmt.Println("traceId", traceId)
	panic("test")
	return &hellothrift.HelloResp{
		Header: respHeader,
		Data:   map[string]string{"User": req.GetName()},
	}, nil
}

func SimpleProcessorLoggingMiddleware(name string, next thrift.TProcessorFunction) thrift.TProcessorFunction {
	return thrift.WrappedTProcessorFunction{
		Wrapped: func(ctx context.Context, seqId int32, in, out thrift.TProtocol) (bool, thrift.TException) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
				}
			}()
			traceId := ctx.Value(TRACEID)
			log.Printf("Before: %q, seqId: %d, traceId:%s", name, seqId, traceId)
			success, err := next.Process(ctx, seqId, in, out)
			log.Printf("After: %q", name)
			log.Printf("Success: %v", success)
			if err != nil {
				log.Printf("Error: %v", err)
			}
			return success, err
		},
	}
}

func main() {
	//命令行参数
	flag.Usage = Usage
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	s := flag.Bool("s", false, "client or server")
	flag.Parse()

	path := "/data/hhh"

	if *s {
		//handler
		handler := &Greeter{}
		//processor
		processor := hellothrift.NewHelloServiceProcessor(handler)
		processor1 := thrift.WrapProcessor(processor, SimpleProcessorLoggingMiddleware)
		err := apusServer(path, *addr, processor1)
		fmt.Println("server", err)
	}
	fmt.Println("client")
	c, err := apusErrClient(path)
	if err != nil {
		fmt.Println("client", err)
	}

	// 自动创建相关配置
	client := hellothrift.NewHelloServiceClient(c)

	Hello(client, "myname")
	Hello(client, "worlod")

}

func Hello(client *hellothrift.HelloServiceClient, name string) {
	traceId := time.Now().String()
	ctx := context.Background()
	ctx = context.WithValue(ctx, TRACEID, traceId)
	req := hellothrift.NewHelloReq()
	reqHeader := hellothrift.NewReqHeader()
	reqHeader.TraceID = traceId
	req.Header = reqHeader
	req.Name = "bulang"
	fmt.Println("traceId", ctx.Value(TRACEID))
	resp, err := client.SayHello(ctx, req)
	fmt.Printf("%+v, %v", resp, err)
}

// 客户端自动创建流程
func apusClient(path string) (*thrift.TStandardClient, error) {
	// 获取zk注册的地址列表流程
	addrs := []string{"127.0.0.1:9090"}
	fmt.Println(addrs)

	var transport thrift.TTransport
	transport, err := thrift.NewTSocket(addrs[0])
	if err != nil {
		return nil, err
	}

	//protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//no buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := thrift.NewTStandardClient(iprot, oprot)
	return client, err

	// 创建对应数据类型
	// v := reflect.ValueOf(s)
	// if v.Kind() == reflect.Ptr && !v.IsNil() {
	// 	v = v.Elem()
	// }

	// fmt.Println(v.Kind())

	// // 建立

	// return nil
}

// 服务端自动创建流程
func apusServer(path string, addr string, processor thrift.TProcessor) error {
	// 获取本机IP，
	// addr := "127.0.0.1:9091"
	fmt.Println(path)

	// 服务退出，则取消注册
	defer func(addr string) {
		fmt.Println("unregister", addr)
	}(addr)

	// 写入zk

	// 建立服务流程
	//protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()
	// transportFactory = thrift.NewTBufferedTransportFactory(8192)
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	//transport,no secure
	var transport thrift.TServerTransport
	transport, err := thrift.NewTServerSocket(addr)

	//start tcp server
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()
	return err
}

// 注册服务流程
func register(port int, path string) {
	// 获取本机IP

	// 写入zk
}

func apusErrClient(path string) (*thrift.TStandardClient, error) {
	// 获取zk注册的地址列表流程
	addrs := []string{"127.0.0.1:9090"}
	fmt.Println(addrs)

	var transport thrift.TTransport
	transport, err := thrift.NewTSocket(addrs[0])
	if err != nil {
		return nil, err
	}

	//protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//no buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()
	// transportFactory = thrift.NewTBufferedTransportFactory(8192)
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := thrift.NewTStandardClient(iprot, oprot)
	return client, err

	// 创建对应数据类型
	// v := reflect.ValueOf(s)
	// if v.Kind() == reflect.Ptr && !v.IsNil() {
	// 	v = v.Elem()
	// }

	// fmt.Println(v.Kind())

	// // 建立

	// return nil
}
