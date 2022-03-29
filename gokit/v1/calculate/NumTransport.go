package calculate

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint EndPointServer) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder), //程序中的全部报错都会走这里面
	}
	r := mux.NewRouter()
	calculate := r.PathPrefix("/calculate").Subrouter()
	calculate.Handle("/{type}/{a}/{b}", httptransport.NewServer(
		endpoint.CalculateEndPoint,
		decodeHTTPCalculateRequest, //解析请求值
		encodeHTTPGenericResponse,  //返回值
		options...,
	))
	// 健康检查
	r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(`{"status":"ok"}`))
	})
	return r
}

func decodeHTTPCalculateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	types := vars["type"]
	a, _ := strconv.ParseInt(vars["a"], 10, 0)
	b, _ := strconv.ParseInt(vars["b"], 10, 0)
	in := CalculateIn{A: int(a), B: int(b), Type: types}
	return in, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	fmt.Println("errorEncoder", err.Error())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

type errorWrapper struct {
	Error string `json:"errors"`
}
