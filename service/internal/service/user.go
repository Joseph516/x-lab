package service

type RegisterRequest struct {
	// ID         uint32 `form:"id" binding:"required,gte=1"`
	Username   string `form:"username" binding:"required,min=2,max=20"`
	Password   string `form:"password" binding:"required,min=4,max=32"`
	Repassword string `form:"repassword" binding:"required,min=4,max=32"`
	// State      uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type LoginRequest struct {
	Username   string `form:"username" binding:"required,min=2,max=20"`
	Password   string `form:"password" binding:"required,min=4,max=32"`
}

func (svc *Service) Register(param *RegisterRequest) error {
	return svc.dao.Register(param.Username, param.Password)
}

func (svc *Service) LoginIn(param *LoginRequest) error {
	return svc.dao.LoginIn(param.Username, param.Password)
}