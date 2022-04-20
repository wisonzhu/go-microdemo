package myTransports
import (
	"context"
	"encoding/json"
	"errors"
	"gokitdemo/myEndPoints"
	"net/http"
	"strconv"
)
//Transport层主要负责与传输协议HTTP，GRPC，THRIFT等相关的逻辑；
func GetNameDecodeRequest (c context.Context,r *http.Request) (interface{}, error){
	id := r.URL.Query().Get("id")
	if id ==""{
		return nil,errors.New("无效参数")
	}
	intid, err := strconv.Atoi(id)
	if err != nil{
		return nil,errors.New("无效参数")
	}
	return myEndPoints.UserRequest{Id: intid},nil
}

func GetNameEncodeResponse(c context.Context,w http.ResponseWriter,res interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	return json.NewEncoder(w).Encode(res)
}