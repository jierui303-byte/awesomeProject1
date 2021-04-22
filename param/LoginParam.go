package param

type LoginParam struct {
	Name     string `json:"name"`     //用户名
	Password string `json:"password"` //密码
	Id       string `json:"id"`       //验证码ID
	Value    string `json:"value"`    //验证码输入值
}
