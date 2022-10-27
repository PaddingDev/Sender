package model

type Config struct {
	IsDebug         bool   `json:"is_debug"`
	ListenAddr      string `json:"listen_addr"`
	StorageBasePath string `json:"storage_path"`

	MaxExpireTime  int  `json:"max_exp_time"`
	MinExpireTime  int  `json:"min_exp_time"`
	ExpireTimeUnit byte `json:"exp_unit"`
	/*
		'm'  => minute
		's'  => second
		'x'  => ms
		'h'  => hour
		'd'  => day
	*/

	AuthType      AuthTypeEnum   `json:"auth_type"`
	BasicAuthList *BasicAuthList `json:"basic_auth"`

	MaxUploadBytes        int64 `json:"max_upload_bytes"`
	MaxNormalRequestBytes int64 `json:"max_normal_req_bytes"`
}

type AuthTypeEnum string

const (
	None         AuthTypeEnum = "None"
	BasicAuth    AuthTypeEnum = "BasicAuth"
	OnlyDebugMod AuthTypeEnum = "OnlyDebugMod"
)

type BasicAuthList map[string]string
