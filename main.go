// This document is the entry of the project.
package main

import (
	"Cerebral-Palsy-Detection-System/Apps"
	"Cerebral-Palsy-Detection-System/Apps/controller/WS/Conf"
	"Cerebral-Palsy-Detection-System/Apps/controller/WS/service"
	"Cerebral-Palsy-Detection-System/Database"
)

func main() {
	Init()
	go service.Manager.Start()
	Apps.InitWebFrameWork()
	Apps.StartServer()
}

func Init() {
	Database.DatabaseInit()
	Conf.Init()
}
