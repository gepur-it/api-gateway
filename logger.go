package main

import (
	"github.com/zbindenren/logrus_mail"
	"github.com/sirupsen/logrus"
	"fmt"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal(msg)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func notify(msg string) {
    logger.Info(msg)
}

var logger = logrus.New()

func initLogger(Config Configuration) {
  	hook, _  := logrus_mail.NewMailAuthHook(
		Config.LogEmailAppName,
		Config.LogEmailHost,
		Config.LogEmailPort,
		Config.LogEmailFrom,
		Config.LogEmailTo,
		Config.LogEmailUser,
		Config.LogEmailPassword,
	)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	formatter := new(logrus.TextFormatter)
	formatter.FullTimestamp = true
	logger.SetFormatter(formatter)
	logger.Hooks.Add(hook)
}
