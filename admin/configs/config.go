package configs

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"time"
)

type DbConfig struct {
	MyBlog *ORMConfig
}

type ORMConfig struct {
	DSN         string
	Active      int
	Idle        int
	IdleTimeout time.Duration
}

type ServerConfig struct {
	Port    string
	Timeout string
}

type Configs struct {
	Db     *DbConfig
	Server *ServerConfig
}

func NewDbConfig() *DbConfig {
	dc := &DbConfig{}
	if _, err := toml.DecodeFile("D:\\GoProjects\\src\\my_blog\\admin\\configs\\db.toml", dc); err != nil {
		log.Error("read db toml false, err [%v]", err)
		panic(err)
	}
	return dc
}

func NewServerConfig() *ServerConfig {
	sc := &ServerConfig{}
	if _, err := toml.DecodeFile("D:\\GoProjects\\src\\my_blog\\admin\\configs\\http.toml", sc); err != nil {
		log.Error("read http toml false, err [%v]", err)
		panic(err)
	}
	return sc
}

func NewConfigs(Db *DbConfig, Sc *ServerConfig) *Configs {
	c := &Configs{
		Db:     Db,
		Server: Sc,
	}
	return c
}
