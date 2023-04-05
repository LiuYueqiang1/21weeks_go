package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, _ := ini.Load("F:\\goland\\go_project\\21weeks\\21weeks_go\\156_logagent\\conf\\config.ini")
	fmt.Println(cfg.Section("kafka").Key("address"))
}
