package discover

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func arithmeticFactory(_ context.Context, method, path string) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, err error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}
		tgt, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path
		var (
			enc httptransport.EncodeRequestFunc
			dec httptransport.DecodeResponseFunc
		)
		fmt.Println(tgt)
		enc, dec = encodeArithmeticRequest, decodeArithmeticReponse
		return httptransport.NewClient(method, tgt, enc, dec).Endpoint(), nil, nil
	}
}

func encodeArithmeticRequest(_ context.Context, req *http.Request, request interface{}) error {
	arithReq := request.(CalculateIn)
	p := "/" + arithReq.Type + "/" + strconv.Itoa(arithReq.A) + "/" + strconv.Itoa(arithReq.B)
	fmt.Println(p)
	req.URL.Path += p
	return nil
}

func decodeArithmeticReponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response CalculateResponse
	var s map[string]interface{}
	if respCode := resp.StatusCode; respCode >= 400 {
		fmt.Println(respCode)
		fmt.Println(resp.Body)
		if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
			return nil, err
		}
		return nil, errors.New(s["error"].(string) + "\n")
	}
	respCode := resp.StatusCode
	fmt.Println(respCode)
	fmt.Println(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	fmt.Println(response)
	return response, nil
}
