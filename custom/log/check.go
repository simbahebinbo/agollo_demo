package main

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)


func main() {
	c := &config.AppConfig{
		AppID:          "testApplication_yang",
		Cluster:        "dev",
		IP:             "http://106.54.227.205:8080",
		NamespaceName:  "testyml.yml",
		IsBackupConfig: false,
		Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
	}

	loggerInterface:=initSeeLog("seelog.xml")
	agollo.SetLogger(&DefaultLogger{loggerInterface})

	client,err:=agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err!=nil{
		fmt.Println("err:", err)
		panic(err)
	}

	checkKey(c.NamespaceName,client)
}


func checkKey(namespace string,client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	count:=0
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		count++
		return true
	})
	if count<1{
		panic("config key can not be null")
	}
}

type DefaultLogger struct {
	log seelog.LoggerInterface
}

func (this *DefaultLogger) Debugf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Infof(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Warnf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Errorf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Debug(v ...interface{}) {
	this.log.Debug(v)
}
func (this *DefaultLogger) Info(v ...interface{}) {
	this.Debug(v)
}

func (this *DefaultLogger) Warn(v ...interface{}) {
	this.Debug(v)
}

func (this *DefaultLogger) Error(v ...interface{}) {
	this.Debug(v)
}

func initSeeLog(configPath string) seelog.LoggerInterface {
	logger, err := seelog.LoggerFromConfigAsFile(configPath)

	//if error is happen change to default config.
	if err != nil {
		logger, err = seelog.LoggerFromConfigAsBytes([]byte("<seelog />"))
	}

	logger.SetAdditionalStackDepth(1)
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	return logger
}