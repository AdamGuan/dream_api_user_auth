package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"],
		beego.ControllerComments{
			"CheckUserAndPwdByPhone",
			`/auth/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"],
		beego.ControllerComments{
			"ResetPwdByPhone",
			`/resetpwd`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"],
		beego.ControllerComments{
			"ModifyPwdByUid",
			`/pwd/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"],
		beego.ControllerComments{
			"CheckUserExists",
			`/exists/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user_auth/controllers:ConsumerController"],
		beego.ControllerComments{
			"ModifyPhone",
			`/phone/:mobilePhoneNumber`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_user_auth/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_user_auth/controllers:SmsController"],
		beego.ControllerComments{
			"Smsvalid",
			`/smsvalid/:mobilePhoneNumber`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dream_api_user_auth/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_user_auth/controllers:SmsController"],
		beego.ControllerComments{
			"ResetPwdGetSms",
			`/resetpwd/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_user_auth/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_user_auth/controllers:SmsController"],
		beego.ControllerComments{
			"ChangePhoneSms",
			`/phone/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

}
