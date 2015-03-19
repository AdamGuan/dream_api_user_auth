package models

import (
	"dream_api_user_auth/helper"
)

func init() {
}

type MSign struct {
}

//检查sign是否正确
func (u *MSign) CheckSign(signOri string,token string) bool{
	return helper.CheckSign(signOri,token)
}