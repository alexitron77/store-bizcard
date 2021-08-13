package main

import (
	"flag"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"biz.card/cmd/api"
	ctrl "biz.card/cmd/api/controllers"
	mw "biz.card/cmd/api/middleware"
	"biz.card/cmd/api/repositories/aws"
	"biz.card/cmd/api/repositories/mongo"
	"biz.card/config"
	cg "biz.card/config"
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

	config := cg.NewConfig(conf, logger)
	storage := cg.NewStorage(dbconn.DB, s3)
	bizCardRepo := mongo.NewDBRepo(config, storage)
	awsRepo := aws.NewAwsRepo()

	card := ctrl.NewBizcardController(config, storage, bizCardRepo, awsRepo)

	a := api.NewApi(addr, r)
	a.Routes(card)
	a.Start()

}
