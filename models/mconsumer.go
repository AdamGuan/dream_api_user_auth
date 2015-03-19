package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"dream_api_user_auth/helper"
//	"fmt"
//	"crypto/md5"
//	"strings"
//	"github.com/astaxie/beego/config" 
//	"github.com/astaxie/beego"
)

func init() {
}

type MConsumer struct {
}

type userInfoa struct {
	F_uid string
	F_phone_number string
	F_crate_datetime string
	F_modify_datetime string
}

//根据手机号码获取uid
func (u *MConsumer) GetUidByPhone(phone string)string{
	uid := ""
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_phone=? LIMIT 1", phone).Values(&maps)
	if err == nil && num > 0 {
		uid = maps[0]["F_user_name"].(string)
	}
	return uid
}

//检查手机号码是否可用
func (u *MConsumer) CheckPhoneValid(phone string)int{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_phone=? LIMIT 1", phone).Values(&maps)
	if err == nil && num <= 0 {
		return 0
	}
	return -23

}

//检查用户ID是否可用
func (u *MConsumer) checkUserIdValid(uid string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=? LIMIT 1", uid).Values(&maps)
	if err == nil && num <= 0 {
		return true
	}
	return false
}

//创建一个可用的用户ID
func (u *MConsumer) CreateUid()string{
	uid := ""
	for{
		uid = helper.GetGuid()
		if u.checkUserIdValid(uid){
			break
		}
	}
	return uid
}

//检查手机号码是否存在
func (u *MConsumer) CheckPhoneExists(phone string)bool{
	if u.CheckPhoneValid(phone) != 0{
		return true
	}
	return false
}

//检查uid是否存在
func (u *MConsumer) CheckUserIdExists(uid string)bool{
	if !u.checkUserIdValid(uid){
		return true
	}
	return false
}

//检查手机号码与密码是否正确
func (u *MConsumer) CheckPhoneAndPwd(phone string,userPwd string)bool{
	if len(phone) <= 0 || len(userPwd) <= 0{
		return false
	}
	o := orm.NewOrm()
	var maps []orm.Params
//	pwd := helper.Sha1(userPwd)
	pwd := userPwd
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_phone=? AND F_user_password = ? LIMIT 1", phone,pwd).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

//检查用户ID与密码是否正确
func (u *MConsumer) CheckUserIdAndPwd(uid string,userPwd string)bool{
	if len(uid) <= 0 || len(userPwd) <= 0{
		return false
	}
	o := orm.NewOrm()
	var maps []orm.Params
	//helper.Sha1(userPwd)
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=? AND F_user_password = ? LIMIT 1", uid,userPwd).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

//检查用户密码
func (u *MConsumer) CheckUserPwdValid(userPwd string)int{
	if helper.CheckPwdValid(userPwd){
		return 0
	}
	return -3
}


//修改用户密码
func (u *MConsumer) ModifyUserPwdByUid(userId string,userPwd string)int{
	result := -1
	res := u.CheckUserIdExists(userId)
	if res {
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		result = -1
		//写入数据库
		o := orm.NewOrm()
		_, err := o.Raw("UPDATE t_user SET F_user_password=?,F_modify_datetime=? WHERE F_user_name=?",userPwd,helper.GetNowDateTime(),userId).Exec()
		if err == nil {
			result = 0
		}
	}
	return result
}

//修改用户密码(根据手机)
func (u *MConsumer) ModifyUserPwdByPhone(phone string,userPwd string)int{
	result := -1
	res := u.CheckPhoneExists(phone)
	if res {
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		result = -1
		//写入数据库
		o := orm.NewOrm()
		_, err := o.Raw("UPDATE t_user SET F_user_password=?,F_modify_datetime=? WHERE F_user_phone=?",userPwd,helper.GetNowDateTime(),phone).Exec()
		if err == nil {
			result = 0
		}
	}
	return result
}

//获取用户的信息
func (u *MConsumer) GetUserInfoByUid(userId string)userInfoa{
	info := userInfoa{}
	if len(userId) > 0 {
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT * FROM t_user WHERE F_user_name=? LIMIT 1", userId).Values(&maps)
		if err == nil && num > 0 {
			info.F_uid = maps[0]["F_user_name"].(string)
			info.F_phone_number = maps[0]["F_user_phone"].(string)
			//创建时间
			info.F_crate_datetime = ""
			if maps[0]["F_crate_datetime"] != nil{
				info.F_crate_datetime = maps[0]["F_crate_datetime"].(string)
			}
			//修改时间
			info.F_modify_datetime = ""
			if maps[0]["F_modify_datetime"] != nil{
				info.F_modify_datetime = maps[0]["F_modify_datetime"].(string)
			}
		}
	}
	return info
}

//获取用户的信息(根据手机号码)
func (u *MConsumer) GetUserInfoByPhone(phone string)userInfoa{
	uid := u.GetUidByPhone(phone)
	if len(uid) > 0{
		return u.GetUserInfoByUid(uid)
	}
	return userInfoa{}
}

//修改用户手机号码
func (u *MConsumer) ModifyUserPhone(phone string,uid string)int{
	result := -1
	
	result = u.CheckPhoneValid(phone)
	if result == 0 {
		o := orm.NewOrm()
		res, err := o.Raw("UPDATE t_user SET F_user_phone=?,F_modify_datetime=? WHERE F_user_name=?",phone,helper.GetNowDateTime(),uid).Exec()
		if err == nil {
			num, err := res.RowsAffected()
			if err == nil && num > 0{
				result = 0
			}
		}
	}
	return result
}