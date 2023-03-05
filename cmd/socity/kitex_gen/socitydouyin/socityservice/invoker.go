// Code generated by Kitex v0.4.4. DO NOT EDIT.

package socityservice

import (
	server "github.com/cloudwego/kitex/server"
	socitydouyin "mini-douyin/cmd/socity/kitex_gen/socitydouyin"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler socitydouyin.SocityService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
