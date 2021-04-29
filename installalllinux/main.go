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
var whatArg string

func init() {
	//Check to see what OS we are running
	whatIsOS()
}

//Used for logging
func logWriter(logMessage string) {
	//Logging info

	wd, _ := os.Getwd()
	logDir := filepath.Join(wd, "logging", "installlinux.txt")
	//Check to see if path exists
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll("logging", 0777) // Make the directory
		//Create file
		emptyFile, err2 := os.Create("./logging/installlinux.txt")
		if err2 != nil {
			log.Fatal(err2)
		}
		emptyFile.Close()
	}

	logFile, err := os.OpenFile(logDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	defer logFile.Close()

	if err != nil {
		fmt.Println("Failed opening log file: " + err.Error())
	}

	log.SetOutput(logFile)

	log.Println(logMessage)
}

func main() {
	if strings.Contains(whatOS, "linux") {
		//Log what we're doing
		theArgs := os.Args
		fmt.Printf("The args is: %v\n", theArgs)
		if theArgs == nil {
			fmt.Printf("Error, no argument given to program!\n")
			return
		}
		whatArg = theArgs[1]
		theRunner := "We are installing with the following setting: " + whatArg
		fmt.Println(theRunner)
		logWriter(theRunner)
		//Execute linux commands
		executeLinuxCommands(whatArg)
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

func executeLinuxCommands(theArg string) {
	switch theArg {
	case "all":
		//All Linux setup

		break
	case "basic":
		//Basic Linux Setup
		fillBasic()
		break
	default:
		//Error, log it
		theErr := "Error, wrong command line arugment entered: " + theArg
		fmt.Println(theErr)
		logWriter(theErr)
		break
	}
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
	logWriter(result)
}

func commandExecute() {

}
