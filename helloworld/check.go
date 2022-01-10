package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"time"
)

func main() {
	c := &config.AppConfig{
		AppID:         "zkdex-explorer",
		Cluster:       "test-1",
		IP:            "http://apollo-config.system-service.huobiapps.com",
		NamespaceName: "application.properties",
	}
	agollo.SetLogger(&DefaultLogger{})

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}

	printKey(c.NamespaceName, client)

	c = &config.AppConfig{
		AppID:         "zkdex-explorer",
		Cluster:       "test-1",
		IP:            "http://apollo-config.system-service.huobiapps.com",
		NamespaceName: "zkdex-explorer-common.properties",
	}

	client, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}

	printKey(c.NamespaceName, client)

	time.Sleep(5 * time.Second)
}

func printKey(namespace string, client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		return true
	})
}

type DefaultLogger struct {
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
	fmt.Println(v)
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
