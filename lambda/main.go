package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"

	awsbilling "github.com/tommzn/hdb-datasource-awsbilling"
	core "github.com/tommzn/hdb-datasource-core"
)

func main() {

	processor, err := bootstrap()
	if err != nil {
		panic(err)
	}
	lambda.Start(processor.Handle)
}

// bootstrap loads config and creates a new scheduled collector with a exchangerate datasource.
func bootstrap() (core.S3EventHandler, error) {

	conf := loadConfig()
	secretsManager := newSecretsManager()
	logger := newLogger(conf, secretsManager)
	processor := awsbilling.New(true)
	queue := conf.Get("hdb.queue", config.AsStringPtr("de.tsl.hdb.weather"))
	return core.NewS3EventHandler(*queue, processor, conf, logger), nil
}

// loadConfig from config file.
func loadConfig() config.Config {

	configSource, err := config.NewS3ConfigSourceFromEnv()
	if err != nil {
		panic(err)
	}

	conf, err := configSource.Load()
	if err != nil {
		panic(err)
	}
	return conf
}

// newSecretsManager retruns a new secrets manager from passed config.
func newSecretsManager() secrets.SecretsManager {
	return secrets.NewSecretsManager()
}

// newLogger creates a new logger from  passed config.
func newLogger(conf config.Config, secretsMenager secrets.SecretsManager) log.Logger {
	return log.NewLoggerFromConfig(conf, secretsMenager)
}
