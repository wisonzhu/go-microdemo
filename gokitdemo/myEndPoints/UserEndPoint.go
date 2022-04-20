package myEndPoints

import (
	"context"
    "gokitdemo/myServices"
    "github.com/go-kit/kit/endpoint"
)
//定义Request、Response格式，并可以使用装饰器(闭包)包装函数,以此来实现各个中间件嵌套
type UserRequest struct {
    Id int `json:"id"`
}


type UserResponse struct {
    Name string `json:"name"`
}

func MakeServerEndPointGetName(s myServices.IUserService) endpoint.Endpoint  {
    return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        r,ok := request.(UserRequest)
        if !ok{
            return response,nil
        }
        return UserResponse{Name: s.GetName(r.Id)},nil
    }
}


