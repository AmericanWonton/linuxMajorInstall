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
	theCommands = append(theCommands, "sudo", "apt", "install", "binutils-for-build", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Install nautilus-admin */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "nautilus-admin", "-y")
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
	/* Build essential */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "build-essential", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install pvscan */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "lvm2", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Need update for below programs */
	theCommands = append(theCommands, "sudo", "apt-get", "update", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt-get", "upgrade", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Git Desktop Core */
	theCommands = append(theCommands, "sudo", "apt", "install", "git-all", "-y")
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
	/* Install rpm */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "rpm", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Install alien */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "alien", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install apache webserver */
	theCommands = append(theCommands, "sudo", "apt", "install", "certbot", "python3-certbot-apache", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "search","libapache2-mod-")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "install", "--assume-yes", "libapache2-mod-security2")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* allow apache traffic */
	theCommands = append(theCommands, "sudo", "ufw", "allow", "'Apache Full'")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "delete","allow", "'Apache'")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* enable security features */
	theCommands = append(theCommands, "sudo", "a2enmod", "proxy")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "a2enmod", "proxy_http")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "a2enmod", "proxy_balancer")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "a2enmod", "lbmethod_byrequests")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* restart apache2 */
	theCommands = append(theCommands, "sudo", "systemctl", "restart","apache2")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Open known ports I work with */
	theCommands = append(theCommands, "sudo", "ufw", "allow","22")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","443")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","8080")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","8000")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","ssh")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","http")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","https")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","80")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","3000")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install golang */
	theCommands = append(theCommands, "sudo", "snap", "install","go", "--classic")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* make directories */
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/bin")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/src")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/pkg")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/src/git_repositories")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/src/other")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* give  exeucte permissions */
	theCommands = append(theCommands, "chmod", "777", "startUpCronJob")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* set the go path */
	theCommands = append(theCommands, "echo", "\"export GOPATH=$HOME/startUpCronJob/go\"", ">>", "~/.bashrc")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "echo", "\"export PATH=$PATH:$GOROOT/bin:$GOPATH/bin\"", ">>", "~/.bashrc")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
}

func fillnewServer() {
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
	theCommands = append(theCommands, "sudo", "apt", "install", "binutils-for-build", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Install nautilus-admin */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "nautilus-admin", "-y")
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
	/* Build essential */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "build-essential", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install pvscan */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "lvm2", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Need update for below programs */
	theCommands = append(theCommands, "sudo", "apt-get", "update", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt-get", "upgrade", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Git Desktop Core */
	theCommands = append(theCommands, "sudo", "apt", "install", "git-all", "-y")
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
	/* Install rpm */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "rpm", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Install alien */
	theCommands = append(theCommands, "sudo", "apt-get", "install", "alien", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install apache webserver */
	theCommands = append(theCommands, "sudo", "apt", "install", "certbot", "python3-certbot-apache", "-y")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "search","libapache2-mod-")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "apt", "install", "--assume-yes", "libapache2-mod-security2")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* allow apache traffic */
	theCommands = append(theCommands, "sudo", "ufw", "allow", "'Apache Full'")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "delete","allow", "'Apache'")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* enable security features */
	theCommands = append(theCommands, "sudo", "a2enmod", "proxy")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "a2enmod", "proxy_http")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "a2enmod", "proxy_balancer")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "a2enmod", "lbmethod_byrequests")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* restart apache2 */
	theCommands = append(theCommands, "sudo", "systemctl", "restart","apache2")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* create example files to work with */
	theCommands = append(theCommands, "touch", "/etc/apache2/sites-available/josephkeller.me.conf")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "touch", "/etc/apache2/sites-available/mariahsplayroom.com.conf")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "chmod", "777", "/etc/apache2/sites-available/josephkeller.me.conf")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "chmod", "777", "/etc/apache2/sites-available/mariahsplayroom.com.conf")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Open known ports I work with */
	theCommands = append(theCommands, "sudo", "ufw", "allow","22")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","443")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","8080")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","8000")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","ssh")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","http")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","https")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","80")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "sudo", "ufw", "allow","3000")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* install golang */
	theCommands = append(theCommands, "sudo", "snap", "install","go", "--classic")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* make directories */
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/bin")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/src")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/pkg")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/src/git_repositories")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "mkdir", "startUpCronJob/go/src/other")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* give  exeucte permissions */
	theCommands = append(theCommands, "chmod", "777", "startUpCronJob")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* set the go path */
	theCommands = append(theCommands, "echo", "\"export GOPATH=$HOME/startUpCronJob/go\"", ">>", "~/.bashrc")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	theCommands = append(theCommands, "echo", "\"export PATH=$PATH:$GOROOT/bin:$GOPATH/bin\"", ">>", "~/.bashrc")
	allPackageCommands = append(allPackageCommands, theCommands)
	theCommands = nil
	/* Install Docker */
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
