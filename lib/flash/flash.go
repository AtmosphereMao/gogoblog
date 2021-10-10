package flash

import (
	"encoding/gob"
	"myblog/lib/session"
)
type Flashes map[string]interface{}

// 存入会话数据里的 key
var flashKey = "_flashes"

func init(){
	gob.Register(Flashes{})
}

func All() Flashes {
	val := session.Get(flashKey)
	// 读取是必须做类型检测
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}
	// 读取即销毁，直接删除
	session.Forget(flashKey)
	return flashMessages
}

func addFlash(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}

func Info(message string) {
	addFlash("info", message)
}

func Warning(message string) {
	addFlash("warning", message)
}

func Success(message string) {
	addFlash("success", message)
}

func Danger(message string) {
	addFlash("danger", message)
}