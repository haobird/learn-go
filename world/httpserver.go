package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
	"go.uber.org/zap"
)

const contentType = "application/json"

var errUnsupportedContentType = errors.New("unsupported content type")

func main() {
	logger, _ := zap.NewProduction()
	port := ":8001"
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(Final),
	}

	mux := bone.New()

	mux.Get("/hello", kithttp.NewServer(
		notifyEndpoint(),
		DecodeJSONRequest,
		kithttp.EncodeJSONResponse,
		opts...,
	))

	mux.Post("/search", kithttp.NewServer(
		notifyEndpoint(),
		DecodeJSONRequest,
		kithttp.EncodeJSONResponse,
		opts...,
	))

	err := http.ListenAndServe(port, LoggingHTTPMiddleware(mux, logger))
	fmt.Println(err)
}

func LoggingHTTPMiddleware(next http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()

		// 创建 req副本保存请求信息
		body, _ := io.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// 创建一个 recorder 用来保存响应信息
		recorder := httptest.NewRecorder()

		// 执行下一个
		next.ServeHTTP(recorder, r)

		// 从 recorder 中提取记录下来的 Response Header，设置为 ResponseWriter 的 Header
		for key, value := range recorder.HeaderMap {
			for _, val := range value {
				w.Header().Set(key, val)
			}
		}

		// 提取 recorder 中记录的状态码，写入到 ResponseWriter 中
		w.WriteHeader(recorder.Code)

		// var contentLength int
		if recorder.Body != nil {
			// 将 recorder 记录的 Response Body 写入到 ResponseWriter 中，客户端收到响应报文体
			w.Write(recorder.Body.Bytes())

			// 计算 Response Body 的大小（即 Content-Length）
			// contentLength = recorder.Body.Len()
		}

		// 打印日志
		logger.Info("HTTP",
			zap.String("URI", r.RequestURI),
			zap.Int("STATUS", recorder.Code),
			zap.Int64("RequestLength", r.ContentLength),
			zap.Any("RequestHeader", r.Header),
			zap.String("RequestBODY", string(body)),
			zap.Int("ResponseLength", recorder.Body.Len()),
			zap.Any("ResponseHeader", recorder.HeaderMap),
			zap.String("ResponseBody", string(recorder.Body.Bytes())),
			zap.String("COST", time.Since(begin).String()),
		)

	})
}

func Final(ctx context.Context, code int, r *http.Request) {
	fmt.Printf("code:%d,uri:%s\n", code, r.RequestURI)
}

func notifyEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println("notifyEndpoint", request)

		return nil, err
	}
}

func DecodeJSONRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req map[string]interface{}
	if r.ContentLength == 0 {
		return nil, nil
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// func EncodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	if headerer, ok := response.(Headerer); ok {
// 		for k, values := range headerer.Headers() {
// 			for _, v := range values {
// 				w.Header().Add(k, v)
// 			}
// 		}
// 	}
// 	code := http.StatusOK
// 	if sc, ok := response.(StatusCoder); ok {
// 		code = sc.StatusCode()
// 	}
// 	w.WriteHeader(code)
// 	if code == http.StatusNoContent {
// 		return nil
// 	}
// 	return json.NewEncoder(w).Encode(response)
// }
