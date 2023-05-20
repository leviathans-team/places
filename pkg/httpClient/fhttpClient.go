package httpClient

//import (
//	"encoding/json"
//	"fmt"
//	"github.com/google/uuid"
//	"github.com/valyala/fasthttp"
//	"golang-pkg/config"
//	"golang-pkg/pkg/secure"
//	"time"
//)
//
//type Client struct {
//	cfg             *config.Config
//	timeout         time.Duration
//	requestSerialId string
//}
//
//func NewClient(cfg *config.Config) *Client {
//	return &Client{
//		cfg:             cfg,
//		timeout:         time.Second * 15,
//		requestSerialId: uuid.New().String(),
//	}
//}
//
//func (c *Client) Request(method string, url string, body interface{},
//	queryParams map[string]string, headers map[string]string) (responseBody []byte, statusCode int, err error) {
//	req := fasthttp.AcquireRequest()
//	req.Header.SetMethod(method)
//	for k, v := range headers {
//		req.Header.Set(k, v)
//	}
//	var queryCount int
//	for k, v := range queryParams {
//		switch queryCount {
//		case 0:
//			url += fmt.Sprintf("?%s=%s", k, v)
//		default:
//			url += fmt.Sprintf("&%s=%s", k, v)
//		}
//		queryCount++
//	}
//	req.Header.Set("requestSerialID", c.requestSerialId)
//
//	bodyData, err := json.Marshal(body)
//	if err != nil {
//		return nil, 0, err
//	}
//	req.SetBody(bodyData)
//	req.SetRequestURI(url)
//	res := fasthttp.AcquireResponse()
//	defer fasthttp.ReleaseRequest(req)
//	defer fasthttp.ReleaseResponse(res)
//
//	client := &fasthttp.Client{}
//
//	if err = client.DoTimeout(req, res, c.timeout); err != nil {
//		return
//	}
//
//	statusCode = res.StatusCode()
//
//	responseBody = make([]byte, len(res.Body()))
//	copy(responseBody, res.Body())
//	req.SetConnectionClose()
//	return responseBody, statusCode, nil
//}
//
//func (c *Client) MakeHeaders(body, privateKey, publicKey string) map[string]string {
//	headers := make(map[string]string)
//	headers["Content-Type"] = "application/json"
//	headers["Accept"] = "application/json"
//	signature := secure.CalcSignature(privateKey, fmt.Sprintf("%v", body))
//	headers["X-Signature"] = signature
//	headers["X-Public-Key"] = publicKey
//	return headers
//}
