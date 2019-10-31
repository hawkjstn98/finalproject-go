package helper

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
)

func SavePassword(pass string) string{
	hsh:= md5.New();
	hsh.Write([]byte(mongo_constant.Salt+pass+mongo_constant.Salt))
	return hex.EncodeToString(hsh.Sum(nil))
}