package micro

import (
	"context"
	"testing"

	"github.com/iGoogle-ink/gotil/micro/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func TestServer(t *testing.T) {
	// 解开注释测试

	//InitServer("service.server.serverName", "v1.0.0", nil, nil, func(s server.Server) {
	//	if err := proto.RegisterGreeterHandler(s, new(Greeter)); err != nil {
	//		panic(fmt.Sprintf("service.server.helloworld start failed: %+v", err))
	//	}
	//})
	//
	//ch := make(chan os.Signal)
	//signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	//for {
	//	si := <-ch
	//	switch si {
	//	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	//		time.Sleep(time.Second)
	//		xlog.Warn("service.server.serverName stop")
	//		time.Sleep(time.Second)
	//		return
	//	case syscall.SIGHUP:
	//	default:
	//		return
	//	}
	//}
}
