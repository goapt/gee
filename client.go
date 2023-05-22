package gee

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/goapt/gee/encoding"
	"github.com/goapt/gee/encoding/json"
	"github.com/goapt/gee/errors"
)

type AuthFunc = func(req *http.Request) error
type ErrorDecoderFunc = func(ctx context.Context, res *http.Response) error

type clientOptions struct {
	timeout       time.Duration
	authorization AuthFunc
	errorDecoder  ErrorDecoderFunc
}

type ClientOption func(c *clientOptions)

func WithTimeout(t time.Duration) ClientOption {
	return func(c *clientOptions) {
		c.timeout = t
	}
}

func WithAuthorization(fn AuthFunc) ClientOption {
	return func(c *clientOptions) {
		c.authorization = fn
	}
}

func WithErrorDecoder(errorDecoder ErrorDecoderFunc) ClientOption {
	return func(c *clientOptions) {
		c.errorDecoder = errorDecoder
	}
}

var defaultErrorDecoder = func(ctx context.Context, res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("[WARN] gee client close response body error: %v", err)
		}
	}()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Errorf(res.StatusCode, "IoReadError", err.Error())
	}
	e := new(errors.Error)
	if err = encoding.GetCodec(json.Name).Unmarshal(data, e); err == nil {
		e.Code = int32(res.StatusCode)
		if e.GetReason() == "" {
			e.Reason = errors.UnknownReason
		}

		if e.GetMessage() == "" {
			e.Message = string(data)
		}
		return e
	}
	return errors.Errorf(res.StatusCode, errors.UnknownReason, string(data))
}

// Client is an HTTP client.
type Client struct {
	opts     clientOptions
	cc       *http.Client
	Endpoint string
}

func NewClient(client *http.Client, endpoint string, opts ...ClientOption) *Client {
	options := clientOptions{
		timeout:      5 * time.Second,
		errorDecoder: defaultErrorDecoder,
	}
	for _, o := range opts {
		o(&options)
	}
	return &Client{
		opts:     options,
		cc:       client,
		Endpoint: endpoint,
	}
}

func (c *Client) Invoke(ctx context.Context, method, path string, args any, reply any) error {
	var (
		body io.Reader
	)
	if args != nil {
		data, err := encoding.GetCodec(json.Name).Marshal(args)
		if err != nil {
			return err
		}
		body = bytes.NewReader(data)
	}
	url := fmt.Sprintf("https://%s%s", strings.TrimRight(c.Endpoint, "/"), path)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Gee-Client/v1.0.0")

	if c.opts.authorization != nil {
		if err := c.opts.authorization(req); err != nil {
			return err
		}
	}

	resp, err := c.cc.Do(req)
	if err == nil {
		err = c.opts.errorDecoder(req.Context(), resp)
	}

	if err != nil {
		return err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("[WARN] gee client close response body error: %s", err)
		}
	}()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return encoding.GetCodec(json.Name).Unmarshal(data, reply)
}
