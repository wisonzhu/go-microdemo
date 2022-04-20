package main
import (
	"gokitdemo/myEndPoints"
	"gokitdemo/myServices"
	"gokitdemo/myTransports"
	"github.com/go-kit/kit/transport/http"
	"log"
	sthttp "net/http"
)

func main() {
	//业务接口服务
	s := myServices.UserService{}
	//使用myEndPoints创建业务服务
	getName := myEndPoints.MakeServerEndPointGetName(s)
	//使用 kit 创建 handler
	// 固定格式
	// 传入 业务服务 以及 定义的 加密解密方法
	server := http.NewServer(getName, myTransports.GetNameDecodeRequest, myTransports.GetNameEncodeResponse)
	//监听服务
	log.Println(sthttp.ListenAndServe(":8080", server))
}