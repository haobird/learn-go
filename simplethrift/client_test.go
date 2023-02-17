package main

import (
	"context"
	"fmt"
	"samplethrift/gen-go/Sample"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
)

var ctx = context.Background()

func GetClient() *Sample.GreeterClient {
	addr := ":9092"
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
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

	client := Sample.NewGreeterClient(thrift.NewTStandardClient(iprot, oprot))
	return client
}

//SayHello
func TestSayHello(t *testing.T) {
	client := GetClient()

	user := &Sample.User{}
	user.Name = ""
	user.Address = "address"

	rep, err := client.SayHello(ctx, user)
	if err != nil {
		t.Errorf("thrift err: %v\n", err)
	} else {
		t.Logf("Recevied: %v\n", rep)
	}
}
