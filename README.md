# pScan
A CLI tool that scans the ports on a host address and tells you which is open and closed. 

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

WIP: How to close open ports