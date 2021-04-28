package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var whatOS string

func init() {
	//Check to see what OS we are running
	whatIsOS()
}

//Used for logging
func logWriter(logMessage string) {
	//Logging info

	wd, _ := os.Getwd()
	logDir := filepath.Join(wd, "logging", "biglinuxstartup.txt")
	logFile, err := os.OpenFile(logDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

	defer logFile.Close()

	if err != nil {
		fmt.Println("Failed opening log file")
	}

	log.SetOutput(logFile)

	log.Println(logMessage)
}

func main() {
	if strings.Contains(whatOS, "linux") {
		//Execute linux commands
		executeLinuxCommands()
	} else {
		fmt.Printf("Wrong, you're running this on windows")
	}
}

func whatIsOS() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		whatOS = runtime.GOOS
	}
}

func executeLinuxCommands() {
	//update and upgrade
	_, err := exec.Command("sudo", "apt-get", "update", "-y").Output()
	if err != nil {
		theErr := "Errored during update: " + err.Error()
		fmt.Println(theErr)
		logWriter(theErr)
	}
	_, err2 := exec.Command("sudo", "apt-get", "upgrade", "-y").Output()
	if err2 != nil {
		theErr := "Errored during upgrade: " + err2.Error()
		fmt.Println(theErr)
		logWriter(theErr)
	}

	theTimeNow := time.Now()

	result := "This ran succesffully on " + theTimeNow.Format("2006-01-02 15:04:05")
	fmt.Println(result)

}
