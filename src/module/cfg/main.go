package cfg

import (
	"encoding/json"
	"github.com/PaddingDEV/Sender/model"
	"github.com/PaddingDEV/Sender/utils"
)

var _cfg *model.Config = nil

func GetCfg() *model.Config {
	return GetCfgFromFile("config.yml")
}

func GetCfgFromFile(path string) *model.Config {
	if _cfg != nil {
		return _cfg
	}
	byteValue, err := utils.ReadFileToByte(path)
	utils.PanicIfNotNil(err, "IO: %v\n")

	err = json.Unmarshal(byteValue, &_cfg)
	utils.PanicIfNotNil(err, "JSON: %v\n")

	return _cfg
}
