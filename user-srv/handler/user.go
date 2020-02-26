package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"
	us "user-srv/model/user"
	s "user-srv/proto/user"
)

//type User struct{}

//// Call is a single request handler called via client.Call or the generated client code
//func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
//	log.Log("Received User.Call request")
//	rsp.Msg = "Hello " + req.Name
//	return nil
//}
//
//// Stream is a server side stream handler called via client.Stream or the generated client code
//func (e *User) Stream(ctx context.Context, req *user.StreamingRequest, stream user.User_StreamStream) error {
//	log.Logf("Received User.Stream request with count: %d", req.Count)
//
//	for i := 0; i < int(req.Count); i++ {
//		log.Logf("Responding: %d", i)
//		if err := stream.Send(&user.StreamingResponse{
//			Count: int64(i),
//		}); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
//func (e *User) PingPong(ctx context.Context, stream user.User_PingPongStream) error {
//	for {
//		req, err := stream.Recv()
//		if err != nil {
//			return err
//		}
//		log.Logf("Got ping %v", req.Stroke)
//		if err := stream.Send(&user.Pong{Stroke: req.Stroke}); err != nil {
//			return err
//		}
//	}
//}

type Service struct{}

var (
	userService us.Service
)

// Init 初始化handler
func Init() {

	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {

	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return err
	}

	rsp.User = user
	rsp.Success = true

	return nil
}
