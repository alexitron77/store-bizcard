package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"net/http"

	"crypto/tls"

	"biz.card/api"
	ctrl "biz.card/api/controllers"
	mw "biz.card/api/middleware"
	"biz.card/api/repositories/aws"
	"biz.card/api/repositories/mongo"
	"biz.card/config"
	_ "biz.card/docs"
	log "github.com/sirupsen/logrus"
)

func main() {
	addr := flag.String("addr", ":8080", "Server address")
	flag.Parse()
	path, _ := filepath.Abs("config/env")
	conf := config.LoadConfig(path)

	r := gin.Default()

	r.Use(mw.GinLogMiddleware())

	s3 := aws.AwsInit(conf.Aws.AccessKey, conf.Aws.Secret)

	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})

	db := &mongo.DBConn{
		Url:      conf.Mongo.Url,
		Username: conf.Mongo.Username,
		Password: conf.Mongo.Password,
	}

	dbconn := db.ConnectDB()
	defer dbconn.DB.Disconnect(dbconn.Ctx)

	config := config.NewConfig(dbconn.DB, dbconn.Ctx, logger, s3)
	bizCardRepo := mongo.NewDBRepo(config)
	awsRepo := aws.NewAwsRepo()

	card := ctrl.NewBizcardController(config, bizCardRepo, awsRepo)

	api.Handlers(r, card)

	server := http.Server{
		Addr:      *addr,
		Handler:   r,
		TLSConfig: &tls.Config{},
	}

	err := server.ListenAndServeTLS("./certificate/cert.pem", "./certificate/key.pem")

	if err != nil {
		fmt.Print(err)
	}
}
