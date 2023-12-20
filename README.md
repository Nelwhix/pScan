# pScan
A CLI tool that scans the ports on a host address and tells you which is open and closed. 

## Installation
Installation is only available on homebrew for now. Run:
```bash 
    brew tap nelwhix/nelwhix
    brew install pscan 
```

## Commands
To add a new host write
```bash
    ./pScan hosts add [host-address]
```
You can also list out the hosts the CLI is monitoring with:
```bash
    ./pScan hosts list
```
To scan the ports on the hosts run:
```bash
    ./pScan scan -p 8080,8000,443
```
replace 8080 etc. with the ports you want to check if closed

a closed port means it is free for use, an open port means it is in use

To kill a port, run:
```bash 
    ./pScan kill -p 8000
```
the command will ask for confirmation and if you agree the port is closed.
