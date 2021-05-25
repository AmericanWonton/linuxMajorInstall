# linuxMajorInstall

Secure copy these into any server/node you need to for the following:

- Every reboot will have sudo apt-update run

- You can install multiple packages needed in Ubuntu, (CURRENTLY DESIGNED FOR UBUNTU 20).
This program take CLI arguments, listed here, for differnt installs based on your need:

all - Just installs everything

basic - a basic linux-buntu setup that will set you up with what you need

dev-machine - installs programs for a dev machine I might be working on,
(please note that some programs will still need additional setup)


Notes: brew really can't be run with this, needs to be ran as a user on the system!
Also, Git LFS Needs an executable
git lfs isn't working either...

Might want to run this binary as sudo, just so the program dosen't have to keep asking permission

# For the scripting session:
Please see the notes here if you have any questions on how it's built: https://www.shellscript.sh/