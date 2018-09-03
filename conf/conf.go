package conf

import "github.com/bigzhu/gobz/confbz"

// Seed 配置文件
type Seed struct {
	Seed string `toml:"seed"`
}

// GetSeedConf seed 配置
func GetSeedConf() (seed Seed) {
	confbz.GetConf("seed", &seed)
	return
}
