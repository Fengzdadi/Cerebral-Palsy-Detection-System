package main

import (
	"Cerebral-Palsy-Detection-System/Database"
	"Cerebral-Palsy-Detection-System/WS/Conf"
)

func Init() {
	Database.DatabaseInit()
	Conf.Init()
}
