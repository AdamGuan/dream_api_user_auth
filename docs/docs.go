package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1","swaggerVersion":"1.2","apis":[{"path":"/sms","description":"短信(每个用户短信发送限制为1分钟的一次)\n"},{"path":"/consumer","description":"用户\n"}],"info":{"title":"用户系统 API v1"}}`
    Subapi string = `{"/consumer":{"apiVersion":"1","swaggerVersion":"1.2","basePath":"","resourcePath":"/consumer","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/auth/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"GET","nickname":"用户登录验证(利用手机号码登录)","type":"","summary":"用户登录验证(利用手机号码登录)(token: md5(pkg))","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"pwd","description":"密码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MUserLoginResp","responseModel":"MUserLoginResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/resetpwd","description":"","operations":[{"httpMethod":"PUT","nickname":"重置密码(利用手机号码重置密码)","type":"","summary":"重置密码(利用手机号码重置密码)(token: md5(pkg))","parameters":[{"paramType":"form","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"pwd","description":"密码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"num","description":"验证码(经过验证成功后的)","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/pwd/:uid","description":"","operations":[{"httpMethod":"PUT","nickname":"修改密码","type":"","summary":"修改密码(token: md5(pkg))","parameters":[{"paramType":"path","name":"uid","description":"用户ID","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"oldPwd","description":"旧密码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"newPwd","description":"新密码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/exists/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"GET","nickname":"验证手机号码是否已注册","type":"","summary":"验证手机号码是否已注册(token: md5(pkg))","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/phone/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"PUT","nickname":"修改用户手机号码","type":"","summary":"修改用户手机号码(token: md5(pkg))","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码(新的手机号码)","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"uid","description":"uid","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"num","description":"验证码(经过验证成功后的)","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MModifyPhoneResp","responseModel":"MModifyPhoneResp"},{"code":401,"message":"无权访问","responseModel":""}]}]}],"models":{"MModifyPhoneResp":{"id":"MModifyPhoneResp","properties":{"responseMsg":{"type":"string","description":"","format":""},"responseNo":{"type":"int","description":"","format":""}}},"MResp":{"id":"MResp","properties":{"responseMsg":{"type":"string","description":"","format":""},"responseNo":{"type":"int","description":"","format":""}}},"MUserLoginResp":{"id":"MUserLoginResp","properties":{"F_crate_datetime":{"type":"string","description":"","format":""},"F_modify_datetime":{"type":"string","description":"","format":""},"F_phone_number":{"type":"string","description":"","format":""},"F_uid":{"type":"string","description":"","format":""},"responseMsg":{"type":"string","description":"","format":""},"responseNo":{"type":"int","description":"","format":""}}}}},"/sms":{"apiVersion":"1","swaggerVersion":"1.2","basePath":"","resourcePath":"/sms","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/smsvalid/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"POST","nickname":"短信验证码验证","type":"","summary":"短信验证码验证(token: md5(pkg))","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"num","description":"验证码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/resetpwd/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"GET","nickname":"发送一条短信验证码(重置密码时)","type":"","summary":"发送一条短信验证码(重置密码时)(token: md5(pkg))","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/phone/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"GET","nickname":"发送一条短信验证码(更换手机号码)","type":"","summary":"发送一条短信验证码(更换手机号码)(token: md5(pkg))","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码(新的号码)","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"uid","description":"uid","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]}],"models":{"MResp":{"id":"MResp","properties":{"responseMsg":{"type":"string","description":"","format":""},"responseNo":{"type":"int","description":"","format":""}}}}}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	err := json.Unmarshal([]byte(Rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(Subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = BasePath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
