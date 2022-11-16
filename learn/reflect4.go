package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type Msg struct {
	Subject string
	Data    []byte
}

type TradeEventPayload struct {
	ID       uint      `json:"id"`
	Symbol   string    `json:"symbol"`
	Entry    float64   `json:"entry"`
	Exit     float64   `json:"exit"`
	Quantity float64   `json:"quantity"`
	Time     time.Time `json:"time"`
}

var emptyMsgType = reflect.TypeOf(&Msg{})

// Dissect the cb Handler's signature
func argInfo(cb any) (reflect.Type, int) {
	cbType := reflect.TypeOf(cb)
	if cbType.Kind() != reflect.Func {
		panic("nats: Handler needs to be a func")
	}
	numArgs := cbType.NumIn()
	if numArgs == 0 {
		return nil, numArgs
	}
	return cbType.In(numArgs - 1), numArgs
}

// Decode
func DecodeDefault(data []byte, vPtr interface{}) error {
	// Figure out what it's pointing to...
	sData := *(*string)(unsafe.Pointer(&data))
	switch arg := vPtr.(type) {
	case *string:
		*arg = sData
		return nil
	case *[]byte:
		*arg = data
		return nil
	case *int:
		n, err := strconv.ParseInt(sData, 10, 64)
		if err != nil {
			return err
		}
		*arg = int(n)
		return nil
	case *int32:
		n, err := strconv.ParseInt(sData, 10, 64)
		if err != nil {
			return err
		}
		*arg = int32(n)
		return nil
	case *int64:
		n, err := strconv.ParseInt(sData, 10, 64)
		if err != nil {
			return err
		}
		*arg = int64(n)
		return nil
	case *float32:
		n, err := strconv.ParseFloat(sData, 32)
		if err != nil {
			return err
		}
		*arg = float32(n)
		return nil
	case *float64:
		n, err := strconv.ParseFloat(sData, 64)
		if err != nil {
			return err
		}
		*arg = float64(n)
		return nil
	case *bool:
		b, err := strconv.ParseBool(sData)
		if err != nil {
			return err
		}
		*arg = b
		return nil
	default:
		vt := reflect.TypeOf(arg).Elem()
		return fmt.Errorf("nats: Default Encoder can't decode to type %s", vt)
	}
}

func DecodeJson(data []byte, vPtr interface{}) (err error) {
	switch arg := vPtr.(type) {
	case *string:
		// If they want a string and it is a JSON string, strip quotes
		// This allows someone to send a struct but receive as a plain string
		// This cast should be efficient for Go 1.3 and beyond.
		str := string(data)
		if strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
			*arg = str[1 : len(str)-1]
		} else {
			*arg = str
		}
	case *[]byte:
		*arg = data
	default:
		err = json.Unmarshal(data, arg)
	}
	return
}

func MockFunction(f any) error {
	argType, numArgs := argInfo(f)
	fmt.Printf("argType:%v, numArgs:%v\n", argType, numArgs)

	cbValue := reflect.ValueOf(f)

	var oV []reflect.Value
	m := 20
	oV = []reflect.Value{reflect.ValueOf(m)}

	out := cbValue.Call(oV)
	fmt.Println(out)
	return nil
}

func Subscribe(ch chan Msg, cb any) error {
	if cb == nil {
		return errors.New("nats: Handler required for EncodedConn Subscription")
	}
	argType, numArgs := argInfo(cb)
	if argType == nil {
		return errors.New("nats: Handler requires at least one argument")
	}

	cbValue := reflect.ValueOf(cb)
	wantsRaw := (argType == emptyMsgType)
	fmt.Println("argType", argType, "wantsRaw", wantsRaw, "argType.Kind", argType.Kind())

	handler := func(m *Msg) error {
		var oV []reflect.Value
		if wantsRaw {
			oV = []reflect.Value{reflect.ValueOf(m)}
		} else {
			var oPtr reflect.Value
			if argType.Kind() != reflect.Ptr {
				oPtr = reflect.New(argType)
			} else {
				fmt.Println("argType.Elem", argType.Elem())
				oPtr = reflect.New(argType.Elem())
			}

			if err := DecodeJson(m.Data, oPtr.Interface()); err != nil {
				return errors.New("nats: Got an error trying to unmarshal: " + err.Error())
			}
			if argType.Kind() != reflect.Ptr {
				oPtr = reflect.Indirect(oPtr)
			}

			// Callback Arity
			switch numArgs {
			case 1:
				oV = []reflect.Value{oPtr}
			case 2:
				subV := reflect.ValueOf(m.Subject)
				oV = []reflect.Value{subV, oPtr}
			}
		}
		cbValue.Call(oV)
		return nil
	}

	val := <-ch
	return handler(&val)

}

var (
	ch = make(chan Msg, 10)
)

func main() {
	// MockFunction(func(i int) error {
	// 	fmt.Println("打印", i)
	// 	if i > 50 {
	// 		return fmt.Errorf("error: %d", i)
	// 	}
	// 	return nil
	// })

	ele := TradeEventPayload{
		ID:     123,
		Symbol: "ddd",
	}
	buf, _ := json.Marshal(ele)
	ch <- Msg{"x", buf}
	err := Subscribe(ch, func(p TradeEventPayload) {
		fmt.Println(p)
	})
	fmt.Println(err)

	ch <- Msg{"y", []byte("dd")}
	err = Subscribe(ch, func(p *Msg) {
		fmt.Println(p)
	})
	fmt.Println(err)

	ch <- Msg{"z", buf}
	err = Subscribe(ch, func(p string) {
		fmt.Println(p)
	})
	fmt.Println(err)

	ch <- Msg{"a", buf}
	err = Subscribe(ch, func(p *TradeEventPayload) {
		fmt.Println(p)
	})
	fmt.Println(err)
}
