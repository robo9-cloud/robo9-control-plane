/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/serverless-cloud-robot/robo9/internal/controller"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Info("starting api web service. v0.0.1 ...")
	controller.StartServer()

}
