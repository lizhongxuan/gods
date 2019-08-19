package main

import (
	"time"
	"fmt"
)

type Options struct {
	timeout time.Duration
	insecure bool
	isOpen bool
}

type Client struct {
	conn struct{}
	opts Options
}

type DialOption func(*Options)

func WithOpen() DialOption {
	return func(o *Options) {
		o.isOpen = true
	}
}

func WithInsecure() DialOption {
	return func(o *Options) {
		o.insecure = true
	}
}

func WithTimeout(d time.Duration) DialOption {
	return func(o *Options) {
		o.timeout = d
	}
}

func NewDialClient(dialopts ...DialOption)*Client {
	cc := &Client{
		conn: struct{}{},
	}

	for _,opt := range dialopts {
		opt(&cc.opts)
		
	}
	return cc
}

func main() {
	c :=NewDialClient()
	fmt.Println("client insecure:",c.opts.insecure)
	fmt.Println("client timeout:",c.opts.timeout)
	fmt.Println("client isOpen:",c.opts.isOpen)

	c2 :=NewDialClient(WithInsecure(),WithTimeout(100),WithOpen())
	fmt.Println("client2 insecure:",c2.opts.insecure)
	fmt.Println("client2 timeout:",c2.opts.timeout)
	fmt.Println("client2 isOpen:",c2.opts.isOpen)
}