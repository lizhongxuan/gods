package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	XUserToken = "x-wps-weboffice-token"
)

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type GetUserInfoBatchInput struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

type ExternalProvider struct {
	http *http.Client
}

func (cli *ExternalProvider) PostFileOnline(ctx context.Context, fileid string, ids []string) error {
	in := &GetUserInfoBatchInput{
		Ids: ids,
	}
	if err := cli.post(ctx, fileid, "/v1/3rd/file/online", in, nil); err != nil {
		return err
	}
	return nil
}


func (cli *ExternalProvider) newRequest(ctx context.Context, method string, fileid string, path string, body io.Reader) *http.Request {
	params := getParams(ctx)
	gateway := getGateway(ctx)
	userAgent := getUserAgent(ctx)

	url := gateway + path + "?" + params.Encode()
	req, _ := http.NewRequest(method, url, body)
	SetHeaderToken(req.Header, GetContextToken(ctx))

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("x-weboffice-file-id", fileid)

	return req.WithContext(ctx)
}


func (cli *ExternalProvider) post(ctx context.Context, fileid, path string, args interface{}, reply interface{}) error {
	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(args)

	req := cli.newRequest(ctx, http.MethodPost, fileid, path, body)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, respErr := cli.http.Do(req)

	err := unmarshalResponse(req, resp, respErr, reply)
	return err
}

func (cli *ExternalProvider) get(ctx context.Context, fileid, path string, reply interface{}) error {
	req := cli.newRequest(ctx, http.MethodGet, fileid, path, nil)
	resp, respErr := cli.http.Do(req)

	err := unmarshalResponse(req, resp, respErr, reply)
	return err
}

func (cli *ExternalProvider) put(ctx context.Context, fileid, path string, args interface{}, reply interface{}) error {
	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(args)

	req := cli.newRequest(ctx, http.MethodPut, fileid, path, body)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, respErr := cli.http.Do(req)

	err := unmarshalResponse(req, resp, respErr, reply)
	return err
}


func (cli *ExternalProvider) GetToken(r *http.Request) string {
	token := r.Header.Get("x-user-token")
	if token == "" {
		cookieToken, err := r.Cookie("officetoken")
		if err == nil {
			token = cookieToken.Value
		}
	}
	return token
}



func unmarshalResponse(req *http.Request, resp *http.Response, respErr error, reply interface{}) error {
	if respErr != nil {
		if urlErr, ok := respErr.(*url.Error); ok {
			switch urlErr.Err {
			case context.Canceled:
				return respErr
			case context.DeadlineExceeded:
				return respErr
			}
		}
		return respErr
	}

	p, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	//klog.Errorf("status:%d body:%s", resp.StatusCode, p)

	if resp.StatusCode != http.StatusOK {
		providerError := &Error{}
		if json.Unmarshal(p, providerError) != nil {
			return fmt.Errorf("invalid response: %s", p)
		}
	} else if reply != nil {
		if err := json.Unmarshal(p, reply); err != nil {
			return fmt.Errorf("invalid response: %s", p)
		} else {
			return nil
		}
	}
	return nil
}


func SetHeaderToken(header http.Header, token string) {
	header.Set(XUserToken, token)
}

func GetHeaderToken(header http.Header) string {
	return header.Get(XUserToken)
}