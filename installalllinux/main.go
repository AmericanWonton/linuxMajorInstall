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
		/*
			theArgs := os.Args
			fmt.Printf("The args is: %v\n", theArgs)
			if theArgs == nil {
				fmt.Printf("Error, no argument given to program!\n")
				return
			}
			whatArg = theArgs[1]
		*/
		//Take argument from User
		fmt.Printf("Please enter what install you are performing: ")
		var whatInstall string
		fmt.Scanln(&whatInstall)
		fmt.Println()
		theRunner := "We are installing with the following setting: " + whatInstall
		fmt.Println(theRunner)
		logWriter(theRunner)
		//Execute linux commands
		executeLinuxCommands(whatInstall)
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
		fmt.Printf("DEBUG: Installing all on linux\n")
		fillBasic()
		fillDevMachine()
		goodUpdate, theResult := commandExecute(allPackageCommands)
		if !goodUpdate {
			bigResult := ""
			for j := 0; j < len(theResult); j++ {
				bigResult = bigResult + theResult[j] + "\n"
			}
			logWriter(bigResult)
		} else {
			theTimeNow := time.Now()
			result := "This ran succesffully on " + theTimeNow.Format("2006-01-02 15:04:05")
			fmt.Println(result)
			logWriter(result)
		}
		break
	case "basic":
		//Basic Linux Setup
		fmt.Printf("DEBUG: Running basic linux setup\n")
		fillBasic()
		goodUpdate, theResult := commandExecute(allPackageCommands)
		if !goodUpdate {
			bigResult := ""
			for j := 0; j < len(theResult); j++ {
				bigResult = bigResult + theResult[j] + "\n"
			}
			logWriter(bigResult)
		} else {
			theTimeNow := time.Now()
			result := "This ran succesffully on " + theTimeNow.Format("2006-01-02 15:04:05")
			fmt.Println(result)
			logWriter(result)
		}
		break
	case "apps":
		//Basic app insatlls
		fmt.Printf("DEBUG: Running app linux setup\n")
		installApps()
		goodUpdate, theResult := commandExecute(allPackageCommands)
		if !goodUpdate {
			bigResult := ""
			for j := 0; j < len(theResult); j++ {
				bigResult = bigResult + theResult[j] + "\n"
			}
			logWriter(bigResult)
		} else {
			theTimeNow := time.Now()
			result := "This ran succesffully on " + theTimeNow.Format("2006-01-02 15:04:05")
			fmt.Println(result)
			logWriter(result)
		}
		break
	case "dev-machine":
		fmt.Printf("Installing on Dev Machine\n")
		fillBasic()
		fillDevMachine()
		goodUpdate, theResult := commandExecute(allPackageCommands)
		if !goodUpdate {
			bigResult := ""
			for j := 0; j < len(theResult); j++ {
				bigResult = bigResult + theResult[j] + "\n"
			}
			logWriter(bigResult)
		} else {
			theTimeNow := time.Now()
			result := "This ran succesffully on " + theTimeNow.Format("2006-01-02 15:04:05")
			fmt.Println(result)
			logWriter(result)
		}
		break
	default:
		//Error, log it
		theErr := "Error, wrong command line arugment entered: " + theArg
		fmt.Println(theErr)
		logWriter(theErr)
		break
	}
}

func commandExecute(packgeCommands [][]string) (bool, []string) {
	goodUpdate, theResult := true, []string{}

	//For length of inserted package commands, run the commands
	for x := 0; x < len(packgeCommands); x++ {
		fmt.Printf("DEBUG: Nubmered Roundabout roundabout: %v\n", x)
		commandFirstWord := packgeCommands[x][0] //Insert first package command into argument
		restOfCommands := packgeCommands[x][1:]  //Take out the initial argument to pass in the rest
		fmt.Printf("DEBUG: restofCommands = %v\n", restOfCommands)
		_, err := exec.Command(commandFirstWord, restOfCommands...).Output()
		if err != nil {
			errorComp := "Errored during the following installation:\n"
			for z := 0; z < len(restOfCommands); z++ {
				errorComp = errorComp + " " + restOfCommands[z]
			}
			errorComp = errorComp + "\n" + err.Error()
			fmt.Println(errorComp)
			logWriter(errorComp)
			goodUpdate = false
			theResult = append(theResult, errorComp)
			break
		}
	}

	return goodUpdate, theResult
}
