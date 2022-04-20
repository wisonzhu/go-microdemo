package myServices


type IUserService interface {
	GetName(userId int) string
}

type UserService struct {
}

func (s UserService) GetName(uid int) string {
	if uid == 10{
		return "wangqiang"
	}
	return "xiaoqiang"
}