package main

import (
	"fmt"
	"os"
	"time"
)

const (
	wsUrl = "https://servicedesk.pbh.gov.br./axis/services/USD_R11_WebService"
)

var usu, pass, server, action, idTicket, refNumTicket, dirConf string
var logFile *os.File

// Novo:        OP   | crs:5200
// Em_execucao: WIP  | crs:5208
// Paralisado:  PARL | crs:400053
// Resolvido:   RE   | crs:5212
var status = map[string]string{"Novo": "crs:5200", "Em_Execucao": "crs:5208", "Paralisado": "crs:400053", "Resolvido": "crs:5212"}

func abrirJornada(handle, handleForUserid string) string {
	now := time.Now()
	data := fmt.Sprintf("%02d/%02d/%04d", now.Day(), now.Month(), now.Year())
	dthr := fmt.Sprintf("%02d/%02d/%04d %02d:%02d:%02d", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())
	return doCreateRequestSdm(handle, handleForUserid, fmt.Sprintf("Início da jornada: %s", dthr), "WIP", fmt.Sprintf("[Registro de ponto] %s", data))
}
func paralisarJornada(handle, objHandle, dataHora string) string {
	return changeStatusSdm(handle, objHandle, fmt.Sprintf("Paralisando jornada para almoço: %s", dataHora), status["Paralisado"])
}
func retomarJornada(handle, objHandle, dataHora string) string {
	return changeStatusSdm(handle, objHandle, fmt.Sprintf("Retornando do almoço: %s", dataHora), status["Em_Execucao"])
}
func finalizarJornada(handle, objHandle, dataHora string) string {
	return changeStatusSdm(handle, objHandle, fmt.Sprintf("Finalizando a jornada: %s", dataHora), status["Resolvido"])
}

func getObjHandle(handle string) string {
	var objHandle string
	if action != "a" && action != "A" && action != "abrir" {
		if idTicket == "" {
			if refNumTicket != "" {
				idTicket = doGetIdTicketByRefNumSdm(handle, refNumTicket)
			} else {
				ticket, err := os.ReadFile(fmt.Sprintf("%s/.idTicket", dirConf))
				idTicket = fmt.Sprintf("%s", ticket)
				if err != nil {
					fmt.Println("É necessário enviar um idTicket, seja através da flag '-t', da flag '-rnt', ou através do arquivo .idTicket .")
					os.Exit(1)
				}
			}
		}
	}
	objHandle = fmt.Sprintf("cr:%s", idTicket)
	return objHandle
}

/****************************** MAIN **********************************************/
func main() {
	usu, pass, server, action, idTicket, refNumTicket, logFile = readArgs()
	now := time.Now()
	dataHora := fmt.Sprintf("%02d/%02d/%04d %02d:%02d:%02d", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())
	if pass == "" || pass == "-" {
		fmt.Println("Password obrigatória para a execução.")
		os.Exit(1)
	}

	handle := doLoginSdm(usu, pass)
	if handle == "" {
		fmt.Println("Login sem sucesso.")
		os.Exit(1)
	}
	objHandle := getObjHandle(handle)
	handleForUserid := doGetHandleForUseridSdm(usu, handle)
	if handleForUserid == "" {
		fmt.Println("handleForUserid sem sucesso.")
		os.Exit(1)
	}

	switch action {
	case "a", "A", "abrir":
		idTicket = abrirJornada(handle, handleForUserid)
		ticketFile, _ := os.OpenFile(fmt.Sprintf("%s/.idTicket", dirConf), os.O_RDWR|os.O_CREATE, 0666)
		ticketFile.WriteString(idTicket)

	case "p", "P", "paralisar":
		paralisarJornada(handle, objHandle, dataHora)
	case "r", "R", "retomar":
		retomarJornada(handle, objHandle, dataHora)
	case "f", "F", "finalizar":
		finalizarJornada(handle, objHandle, dataHora)
	}

	doLogoutSdm(handle)

	fmt.Printf("%s", dataHora)
	fmt.Printf("\nlogin handle (sid): %s", handle)
	fmt.Printf("\nuser handle (handleForUserid): %s", handleForUserid)
	fmt.Printf("\nticket handle (objHandle): %s", objHandle)
	fmt.Printf("\nticket id: %s", idTicket)
	fmt.Printf("\nticket ref_num: %s", refNumTicket)
	fmt.Println()
	if logFile != nil {
		logFile.Close()
	}
}
