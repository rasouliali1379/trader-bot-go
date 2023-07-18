package config

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	confs = Config{}
	lock  = sync.Mutex{}
)

const (
	EnvProduction = "prod"
	EnvDev        = "dev"
	EnvTest       = "test"
)

type Config struct {
	App   App   `yaml:"app" required:"true"`
	Redis Redis `yaml:"redis" required:"true"`
	Mysql Mysql `yaml:"database" required:"true"`
	OKX   OKX   `yaml:"OKX"`
}

type App struct {
	Name        string `yaml:"app.name" required:"true"`
	Environment string `yaml:"app.environment" required:"true"`
}

type Mysql struct {
	Username string `yaml:"mysql.username" required:"true"`
	Password string `yaml:"mysql.password" required:"true"`
	Host     string `yaml:"mysql.host" required:"true"`
	Schema   string `yaml:"mysql.schema" required:"true"`
	Port     string `yaml:"mysql.port" required:"true"`
}

type Redis struct {
	Username string        `yaml:"redis.username"`
	Password string        `yaml:"redis.password"`
	DB       int           `yaml:"redis.db"`
	Host     string        `yaml:"redis.host" required:"true"`
	Timeout  time.Duration `yaml:"timeout"`
}

type OKX struct {
	Url        string `yaml:"exchanges.okx.url"  required:"true"`
	Origin     string `yaml:"exchanges.okx.origin"  required:"true"`
	ApiKey     string `yaml:"exchanges.okx.apiKey"  required:"true"`
	SecretKey  string `yaml:"exchanges.okx.secretKey"  required:"true"`
	Passphrase string `yaml:"exchanges.okx.passphrase"  required:"true"`
}

func Validate(c any) error {
	errMsg := ""
	numFields := reflect.TypeOf(c).NumField()
	for i := 0; i < numFields; i++ {
		fieldType := reflect.TypeOf(c).Field(i)
		tagval, ok := fieldType.Tag.Lookup("required")
		isRequired := ok && tagval == "true"
		if !isRequired {
			continue
		}
		fieldVal := reflect.ValueOf(c).Field(i)
		if fieldVal.Kind() == reflect.Struct {
			if err := Validate(fieldVal.Interface()); err != nil {
				errMsg += fmt.Sprintf("%s > [%v], ", fieldType.Name, err)
			}
		} else {
			if fieldVal.IsZero() {
				errMsg += fmt.Sprintf("%s, ", fieldType.Name)
			}
		}
	}
	if errMsg == "" {
		return nil
	}
	return errors.New(errMsg)
}

func C() *Config {
	return &confs
}

func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../../dev/config/trader/")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	loadConfigs()
	viper.OnConfigChange(func(in fsnotify.Event) {
		lock.Lock()
		defer lock.Unlock()
		lastUpdate := viper.GetTime("fsnotify")
		if time.Since(lastUpdate) < time.Second {
			return
		}
		viper.Set("fsnotify", time.Now())
		log.Println("config file changed. restarting...")
	})
	viper.WatchConfig()
}

func loadConfigs() {
	must(viper.Unmarshal(&confs),
		"could not unmarshal config file")
	must(Validate(confs),
		"some required configurations are missing")
	log.Printf("configs loaded from file successfully \n")
}

func must(err error, logMsg string) {
	if err != nil {
		log.Println(logMsg)
		panic(err)
	}
}
