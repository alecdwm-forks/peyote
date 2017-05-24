package main

import "fmt"
import "flag"
import "log"
import "os"
import "strings"
import "strconv"

func main() {
	returncode := os.Args[1]
	username := os.Getenv("USER")
	euid := os.Args[2]
	euidnum, err := strconv.Atoi(euid)
	if err != nil {
		log.Fatal(err)
	}
	usercolor := "1"
	if euidnum > 0 {
		usercolor = "23"
	} else {
		usercolor = "88"
	}
	hostname, err := os.Hostname()
	dir, err := os.Getwd()
	homedir := os.Getenv("HOME")
	dirname := strings.Replace(dir, homedir, "~", 1)
	if err != nil {
		log.Fatal(err)

	}
	flag.Parse()
	returncodeint, err := strconv.Atoi(returncode)
	if err != nil {
		log.Fatal(err)
	}
	if returncodeint > 0 {
		fmt.Printf("%%F{254}%%K{%s} %s@%s %%K{241}%%F{%s}%%F{254}%%K{241} %s %%K{236}%%F{241}%%F{203}%%K{236} %s %%f%%k%%F{236} %%f%%k", usercolor, username, hostname, usercolor, dirname, returncode)
	} else {
		fmt.Printf("%%F{254}%%K{%s} %s@%s %%K{241}%%F{%s}%%F{254}%%K{241} %s %%f%%k%%F{241}%%f%%k ", usercolor, username, hostname, usercolor, dirname)
	}
}
