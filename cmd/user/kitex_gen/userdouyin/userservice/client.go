// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	userdouyin "mini-douyin/cmd/user/kitex_gen/userdouyin"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserRegister(ctx context.Context, req *userdouyin.UserRegisterRequest, callOptions ...callopt.Option) (r *userdouyin.UserRegisterResponse, err error)
	UserLogin(ctx context.Context, req *userdouyin.UserLoginRequest, callOptions ...callopt.Option) (r *userdouyin.UserLoginResponse, err error)
	UserInfo(ctx context.Context, req *userdouyin.UserInfoRequest, callOptions ...callopt.Option) (r *userdouyin.UserInfoResponse, err error)
	PublishList(ctx context.Context, req *userdouyin.PublishListRequest, callOptions ...callopt.Option) (r *userdouyin.PublishListResponse, err error)
	FavoriteList(ctx context.Context, req *userdouyin.FavoriteListRequest, callOptions ...callopt.Option) (r *userdouyin.FavoriteListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) UserRegister(ctx context.Context, req *userdouyin.UserRegisterRequest, callOptions ...callopt.Option) (r *userdouyin.UserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserRegister(ctx, req)
}

func (p *kUserServiceClient) UserLogin(ctx context.Context, req *userdouyin.UserLoginRequest, callOptions ...callopt.Option) (r *userdouyin.UserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, req)
}

func (p *kUserServiceClient) UserInfo(ctx context.Context, req *userdouyin.UserInfoRequest, callOptions ...callopt.Option) (r *userdouyin.UserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserInfo(ctx, req)
}

func (p *kUserServiceClient) PublishList(ctx context.Context, req *userdouyin.PublishListRequest, callOptions ...callopt.Option) (r *userdouyin.PublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, req)
}

func (p *kUserServiceClient) FavoriteList(ctx context.Context, req *userdouyin.FavoriteListRequest, callOptions ...callopt.Option) (r *userdouyin.FavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, req)
}
