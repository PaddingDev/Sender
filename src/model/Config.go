package model

type Config struct {
	IsDebug    bool   `json:"is_debug"`
	ListenAddr string `json:"listen_addr"`
}