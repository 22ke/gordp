package main

import (
	"fmt"
	"gordp/glog"
	"strings"
)

func main(){
	domain   := ""
	username := "kekee"
	password := "97555"
	target   := "127.0.0.1:3389"

	var err error
	g := NewClient(target, glog.NONE)
	//SSL协议登录测试
	err = g.loginForSSL(domain, username, password)
	if err == nil {
		return
	}
	//println(err.Error())
	if strings.Contains(err.Error(),"success") {
		println("login success , ",target," : ",username," : ",password)
	}
	//if err.Error() != "PROTOCOL_RDP" {
	//	fmt.Println("Login Error:", err)
	//	return
	//}
}


func testrdp(target string) {
	domain := ""
	username := "kekee"
	password := "970903Fj"
	//target = "180.102.17.30:3389"
	var err error
	g := NewClient(target, glog.NONE)
	//SSL协议登录测试
	err = g.loginForSSL(domain, username, password)
	if err == nil {
		fmt.Println("Login Success")
		return
	}
	if err.Error() != "PROTOCOL_RDP" {
		fmt.Println("Login Error:", err)
		return
	}
	//RDP协议登录测试 //个人使用起来不太准确
	//err = g.loginForRDP(domain, username, password)
	//if err == nil {
	//	fmt.Println("Login Success")
	//	return
	//} else {
	//	fmt.Println("Login Error:", err)
	//	return
	//}
}

//func TestName(t *testing.T) {
//	targetArr := []string{
//		//"50.57.49.172:3389",
//		//"20.49.22.250:3389",
//		"172.31.61.82:3389",
//	}
//	for _, target := range targetArr {
//		fmt.Println(target)
//		testrdp(target)
//	}
//}
