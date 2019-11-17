package main

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
	"strings"
	"encoding/json"
)

type Dato struct {
	Id int  `json:"id"`
	Columna int  `json:"columna"`
	Valor string  `json:"valor"`
}

func main(){
	con,_ := net.Dial("tcp", "192.168.0.22:8000")
	resp := ""
	r := bufio.NewReader(con)

	for{
		resp = ""
		fmt.Print("Desea agregar una nueva Muestra o predecir? (1: Agregar - 2: Predecir): ")
		fmt.Scanf("%s\n", &resp)
		
		switch resp {
		case "1":
			fmt.Fprintf(con, "AGREGAR\n");

			msgId, _ := r.ReadString('\n')
			msgId = strings.TrimSpace(msgId);
		
			for {
				dato := Dato{};
				dato.Id, _ = strconv.Atoi(msgId)

				fmt.Print("Que dato desea ingresar (1: Dias transcurridos - 2: Color a* - 3: Salir: ")
				fmt.Scanf("%s\n", &resp)

				col, _ := strconv.Atoi(resp);
				dato.Columna = col;
				brake := false;

				var val string;
				switch col {
				case 1:
					fmt.Print("Ingrese los dias transcurridos: ");
					fmt.Scanf("%s\n", &val);
				case 2:
					fmt.Print("Ingrese el color a*: ");
					fmt.Scanf("%s\n", &val);
				case 3:
					fmt.Fprintf(con, "BREAK\n");
					brake = true;
				}

				if brake { 
					break ;
				}

				dato.Valor = val;

				jsonBytes, _ := json.Marshal(dato);
				jsonStr := string(jsonBytes) + "\n";

				fmt.Println(jsonStr);
				fmt.Fprintf(con, jsonStr);
				
			}			
		case "2":
			fmt.Fprintf(con, "PREDECIR\n");
		}
	}
}

func update(){}

func retry(){
	//El dato ingresado es erroneo, intentelo otra vez

}