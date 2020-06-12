package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/shurcooL/trayhost"
)

func main() {

	//Search in .config dir if there is the directory simple-tasks
	home, _ := os.LookupEnv("HOME")
	configPath := filepath.Join(home, ".config")
	configDirPath := filepath.Join(configPath, "simple-tasks")
	configIconPath := filepath.Join(configDirPath, "tasks.png")

	if !findConfig(configPath) {
		//Create config directory
		os.Mkdir(configDirPath, 0700)

		//Download icon in the new directory
		downloadFile("https://raw.githubusercontent.com/0xfederama/simple-tasks/master/resources/tasks.png", configIconPath)
	}

	//Create tray app
	trayMenu := createTray(configIconPath)
	//Load tray icon
	iconData, err := ioutil.ReadFile(configIconPath)
	if err != nil {
		return
	}
	trayhost.Initialize("Simple Tasks", iconData, trayMenu)
	trayhost.EnterLoop()
}
