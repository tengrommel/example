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
	w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
}

func GetToken(userName string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(userName + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
