package cfg

import (
	"encoding/json"
	"github.com/PaddingDEV/Sender/model"
	"github.com/PaddingDEV/Sender/utils"
	"time"
)

var _cfg *model.Config = nil
var _expUnit time.Duration

func GetCfg() *model.Config {
	return GetCfgFromFile("config.yml")
}

func GetExpTimeUnit() time.Duration {
	if _cfg == nil {
		GetCfg()
	}
	return _expUnit
}

func GetCfgFromFile(path string) *model.Config {
	if _cfg != nil {
		return _cfg
	}
	byteValue, err := utils.ReadFileToByte(path)
	utils.PanicIfNotNil(err, "IO: %v\n")

	err = json.Unmarshal(byteValue, &_cfg)
	utils.PanicIfNotNil(err, "JSON: %v\n")

	switch _cfg.ExpireTimeUnit {
	case 'x':
		_expUnit = time.Millisecond
	case 's':
		_expUnit = time.Second
	case 'm':
		_expUnit = time.Minute
	case 'h':
		_expUnit = time.Hour
	case 'd':
		_expUnit = 24 * time.Hour
	default:
		// TODO
	}

	return _cfg
}
