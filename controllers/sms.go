package controllers

import (
	"dream_api_user_auth/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_user_auth/helper"
	"github.com/astaxie/beego/config" 
)

//短信(每个用户短信发送限制为1分钟的一次)
type SmsController struct {
	beego.Controller
}

//json echo
func (u0 *SmsController) jsonEcho(datas map[string]interface{},u *SmsController) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	
	datas["responseMsg"] = ""
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	}

	u.Data["json"] = datas
	//log
	u.logEcho(datas)

	u.ServeJson()
}

//sign check, token为包名md5
func (u0 *SmsController) checkSign(u *SmsController)int {
	result := -6
	pkg := u.Ctx.Request.Header.Get("pkg")
	sign := u.Ctx.Request.Header.Get("sign")
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		result = -7
	}else{
		var signObj *models.MSign
		if re := signObj.CheckSign(sign, helper.Md5(pkg)); re == true {
			result = 0
		}
	}
	return result
}

// @Title 短信验证码验证
// @Description 短信验证码验证(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	num					form	string	true	验证码
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /smsvalid/:mobilePhoneNumber [post]
func (u *SmsController) Smsvalid() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && len(num) > 0 {
		datas["responseNo"] = -1
		pkgConfig := pkgObj.GetPkgConfig(pkg)
		if len(pkgConfig) > 0{
			res := smsObj.ValidMsm(pkg,num,mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"])
			if len(res) == 0{
				datas["responseNo"] = 0
				smsObj.AddMsmActionvalid(mobilePhoneNumber,pkg,num)
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 发送一条短信验证码(重置密码时)
// @Description 发送一条短信验证码(重置密码时)(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /resetpwd/:mobilePhoneNumber [get]
func (u *SmsController) ResetPwdGetSms() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MConsumer
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		res := userObj.CheckPhoneExists(mobilePhoneNumber)
		if res {
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(mobilePhoneNumber,pkg) {
				smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"])
				if len(res) == 0{
					datas["responseNo"] = 0
					smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				}else{
					smsObj.DeleteMsmRate(mobilePhoneNumber,pkg)
				}
			}
		}else{
			datas["responseNo"] = -4
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -4
	}

	//return
	u.jsonEcho(datas,u)
}

// @Title 发送一条短信验证码(更换手机号码)
// @Description 发送一条短信验证码(更换手机号码)(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码(新的号码)
// @Param	uid					query	string	true	uid
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /phone/:mobilePhoneNumber [get]
func (u *SmsController) ChangePhoneSms() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MConsumer
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	uid := u.Ctx.Request.FormValue("uid")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && userObj.CheckUserIdExists(uid) {
		datas["responseNo"] = -1
		if userObj.CheckPhoneExists(mobilePhoneNumber){
			datas["responseNo"] = -23
		}else{
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(mobilePhoneNumber,pkg) {
				smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"])
				if len(res) == 0{
					datas["responseNo"] = 0
					smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				}else{
					smsObj.DeleteMsmRate(mobilePhoneNumber,pkg)
				}
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -10
	}

	//return
	u.jsonEcho(datas,u)
}

//记录请求
func (u *SmsController) logRequest() {
	var logObj *models.MLog
	logObj.LogRequest(u.Ctx)
}

//记录返回
func (u *SmsController) logEcho(datas map[string]interface{}) {
	var logObj *models.MLog
	logObj.LogEcho(datas)
}