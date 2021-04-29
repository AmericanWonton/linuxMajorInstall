package main

import "fmt"

var theCommands []string
var allPackageCommands [][]string

func fillBasic() {
	/* Update and upgrade */
	theCommands = append(theCommands, "sudo", "apt-get", "update", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt-get", "upgrade", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* snapd/ curl */
	theCommands = append(theCommands, "sudo", "apt", "install", "snapd", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "install", "curl", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install leafpad */
	theCommands = append(theCommands, "sudo", "snap", "install", "leafpad", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install gimp */
	theCommands = append(theCommands, "sudo", "apt", "install", "gimp", "gimp-gmic", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* lsof/ ufw*/
	theCommands = append(theCommands, "sudo", "apt", "install", "lsof", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "install", "ufw", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install windows xrp */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "xrdp", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install brew */
	theCommands = append(theCommands, "/bin/bash", "-c", "\"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\"", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "echo", "'eval \"$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)\"'", ">>", "/home/joek/.profile")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "eval", "\"$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)\"")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* update for final section of brew */
	theCommands = append(theCommands, "sudo", "apt-get", "update", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt-get", "upgrade", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt-get", "install", "build-essential", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install git/github desktop core */
	theCommands = append(theCommands, "sudo", "apt", "install", "git-all", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "git", "lfs", "install", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "install", "git-review", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt-get", "install", "gdebi-core", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "wget", "https://github.com/shiftkey/desktop/releases/download/release-2.6.3-linux1/GitHubDesktop-linux-2.6.3-linux1.deb")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "gdebi", "GitHubDesktop-linux-2.6.3-linux1.deb")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil

	fmt.Printf("All package commands is: %v\n", allPackageCommands)
}
