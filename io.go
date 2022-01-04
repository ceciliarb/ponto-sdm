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
	fmt.Printf("--------- Exemplo de uso do programa ----------------------------------------------------------------\n")
	fmt.Printf("./ponto-sdm -u username -p senha\n")
	fmt.Printf("./ponto-sdm --username username --password senha --log\n")
	fmt.Printf("./ponto-sdm --u username --p senha --l --action a \n")
	fmt.Printf("./ponto-sdm --u username --p senha --l --action abrir \n")
	fmt.Printf("./ponto-sdm --u username --p senha --l --action p \n")
	fmt.Printf("./ponto-sdm --u username --p senha --l --action paralisar \n")
	fmt.Printf("./ponto-sdm --u username --p senha --l --action r \n")
	fmt.Printf("./ponto-sdm --u username --p senha --l --action retomar \n")
	fmt.Printf("./ponto-sdm --u username --p senha --l --action f \n")
	fmt.Printf("./ponto-sdm --u username --p senha --l --action F --id-ticket 2455454 \n")
	fmt.Printf("./ponto-sdm --u username --p senha --l --action finalizar --num-ref-ticket 934850 \n")
	fmt.Printf("\n")
	fmt.Printf("------------------ Parametros -----------------------------------------------------------------------\n")
	flag.PrintDefaults() // prints default usage
}

func readArgs() (uname, pass, server, action, idTicket, refNumTicket string, logFile *os.File) {
	// variables declaration
	var bLog bool
	var defaultUname string
	defaultUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	defaultUname = defaultUser.Username

	// flags declaration using flag package
	flag.StringVar(&uname, "username", defaultUname, "Especifique um username.")
	flag.StringVar(&uname, "u", uname, "alias para -username")
	flag.StringVar(&pass, "password", "-", "Especifique uma senha.")
	flag.StringVar(&pass, "p", pass, "alias para -password")
	flag.StringVar(&server, "server", wsUrl, "Especifique um servidor.")
	flag.StringVar(&server, "s", server, "alias para -server.")
	flag.StringVar(&action, "action", "abrir", "Especifique uma ação ('abrir', 'paralisar', 'retormar', 'finalizar').")
	flag.StringVar(&action, "a", action, "alias para -action.")
	flag.StringVar(&idTicket, "id-ticket", "", "Especifique um id ticket (talvez exista um no arquivo .idTicket).")
	flag.StringVar(&idTicket, "t", idTicket, "alias para -id-ticket.")
	flag.StringVar(&refNumTicket, "ref-num-ticket", "", "Especifique um ref_num do ticket (número do ticket no SDM).")
	flag.StringVar(&refNumTicket, "rnt", refNumTicket, "alias para -ref-num-ticket.")
	flag.BoolVar(&bLog, "log", false, "Se estiver presente, armazena log das operações.")
	flag.BoolVar(&bLog, "l", bLog, "alias para -log.")

	flag.Usage = printUsage

	flag.Parse() // after declaring flags we need to call it

	dirConf = fmt.Sprintf("/home/%s/.ponto-sdm/", defaultUname)
	_, err = os.Stat(dirConf)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirConf, 0755)
		if err != nil {
			dirConf = "."
		}
	}

	if bLog {
		logFile, err = os.OpenFile(fmt.Sprintf("%s/sdm.log", dirConf), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
