package main

import (
	"sync"
	"time"
)

func main() {
	println(config == nil) // true
	cfg := GetConfig()
	println(cfg.Url)
	println(config == nil) // false

	cfg = GetConfig()
	println(cfg.Url)
}

type Config struct {
	Port int
	Url  string
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() { // スレッドセーフ
		config = &Config{
			Port: 443,
			Url:  "URL",
		}
		time.Sleep(1 * time.Second)
	})
	return config
}
