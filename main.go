/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/serverless-cloud-robot/robo9/internal/controller"
)

func main() {
	fmt.Println("starting api web service. v0.0.1 ...")
	controller.StartServer()

}
