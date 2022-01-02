package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/antchfx/xmlquery"
)

type loginRequest struct {
	Username string
	Password string
}

type getHandleForUseridRequest struct {
	Handle   string
	Username string
}

type createRequestRequest struct {
	Handle        string
	CreatorHandle string
	Description   string
	Status        string
	Summary       string
}

type changeStatusRequest struct {
	Handle       string
	ObjHandle    string
	StatusHandle string
	Desc         string
}

type logoutRequest struct {
	Handle string
}

// funcao que executa login no SDM
// input: user (string) e pass (string)
// output: handle (string)
func doLoginSdm(user, pass string) string {
	logReq := loginRequest{
		Username: user,
		Password: pass,
	}
	request := prepareSoapRequest(logReq, loginXml, "login")
	body := sendRequest(request)
	handle := getInnerTextFromTag(body, "loginResponse", "loginReturn")
	return handle
}

// funcao que recupera handle do SDM
// input: user (string) e pass (string)
// output: handle (string)
func doGetHandleForUseridSdm(user, handle string) string {
	getHandleForUserIdReq := getHandleForUseridRequest{
		Handle:   handle,
		Username: user,
	}
	request := prepareSoapRequest(getHandleForUserIdReq, getHandleForUseridXml, "getHandleForUserid")
	body := sendRequest(request)
	handleForUserid := getInnerTextFromTag(body, "getHandleForUseridResponse", "getHandleForUseridReturn")
	return handleForUserid
}

// funcao que cria ticket no SDM
// input: user (string) e pass (string)
// output: handle (string)
func doCreateRequestSdm(handle, creatorHandle, description, status, summary string) string {
	createRequestReq := createRequestRequest{
		Handle:        handle,
		CreatorHandle: creatorHandle,
		Description:   description,
		Status:        status,
		Summary:       summary,
	}
	request := prepareSoapRequest(createRequestReq, createRequestXml, "createRequest")
	body := sendRequest(request)
	response := getInnerTextFromTag(body, "createRequestResponse", "createRequestReturn")
	idTicketArr := strings.Split(response, "</Handle>")
	idTicketArr = strings.Split(idTicketArr[0], ":")
	idTicket := idTicketArr[1]
	return idTicket
}

// funcao que muda status do ticket no SDM
// input: user (string) e pass (string)
// output: handle (string)
func changeStatusSdm(handle, objHandle, description, status string) string {
	changeStatusReq := changeStatusRequest{
		Handle:       handle,
		ObjHandle:    objHandle,
		Desc:         description,
		StatusHandle: status,
	}
	request := prepareSoapRequest(changeStatusReq, changeStatusXml, "changeStatus")
	body := sendRequest(request)
	response := getInnerTextFromTag(body, "changeStatusResponse", "changeStatusReturn")
	return response
}

// funcao que executa logout no SDM
// input: handle (string)
// output: body (string)
func doLogoutSdm(handle string) string {
	logoutReq := logoutRequest{
		Handle: handle,
	}
	request := prepareSoapRequest(logoutReq, logoutXml, "logout")
	body := sendRequest(request)
	return string(body)
}

// funcao que prepara uma requisicao SOAP
// com um template de xml, populado por valores de uma struct
// input: data (struct), baseXML (string) e soapAction (string)
// output: request (*http.Request)
func prepareSoapRequest(data interface{}, baseXML, soapAction string) *http.Request {
	tmpl, err := template.New(soapAction).Parse(baseXML)
	if err != nil {
		fmt.Printf("Error while marshling object. %s \n", err.Error())
		if logFile != nil {
			logFile.WriteString(fmt.Sprintf("Error while marshling object. %s \n", err.Error()))
		}
	}

	// substituindo valores no template e retornando valor para um buffer de bytes (doc)
	doc := &bytes.Buffer{}
	err = tmpl.Execute(doc, data)
	if err != nil {
		fmt.Printf("template.Execute error. %s \n", err.Error())
		if logFile != nil {
			logFile.WriteString(fmt.Sprintf("template.Execute error. %s \n", err.Error()))
		}
	}

	// codificando o buffer de bytes como xml e salvando em doc
	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		if logFile != nil {
			fmt.Printf("encoder.Encode error. %s \n", err.Error())
			logFile.WriteString(fmt.Sprintf("encoder.Encode error. %s \n", err.Error()))
		}
	}

	if logFile != nil {
		logFile.WriteString(fmt.Sprintf("\n<====================== %s ==========================> \n", soapAction))
		if soapAction != "login" {
			logFile.WriteString(fmt.Sprintf("Request %s \n", doc.String()))
		}
	}

	// criando post request com o conteudo de doc (template modificado e codificado como xml)
	req, err := http.NewRequest(http.MethodPost, server, bytes.NewBuffer([]byte(doc.String())))
	if err != nil {
		fmt.Printf("Error making a request. %s \n", err.Error())
		if logFile != nil {
			logFile.WriteString(fmt.Sprintf("Error making a request. %s \n", err.Error()))
		}
	}
	req.Header = map[string][]string{
		"SOAPAction": {fmt.Sprintf("%s/%s", server, soapAction)},
	}

	return req
}

// funcao que envia uma requisicao SOAP e retorna o body da response
// input: request (*http.Request)
// output: body ([]byte)
func sendRequest(request *http.Request) []byte {
	// enviando post request criada e salvando resposta em resp
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error sending a request. %s \n", err.Error())
		if logFile != nil {
			logFile.WriteString(fmt.Sprintf("Error sending a request. %s \n", err.Error()))
		}
	}
	// lendo conteudo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading a response. %s \n", err.Error())
		if logFile != nil {
			logFile.WriteString(fmt.Sprintf("Error reading a response. %s \n", err.Error()))
		}
	}

	if logFile != nil {
		logFile.WriteString(fmt.Sprintf("Response %s \n", body))
	}

	return body
}

// funcao que recupera o conteudo das tags de uma response
// input: response (*http.Response), parentTag e tags (string)
// output: texts (array de string)
func getInnerTextFromTag(body []byte, parentTag, tag string) string {
	rr, err := xmlquery.Parse(strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}
	channel := xmlquery.FindOne(rr, fmt.Sprintf("//%s", parentTag))
	handle := channel.SelectElement(fmt.Sprintf("%s", tag)).InnerText()
	return string(handle)
}
