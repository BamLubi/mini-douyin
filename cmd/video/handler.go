package main

import (
	"context"
	videodouyin "mini-douyin/cmd/video/kitex_gen/videodouyin"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videodouyin.FeedRequest) (resp *videodouyin.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videodouyin.PublishActionRequest) (resp *videodouyin.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}
