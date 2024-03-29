package env

import (
	"flag"
	"fmt"
	conf "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/infrastructure/env/conf"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

type viperConf struct {
	Path string
}

func NewViperConf(path ...string) *viperConf {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "configFile", ConfigFile, "choose config file.")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				config = ConfigFile
				fmt.Printf("Using default config %v\n", ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("Using default ConfigEnv as APP_CONFIG%v\n", config)
			}
		} else {
			fmt.Printf("You are using the value passed by the -configFile parameter of the command line, the path to config is%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("You are using default path %v\n", config)
	}

	return &viperConf{
		config,
	}
}

func (v viperConf) ReadConfig() error {
	fmt.Println("Loading Env...")
	viper.SetConfigFile(v.Path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Viper load fail: %v\n", err)
		return err
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := viper.Unmarshal(&ENV); err != nil {
			fmt.Println(err)
		}
	})

	if err := viper.Unmarshal(&ENV); err != nil {
		fmt.Println(err)
	}
	return nil
}

func (v viperConf) WriteConfig(conf conf.AppConf) error {
	cs := StructToMap(conf, "mapstructure")
	for k, v := range cs {
		viper.Set(k, v)
	}
	if err := viper.SafeWriteConfigAs(v.Path); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfigAs(v.Path)
			if err != nil {
				log.Fatalf("Viper write config fail: %v", err)
				return err
			}
		} else {
			confPath := filepath.Dir(v.Path)
			viper.AddConfigPath(confPath)
			err = viper.WriteConfig()
			if err != nil {
				log.Fatalf("Viper write config fail: %v", err)
				return err
			}
		}
	}
	return nil
}

func StructToMap(obj interface{}, tagString string) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		if objType.Field(i).Tag.Get(tagString) != "" {
			data[objType.Field(i).Tag.Get(tagString)] = objVal.Field(i).Interface()
		} else {
			data[objType.Field(i).Name] = objVal.Field(i).Interface()
		}
	}
	return data
}
