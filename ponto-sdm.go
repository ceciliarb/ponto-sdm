package main

import (
  "flag"
  "fmt"
  "os"
  "os/user"
)

func printUsage() {
   fmt.Printf("--------- Exemplo de uso do programa ----------------\n")
   fmt.Printf("./ponto-sdm -u username -p senha\n")
   flag.PrintDefaults()  // prints default usage
   fmt.Printf("-----------------------------------------------------\n")
}

func readArgs() (uname, pass string) {
  // variables declaration  
  var defaultUname string
  defaultUser, err := user.Current()
  if err != nil {
    panic(err)
  }
  defaultUname = defaultUser.Username

  // flags declaration using flag package
  flag.StringVar(&uname, "u", defaultUname, "Especifique um username." )
  flag.StringVar(&pass, "p", "-", "Especifique uma senha.")

  flag.Usage = printUsage

  flag.Parse()  // after declaring flags we need to call it

  return
}


func main() {
  fmt.Println("Olá, mundo!")
  fmt.Println(os.Args[0])
  u, p := readArgs()
  fmt.Printf("user: %s | pass: %s\n", u, p)
}

