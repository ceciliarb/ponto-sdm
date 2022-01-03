# Ponto SDM

## Usage
```bash
--------- Exemplo de uso do programa ----------------------------------------------------------------
./ponto-sdm -u username -p senha
./ponto-sdm --username username --password senha --log
./ponto-sdm --u username --p senha --l --action a 
./ponto-sdm --u username --p senha --l --action abrir 
./ponto-sdm --u username --p senha --l --action p 
./ponto-sdm --u username --p senha --l --action paralisar 
./ponto-sdm --u username --p senha --l --action r 
./ponto-sdm --u username --p senha --l --action retomar 
./ponto-sdm --u username --p senha --l --action f 
./ponto-sdm --u username --p senha --l --action F --id-ticket 2455454 
./ponto-sdm --u username --p senha --l --action finalizar --num-ref-ticket 934850 

------------------ Parametros -----------------------------------------------------------------------
  -a string
        alias para -action.
  -action string
        Especifique uma ação ('abrir', 'paralisar', 'retormar', 'finalizar'). (default "abrir")
  -id-ticket string
        Especifique um id ticket (talvez exista um no arquivo .idTicket).
  -l    alias para -log.
  -log
        Se estiver presente, armazena log das operações.
  -p string
        alias para -password.
  -password string
        Especifique uma senha. (default "-")
  -ref-num-ticket string
        Especifique um ref_num do ticket (número do ticket no SDM).
  -rnt string
        alias para -ref-num-ticket.
  -s string
        alias para -server.
  -server string
        Especifique um servidor.
  -t string
        alias para -id-ticket.
  -u string
        alias para -username
  -username string
        Especifique um username. (default "ceci")

```


## Manutenção
```bash
git clone github.com/ceciliarb/ponto-sdm
# windows
export CGO_ENABLED=0
env GOOS=windows GOARCH=386 go build
# linux
go build
```