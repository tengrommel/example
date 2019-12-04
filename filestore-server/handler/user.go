package handler

import (
	"awesomeProject/filestore-server/db"
	"awesomeProject/filestore-server/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	pwd_salt = "*#890"
)

// SignUpHandler 处理用户请求
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	userName := r.Form.Get("username")
	passwd := r.Form.Get("password")
	if len(userName) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}
	enc_passwd := util.Sha1([]byte(passwd + pwd_salt))
	suc := db.UserSignUp(userName, enc_passwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAIL"))
	}
}

// SignInHandler:登录接口
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("username")
	passWord := r.Form.Get("password")

	encPassword := util.Sha1([]byte(passWord + pwd_salt))
	pwdChecked := db.UserSignIn(userName, encPassword)
	// 1、校验用户名及密码
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}
	// 2、生成访问凭证(token)
	token := GetToken(userName)
	upRes := db.UpdateToken(userName, token)
	if !upRes {
		w.Write([]byte("FAILED"))
		return
	}
	// 3、登录成功重定向到首页
	//w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: userName,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())
}

// UserInfoHandler：查询用户信息
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1、解析请求参数
	r.ParseForm()
	userName := r.Form.Get("username")
	token := r.Form.Get("token")
	// 2、验证token是否有效
	isValidToken := IsTokenValid(token)
	if !isValidToken {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 3、查询用户信息
	user, err := db.GetUserInfo(userName)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 4、组装并且响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

func GetToken(userName string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(userName + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// GenToken：生成token
func IsTokenValid(token string) bool {
	// TODO:判断token的时效性，是否过期
	// TODO:从数据库表tbl_user_token查询userName对应的token信息
	// TODO:对比两个token是否一致
	return true
}
