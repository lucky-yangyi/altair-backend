package cache

import (
	"altair-backend/internal/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

const UserInfoPrefix = "cache:userInfo:"
const UserIdPrefix = "cache:id:"

const AdminInfoPrefix = "cache:adminInfo:"
const CompanyInfoPrefix = "cache:companyInfo:"

type AdminLogin struct {
	Admin model.AdminNoPassword `json:"admin"`
	Token Token                 `json:"token"`
}

type Login struct {
	User  model.CompanyMember `json:"user"`
	Token Token               `json:"token"`
}
type Token struct {
	AccessKey string `json:"accessKey"` //公钥
	SecretKey string `json:"secretKey"` //私钥
}

func SetUserInRedis(userInfo Login) error {
	rds := RedisPool().Get()
	defer rds.Close()
	//构建
	userInfoStr, _ := json.Marshal(userInfo)

	_, err := rds.Do("SETEX", UserInfoPrefix+":"+userInfo.Token.AccessKey, "86400", userInfoStr)
	_, err = rds.Do("SETEX", UserIdPrefix+":"+fmt.Sprintf("%v", userInfo.User.ID)+":"+userInfo.Token.AccessKey, "86400", userInfo.Token.AccessKey)
	return err
}

func SetAdminInRedis(adminInfo AdminLogin) error {
	rds := RedisPool().Get()
	defer rds.Close()
	//构建
	userInfoStr, _ := json.Marshal(adminInfo)

	_, err := rds.Do("SETEX", AdminInfoPrefix+":"+adminInfo.Token.AccessKey, "86400", userInfoStr)
	return err
}

func DelUserInKeyRedis(key string) error {
	rds := RedisPool().Get()
	defer rds.Close()
	_, err := rds.Do("DEL", UserInfoPrefix+":"+key)
	return err
}

func DelUserInIdRedis(id uint64) (accessKey []string) {
	rds := RedisPool().Get()
	defer rds.Close()
	res, _ := rds.Do("KEYS", UserIdPrefix+":"+fmt.Sprintf("%v", id)+":*")
	bytes, _ := json.Marshal(res)
	fmt.Println(string(bytes))
	var array []string

	json.Unmarshal(bytes, &array)
	//base64解码
	for _, v := range array {
		ress, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			fmt.Println("err.Error", err.Error())
		}
		fmt.Println("string(ress),", string(ress))
		credential := strings.Split(string(ress), ":")
		if len(credential) > 1 {
			accessKey = append(accessKey, credential[len(credential)-1])
			err = DelUserInKeyRedis(credential[len(credential)-1])
			_, err = rds.Do("DEL", string(ress))
		}
	}

	return
}

func DelAdminInRedis(key string) error {
	rds := RedisPool().Get()
	defer rds.Close()
	_, err := rds.Do("DEL", AdminInfoPrefix+":"+key)
	return err
}

// 从 Redis 缓存中获取用户数据
func UserInfoGet(accessKey string) ([]byte, error) {
	rds := RedisPool().Get()
	defer rds.Close()
	data, err := redis.Bytes(rds.Do("GET", UserInfoPrefix+":"+accessKey))
	return data, err
}

// 从 Redis 缓存中获取admin数据
func AdminInfoGet(accessKey string) ([]byte, error) {
	rds := RedisPool().Get()
	defer rds.Close()
	data, err := redis.Bytes(rds.Do("GET", AdminInfoPrefix+":"+accessKey))
	return data, err
}

func SetCompanyInRedis(id uint64, status bool) error {
	rds := RedisPool().Get()
	defer rds.Close()
	//构建
	companyInfoStr, _ := json.Marshal(status)

	_, err := rds.Do("SETEX", CompanyInfoPrefix+":"+fmt.Sprintf("%+v", id), "86400", companyInfoStr)
	return err
}

func DelCompanyInRedis(key uint64) error {
	rds := RedisPool().Get()
	defer rds.Close()
	_, err := rds.Do("DEL", CompanyInfoPrefix+":"+fmt.Sprintf("%+v", key))
	return err
}

func CompanyInfoGet(id uint64) (status bool, err error) {
	rds := RedisPool().Get()
	defer rds.Close()
	data, err := redis.Bytes(rds.Do("GET", CompanyInfoPrefix+":"+fmt.Sprintf("%+v", id)))
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(data, &status)

	return
}
