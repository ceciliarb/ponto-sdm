package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"time"
)

func printUsage() {
	fmt.Printf("--------- Exemplo de uso do programa ----------------\n")
	fmt.Printf("./ponto-sdm -u username -p senha\n")
	flag.PrintDefaults() // prints default usage
	fmt.Printf("-----------------------------------------------------\n")
}

func readArgs() (uname, pass, server, action, ticket string, logFile *os.File) {
	// variables declaration
	var bLog bool
	var defaultUname string
	defaultUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	defaultUname = defaultUser.Username

	// flags declaration using flag package
	flag.StringVar(&uname, "u", defaultUname, "Especifique um username.")
	flag.StringVar(&pass, "p", "-", "Especifique uma senha.")
	flag.StringVar(&server, "s", wsUrl, "Especifique um servidor.")
	flag.StringVar(&action, "a", "abrir", "Especifique uma ação ('abrir', 'paralisar', 'retormar', 'finalizar').")
	flag.StringVar(&ticket, "t", "", "Especifique um id ticket (talvez exista um no arquivo .idTicket).")
	flag.BoolVar(&bLog, "l", false, "Se estiver presente, armazena log das operações.")

	flag.Usage = printUsage

	flag.Parse() // after declaring flags we need to call it

	if bLog {
		logFile, err = os.OpenFile("sdm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		} else {
			now := time.Now()
			logFile.WriteString(fmt.Sprintf("\nLOG: %d/%d/%d %d:%d\n", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute()))
		}
	} else {
		logFile = nil
	}

	return
}
