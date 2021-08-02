/*
@Title : cache
@Description :
@Author : 谭靖渝
@Update : 2021/8/1 23:10
*/
package history

import (
	"DanMu/pkg/db/redis"
	"encoding/json"
)

type Massage struct {
	DanMu []string
}

func SetInfo(key string, massage Massage) (err error) {
	valueByte, err := json.Marshal(massage)
	if err != nil {
		return
	}

	err = redis.Instance().Set(key, string(valueByte), 0)
	if err != nil {
		return
	}
	return
}

func GetInfo(key string) (massage Massage, err error) {
	data, err := redis.Instance().Get(key)
	if err != nil {
		return
	}
	massage = Massage{}
	err = json.Unmarshal(data.([]byte), &massage)
	if err != nil {
		return
	}
	return
}
