package database

import (
	"fmt"
	model2 "github.com/Yegor-own/Chat/src/v0/internal/domain/model"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Driver struct{}

func NewDriver() Driver {
	return Driver{}
}

type confYml struct {
	Host     string `yaml:"host" json:"host"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Dbname   string `yaml:"dbname" json:"dbname"`
	Port     string `yaml:"port" json:"port"`
	Sslmode  string `yaml:"sslmode" json:"sslmode"`
	TimeZone string `yaml:"TimeZone" json:"TimeZone"`
}

func (c *confYml) getConf() error {
	absPath, _ := filepath.Abs("../configs/conf-db.yml")
	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		return err
	}

	return nil
}

func (d Driver) ConnectDB() *gorm.DB {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	var yml confYml
	err := yml.getConf()
	if err != nil {
		log.Fatalln(err)
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		yml.Host,
		yml.User,
		yml.Password,
		yml.Dbname,
		yml.Port,
		yml.Sslmode,
		yml.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func migrateModels(db *gorm.DB) error {
	return db.AutoMigrate(&model2.User{}, &model2.Message{}, &model2.Chat{})
}
