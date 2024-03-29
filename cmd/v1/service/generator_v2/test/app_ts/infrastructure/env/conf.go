package env

import (
	"fmt"
	conf "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/infrastructure/env/conf"
	"log"
	"os"
)

type Conf interface {
	ReadConfig()
	WriteConfig(env2 conf.AppConf) error
}

func LoadConf(path ...string) {
	viper := NewViperConf(path...)
	err := viper.ReadConfig()
	if err != nil {
		log.Fatalf("Viper load fail: %v", err)

		err = viper.WriteConfig(conf.DefaultAppConf)
		if err != nil {
			log.Fatalf("Viper load fail: %v", err)
		}

	}

	loadFromEnv()

	fmt.Println("Reload Config Structure")
	_ = viper.WriteConfig(ENV)
	fmt.Println("--Load env success")
}

func loadFromEnv() {
	envAppMode := os.Getenv("APP_MODE")
	envAppAddress := os.Getenv("APP_ADDRESS")

	if envAppMode == "" && ENV.System.Env == "" {
		ENV.System.Env = AppModeDev
	}

	if envAppAddress == "" {
		envAppAddress = ENV.System.Address
		if envAppAddress == "" {
			ENV.System.Address = ""
		}
	} else {
		ENV.System.Address = envAppAddress
	}
}
