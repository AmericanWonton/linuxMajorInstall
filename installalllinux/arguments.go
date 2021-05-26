package main

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
	/* Install ifconfig; not really needed, but nice to have */
	theCommands = append(theCommands, "apt-get", "install", "net-tools")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* snapd/ curl */
	theCommands = append(theCommands, "sudo", "apt", "install", "snapd", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "install", "curl", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Install Binarturals */
	theCommands = append(theCommands, "sudo", "apt", "install", "binutils-for-build")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Install nautilus-admin */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "nautilus-admin")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install leafpad */
	theCommands = append(theCommands, "sudo", "snap", "install", "leafpad")
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
	/* install brew
	Note, this cannot run as root; might need to be determined that this runs manually */
	/*
		theCommands = append(theCommands, "/bin/bash", "-c", "\"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\"", "-y")
		allPackageCommands = append(allPackageCommands, theCommands)
		theCommands = nil
		theCommands = append(theCommands, "echo", "'eval \"$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)\"'", ">>", "/home/joek/.profile")
		allPackageCommands = append(allPackageCommands, theCommands)
		theCommands = nil
		theCommands = append(theCommands, "eval", "\"$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)\"")
		allPackageCommands = append(allPackageCommands, theCommands)
		theCommands = nil
	*/
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
	/* Can't do git lfs for some reason...
	theCommands = append(theCommands, "sudo", "curl", "-s", "https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh",
		"|", "sudo", "bash")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	*/
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
}

func installApps() {
	/* Update and upgrade */
	theCommands = append(theCommands, "sudo", "apt-get", "update", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt-get", "upgrade", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Google Chrome */
	theCommands = append(theCommands, "wget", "https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "install", "./google-chrome-stable_current_amd64.deb")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* discord */
	theCommands = append(theCommands, "sudo", "snap", "install", "discord")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Install VS Code */
	theCommands = append(theCommands, "sudo", "snap", "install", "--classic", "code")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* inkscape */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "inkscape")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* notepad++ */
	theCommands = append(theCommands, "sudo", "snap", "install", "notepad-plus-plus")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* blender */
	theCommands = append(theCommands, "sudo", "snap", "install", "blender", "--classic")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
}

func fillDevMachine() {
	/* Installing Docker */
	/* Start at home directory */
	theCommands = append(theCommands, "cd")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "update", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "install", "apt-transport-https", "ca-certificates", "curl", "software-properties-common")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "curl", "-fsSL", "https://download.docker.com/linux/ubuntu/gpg",
		"|", "sudo", "apt-key", "add", "-")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "add-apt-repository", "\"deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable\"")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "update")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "apt-cache", "policy", "docker-ce")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "install", "docker-ce")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Install Java JDK...MIGHT BE BEST TO USE BREW */

	/* Have a Docker User created at some point */

	/* Install Golang (NOT USED, DO THIS WITH BREW) */
	/* Start at home directory */
	/*
		theCommands = append(theCommands, "cd")
		allPackageCommands = append(allPackageCommands, theCommands)
		theCommands = nil
		theCommands = append(theCommands, "sudo", "apt", "install", "golang-go", "-y")
		allPackageCommands = append(allPackageCommands, theCommands)
		theCommands = nil
	*/
	/* make go working directory */
	/* NOT USED, DO THIS WITH BREW
	theCommands = append(theCommands, "cd", "go")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "pkg")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "src")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "bin")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "cd")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	*/
}
