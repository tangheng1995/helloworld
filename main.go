package main

import (
	"fmt"
	"time"
)

// Server 服务端
type Server struct {
	host string
	port int
	opts Options
}

// Options 服务端配置项
type Options struct {
	timeout time.Duration
	ssl     bool
	maxConn int
}

// Option 配置接口
type Option interface {
	apply(*Options)
}

type funcOption struct {
	f func(*Options)
}

func (fo *funcOption) apply(o *Options) {
	fo.f(o)
}

func newFuncOption(f func(*Options)) *funcOption {
	return &funcOption{
		f:f,
	}
}

// WithTimeout 配置服务端超时时间
func WithTimeout(t time.Duration) Option {
	return newFuncOption(func(o *Options){
		o.timeout = t
	})
}

// WithSSL 配置服务端SSL开关
func WithSSL(f bool) Option {
	return newFuncOption(func(o *Options){
		o.ssl = f
	})
}

// WithMaxConn 配置服务端最大连接数
func WithMaxConn(m int) Option {
	return newFuncOption(func(o *Options){
		o.maxConn = m
	})
}

// defaultOptions 服务端默认配置
func defaultOptions() Options {
	return Options{
		timeout: 30 * time.Second,
		ssl:     false,
		maxConn: 200,
	}
}

// NewServer 创建服务端实例
func NewServer(host string, port int, opts ...Option) (server *Server,err error){
	s := &Server{
		host: host,
		port: port,
		opts: defaultOptions(),
	}

	for _, opt := range opts {
		opt.apply(&s.opts)
	}

	return s,nil
}

func main() {
	s,_ := NewServer("localhost", 8000,WithTimeout(10),WithSSL(true),WithMaxConn(100))
	fmt.Println(s)
}
