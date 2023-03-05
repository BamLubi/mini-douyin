// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	userdouyin "mini-douyin/cmd/socity/kitex_gen/userdouyin"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*userdouyin.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserRegister": kitex.NewMethodInfo(userRegisterHandler, newUserServiceUserRegisterArgs, newUserServiceUserRegisterResult, false),
		"UserLogin":    kitex.NewMethodInfo(userLoginHandler, newUserServiceUserLoginArgs, newUserServiceUserLoginResult, false),
		"UserInfo":     kitex.NewMethodInfo(userInfoHandler, newUserServiceUserInfoArgs, newUserServiceUserInfoResult, false),
		"PublishList":  kitex.NewMethodInfo(publishListHandler, newUserServicePublishListArgs, newUserServicePublishListResult, false),
		"FavoriteList": kitex.NewMethodInfo(favoriteListHandler, newUserServiceFavoriteListArgs, newUserServiceFavoriteListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "userdouyin",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userdouyin.UserServiceUserRegisterArgs)
	realResult := result.(*userdouyin.UserServiceUserRegisterResult)
	success, err := handler.(userdouyin.UserService).UserRegister(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserRegisterArgs() interface{} {
	return userdouyin.NewUserServiceUserRegisterArgs()
}

func newUserServiceUserRegisterResult() interface{} {
	return userdouyin.NewUserServiceUserRegisterResult()
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userdouyin.UserServiceUserLoginArgs)
	realResult := result.(*userdouyin.UserServiceUserLoginResult)
	success, err := handler.(userdouyin.UserService).UserLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserLoginArgs() interface{} {
	return userdouyin.NewUserServiceUserLoginArgs()
}

func newUserServiceUserLoginResult() interface{} {
	return userdouyin.NewUserServiceUserLoginResult()
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userdouyin.UserServiceUserInfoArgs)
	realResult := result.(*userdouyin.UserServiceUserInfoResult)
	success, err := handler.(userdouyin.UserService).UserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserInfoArgs() interface{} {
	return userdouyin.NewUserServiceUserInfoArgs()
}

func newUserServiceUserInfoResult() interface{} {
	return userdouyin.NewUserServiceUserInfoResult()
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userdouyin.UserServicePublishListArgs)
	realResult := result.(*userdouyin.UserServicePublishListResult)
	success, err := handler.(userdouyin.UserService).PublishList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServicePublishListArgs() interface{} {
	return userdouyin.NewUserServicePublishListArgs()
}

func newUserServicePublishListResult() interface{} {
	return userdouyin.NewUserServicePublishListResult()
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userdouyin.UserServiceFavoriteListArgs)
	realResult := result.(*userdouyin.UserServiceFavoriteListResult)
	success, err := handler.(userdouyin.UserService).FavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceFavoriteListArgs() interface{} {
	return userdouyin.NewUserServiceFavoriteListArgs()
}

func newUserServiceFavoriteListResult() interface{} {
	return userdouyin.NewUserServiceFavoriteListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserRegister(ctx context.Context, req *userdouyin.UserRegisterRequest) (r *userdouyin.UserRegisterResponse, err error) {
	var _args userdouyin.UserServiceUserRegisterArgs
	_args.Req = req
	var _result userdouyin.UserServiceUserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, req *userdouyin.UserLoginRequest) (r *userdouyin.UserLoginResponse, err error) {
	var _args userdouyin.UserServiceUserLoginArgs
	_args.Req = req
	var _result userdouyin.UserServiceUserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, req *userdouyin.UserInfoRequest) (r *userdouyin.UserInfoResponse, err error) {
	var _args userdouyin.UserServiceUserInfoArgs
	_args.Req = req
	var _result userdouyin.UserServiceUserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, req *userdouyin.PublishListRequest) (r *userdouyin.PublishListResponse, err error) {
	var _args userdouyin.UserServicePublishListArgs
	_args.Req = req
	var _result userdouyin.UserServicePublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, req *userdouyin.FavoriteListRequest) (r *userdouyin.FavoriteListResponse, err error) {
	var _args userdouyin.UserServiceFavoriteListArgs
	_args.Req = req
	var _result userdouyin.UserServiceFavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
