package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

var whatOS string

func init() {
	//Check to see what OS we are running
	whatIsOS()
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
	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	out, err := exec.Command("ls").Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	// let's try the pwd command herer
	out, err = exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output = string(out[:])
	fmt.Println(output)

	out2, err := exec.Command("ls", "-ltr").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output2 := string(out2[:])
	fmt.Println(output2)
}
