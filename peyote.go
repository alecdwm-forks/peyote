package main

import "fmt"
import "flag"
import "log"
import "os"
import "strings"

func main(){
  username := os.Getenv("USER")
  hostname := "z"
  dir, err := os.Getwd()
  homedir := os.Getenv("HOME")
  dirname := strings.Replace(dir, homedir, "~", 1)
  if err != nil {
    log.Fatal(err)

  }
  flag.Parse()
  fmt.Printf("%%F{254}%%K{23} %s@%s %%K{241}%%F{23}%%F{254}%%K{241} %s %%f%%k%%F{241}%%f%%k ", username, hostname, dirname)
}


