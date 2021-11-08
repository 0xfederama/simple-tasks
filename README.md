# Simple Tasks

[![Build Status](https://travis-ci.com/0xfederama/simple-tasks.svg?branch=master)](https://travis-ci.com/0xfederama/water-reminder) [![Go Report Card](https://goreportcard.com/badge/github.com/0xfederama/simple-tasks?)](https://goreportcard.com/report/github.com/0xfederama/simple-tasks?) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

:hammer_and_wrench: :hammer_and_wrench: :hammer_and_wrench:

Manage webcam :camera: and microphone :microphone: with this simple taskbar app.

![Simple-Tasks](https://github.com/0xfederama/simple-tasks/blob/master/.screenshots/tasks.png)

# Download and compile from source

- Download Go 1.16.3 [here](https://golang.org/dl/)
- Clone the repository and cd into it
  ```
  git clone https://github.com/0xfederama/simple-tasks
  cd simple-tasks
  ```
- Then you can execute the code in two ways:
  - Run directly with Go
    ```
    go run .
    ```
  - Compile and run the executable
    ```
    go build -o simple-tasks
    ./simpe-tasks
    ```

# Installation

### Linux

- Download the latest [release](https://github.com/0xfederama/simple-tasks/releases) (a binary file) and place it wherever you want
- To make the app runnable from the applications grid, first launch it from the terminal, then you need to create `simple-tasks.desktop` file in `$HOME/.local/share/applications` and copy this, changing [your username]
```
  [Desktop Entry]
  Name=Simple Tasks
  Exec=/path/to/tasks
  Terminal=false
  Type=Application
  Comment=Manage microphone and webcam
  Icon=/home/[your username]/.config/simple-tasks/tasks.png
```
- To make the app run at startup (using `simple-tasks.desktop` and if you have `gnome-tweak-tools` installed) you can open Tweaks and add Water Reminder to the startup applications. Otherwise, if you didn't create `simple-tasks.desktop` or if you don't have `gnome-tweak-tools`, open "Startup Applications", press "Add" and in the command section type `path/to/tasks`. Give the app the name you want
