package discover

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

type CalculateIn struct {
	A    int    `json:"a"`
	B    int    `json:"b"`
	Type string `json:"type"`
}

// ArithmeticResponse define response struct
type CalculateResponse struct {
	Res   int   `json:"res"`
	Error error `json:"error"`
}

func MakeHttpHandler(endpoint endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/calculate/{type}/{a}/{b}").Handler(httptransport.NewServer(
		endpoint,
		decodeDiscoverRequest,
		encodeDiscoverResponse,
	))
	return r
}

func encodeDiscoverResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeDiscoverRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	types := vars["type"]
	a, _ := strconv.ParseInt(vars["a"], 10, 0)
	b, _ := strconv.ParseInt(vars["b"], 10, 0)
	in := CalculateIn{A: int(a), B: int(b), Type: types}
	return in, nil
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	fmt.Println("errorEncoder", err.Error())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

type errorWrapper struct {
	Error string `json:"errors"`
}
