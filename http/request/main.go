package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"util/klog"
	"weboffice/cgo"
	"weboffice/errors"
)
type ErrorReply struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Call(ctx context.Context, addr, path string, args, reply interface{}) error {
	reqBody, _ := json.Marshal(args)
	reqURL := fmt.Sprintf("http://%s%s", addr, path)
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewReader(reqBody))
	if err != nil {
		return errors.Newf(errors.InvalidArgument, err.Error())
	}
	klog.Infof("file id is :%v", req.Header.Get("file-unique-id"))
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		if err == cgo.ErrNothingToPrint {
			return errors.Newf(errors.EmptyFile, err.Error())
		}
		return errors.Newf(errors.InternalError, "call editserver failed: %s", err.Error())
	}

	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errReply ErrorReply
		if err := json.Unmarshal(data, &errReply); err != nil {
			return errors.Newf(errors.InternalError, "%s %s failed, status=%d body=%s",
				req.Method, req.URL.String(), resp.StatusCode, data)
		}
		return errors.Newf(errReply.Code, "%s %s failed, status=%d code=%s detail=%s",
			req.Method, req.URL.String(), resp.StatusCode, errReply.Code, errReply.Message)
	}

	if reply != nil {
		if err := json.Unmarshal(data, reply); err != nil {
			klog.Errorf("unmarshal data is :%v,error is :%v", string(data), err)
			return errors.Newf(errors.InternalError, "%s %s failed, status=%d body=%s",
				req.Method, req.URL.String(), resp.StatusCode, data)
		}
	}
	return nil
}