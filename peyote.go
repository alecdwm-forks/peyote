package main

import "fmt"
import "flag"
import "log"
import "os"
import "strings"

func main(){
  username := os.Getenv("USER")
  hostname := "z"
  dirname := strings.Replace(dir, homedir, "~", 1)
  homedir := os.Getenv("HOME")
  dir, err := os.Getwd()
  if err != nil {
    log.Fatal(err)

  }
  flag.Parse()
  fmt.Printf("%%F{254}%%K{23} %s@%s %%K{241}%%F{23}%%F{254}%%K{241} %s %%f%%k%%F{241}%%f%%k ", username, hostname, dirname)
}


