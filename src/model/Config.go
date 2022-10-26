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
}
