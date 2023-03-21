package configs

import (
	"github.com/MrMohebi/sar-khati-bot/common"
	"gopkg.in/ini.v1"
)

var isIniInitOnce = false
var IniData *ini.File

func IniSetup() {
	if !isIniInitOnce {
		var err error
		IniData, err = ini.Load("config.ini")
		common.IsErr(err, false, "Error loading .ini file")
		isIniInitOnce = true
	} else {
		println("initialized inis once")
	}
}

func IniGet(section string, key string) string {
	if IniData == nil {
		IniSetup()
	}
	return IniData.Section(section).Key(key).String()
}
