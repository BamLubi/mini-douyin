#!/usr/bin/env bash

echo "update userservice"
cd cmd/user
kitex -module mini-douyin -service userservice ../../idl/user.thrift

echo "update videoservice"
cd ../video
kitex -module mini-douyin -service videoservice ../../idl/video.thrift

echo "update socityservice"
cd ../socity
kitex -module mini-douyin -service socityservice ../../idl/socity.thrift
kitex -module mini-douyin ../../idl/user.thrift

echo "update client"
cd ../api
kitex -module mini-douyin ../../idl/user.thrift
kitex -module mini-douyin ../../idl/video.thrift
kitex -module mini-douyin ../../idl/socity.thrift