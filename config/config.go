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
	"go.uber.org/fx"
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

func Validate(c any) error {
	errmsg := ""
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
				errmsg += fmt.Sprintf("%s > [%v], ", fieldType.Name, err)
			}
		} else {
			if fieldVal.IsZero() {
				errmsg += fmt.Sprintf("%s, ", fieldType.Name)
			}
		}
	}
	if errmsg == "" {
		return nil
	}
	return errors.New(errmsg)
}

func C() *Config {
	return &confs
}

func Init(shutdowner fx.Shutdowner, configPath string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
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
		shutdowner.Shutdown()
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
