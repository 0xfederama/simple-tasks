package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gen2brain/beeep"
	"github.com/shurcooL/trayhost"
)

func createTray(icon string) []trayhost.MenuItem {
	menuItems := []trayhost.MenuItem{
		{
			Title: "Webcam - enable",
			Handler: func() {
				cmd := exec.Command("sudo", "modprobe", "uvcvideo")
				cmd.Run()
				beeep.Alert("Simple tasks", "Webcam turned on", icon)
			},
		},
		{
			Title: "Webcam - disable",
			Handler: func() {
				cmd := exec.Command("sudo", "modprobe", "-r", "uvcvideo")
				cmd.Run()
				beeep.Alert("Simple tasks", "Webcam turned off", icon)
			},
		},
		//TODO: Handler function to enable/disable micrphone
		/*
		{
			Title: "Microphone - enable",
			Handler: func() {

				beeep.Alert("Simple tasks", "Microphone turned on", icon)
			},
		},
		{
			Title: "Microphone - disable",
			Handler: func() {

				beeep.Alert("Simple tasks", "Microphone turned off", icon)
			},
		},
		*/
		trayhost.SeparatorMenuItem(),
		{
			Title:   "Quit",
			Handler: trayhost.Exit,
		},
	}
	return menuItems
}

func findConfig(config string) bool {
	if _, err := os.Stat(filepath.Join(config, "simple-tasks")); os.IsNotExist(err) {
		return false
	}
	return true
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
