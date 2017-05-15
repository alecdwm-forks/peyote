package main

import "fmt"
import "flag"
import "log"
import "os"
import "strings"
import "strconv"

func main(){
  returncode := os.Args[1]
  username := os.Getenv("USER")
  hostname := "acid"
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
  if (returncodeint > 0) {
    fmt.Printf("%%F{254}%%K{23} %s@%s %%K{241}%%F{23}%%F{254}%%K{241} %s %%K{236}%%F{241}%%F{203}%%K{236} %s %%f%%k%%F{236} %%f%%k", username, hostname, dirname, returncode)
  } else {
    fmt.Printf("%%F{254}%%K{23} %s@%s %%K{241}%%F{23}%%F{254}%%K{241} %s %%f%%k%%F{241}%%f%%k ", username, hostname, dirname)
  }
}


