package main

import (
	"fmt"
	mcrcon "github.com/Kelwing/mc-rcon"
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"strings"
)

//go:generate cqcfg -c .
// cqp: 名称: CoolQ-mcrcon
// cqp: 版本: 1.0.0:1
// cqp: 作者: Underworld
// cqp: 简介: 在QQ群里与Minecraft Server交互~
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "me.cqp.underworld.mcrcon" // TODO: 修改为这个插件的ID
	cqp.GroupMsg = onGroupMsg
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	defer handlePanic()
	reply, flag := handleCmd(msg,fromQQ)
	if flag {
		cqp.SendGroupMsg(fromGroup, reply)
	}
	return 0
}

func handleCmd(msg string,fromQQ int64) (string, bool) {
	if strings.HasPrefix(msg, "[CQ:at,qq=3***********]") {  //当被@时
		//检查格式
		if fromQQ != 2************ {
			return "您没有这个权限！", true
		} 
		command := strings.TrimSpace(msg[22:])
		conn := new(mcrcon.MCConn)
		err := conn.Open("localhost:25575", "Like")
		if err != nil {
			return "Open failed!", true
		}
		defer conn.Close()

		err = conn.Authenticate()
		if err != nil {
			return "Auth failed!", true
		}

		resp, err := conn.SendCommand(command)
		if err != nil {
			return "Command failed!", true
		}
		return resp, true
	}
	return "",false
}

func handlePanic() {
	if r := recover(); r != nil {
		cqp.AddLog(cqp.Error, "未知错误", fmt.Sprint(r))
		return
	}
}
