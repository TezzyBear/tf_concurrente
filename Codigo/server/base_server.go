package main

import (
	"fmt"
	"net"
	"strings"
	"bufio"
	"encoding/json"
)

type Dato struct {
	Id int  `json:"id"`
	Columna int  `json:"columna"`
	Valor string  `json:"valor"`
}

const dim = 4
var db [][dim]string //Slice de muestras
var counter = 0
var newInst [dim]string

type Muestra struct {
	id int
	valores [dim]string
}

func main(){

	myIP := getOutboundIP();
    
	host := fmt.Sprintf("%s:8000", myIP)
	ln, _ := net.Listen("tcp", host)

	for {
		con, _ := ln.Accept()
		go handle(con)
	}
}

func getOutboundIP() string {
    con, _ := net.Dial("udp", "8.8.8.8:80")
    defer con.Close()
	localAddr := con.LocalAddr().String()
    return strings.Split(localAddr, ":")[0]
}

func handle(con net.Conn){
	defer con.Close()

	r := bufio.NewReader(con)

	for {
		msg, err := r.ReadString('\n');
		if err != nil {
			continue;
		}
		msg = strings.TrimSpace(msg);
		fmt.Println(msg)
		switch msg {
		case "AGREGAR":
			fmt.Fprintf(con, string(len(db)+48)+"\n") //Add new column
			db = append(db, [dim]string{})
			var muestra Muestra;
			chDato := make(chan Muestra)
			go func(){
				for {
					msgDato, err := r.ReadString('\n');
					msgDato = strings.TrimSpace(msgDato);
					if err != nil {
						continue;
					}
					fmt.Println(msgDato)
					if msgDato == "BREAK" {
						chDato<- muestra;
						break;
					}
					var dato Dato;	
					json.Unmarshal([]byte(msgDato), &dato);
					muestra.id = dato.Id;
					muestra.valores[dato.Columna] = dato.Valor;
				}
			}()
				
			muestra = <-chDato;
			db[muestra.id] = muestra.valores;
			fmt.Println(db)

		case "PREDECIR":
			fmt.Println("woooooo prediccion yeah!")
		}

		

		

		//Update

		

		//Predict 

	}	

	//counter = 0
}

func update(db [][4]string, r bufio.Reader){

	

	//fmt.Println(msg)
	//fmt.Println("received")
}