package main

import (
	"fmt"
	"net"
	"strings"
	"strconv"
	"bufio"
	"encoding/json"
	"os/exec"
	"os"
	"encoding/csv"
	"math"
	"io"
)

type Dato struct {
	Id int  `json:"id"`
	Columna int  `json:"columna"`
	Valor string  `json:"valor"`
}

const dim = 5
var db [][dim]string //Slice de muestras
var counter = 0
var newInst [dim]string

type Muestra struct {
	id int
	valores [dim]string
}

type Pred struct {
	a float64
	b float64
	L float64
}

func main(){

	csvToDb("../machine_learning/strawberriesData.csv");

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
					if err != nil {
						continue;
					}
					msgDato = strings.TrimSpace(msgDato);
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

			//con los datos recibidos

			predChan := make(chan Pred)
			
			go func(){
				for{
					msgPred, err := r.ReadString('\n');
					if err != nil {
						continue;
					}
					msgPred = strings.TrimSpace(msgPred);
					var predData Pred;
					json.Unmarshal([]byte(msgPred), &predData);
					predChan <-predData;
				}
				
			}()
			
			predData := <-predChan;		
			
			predLinear, _:= predecir(calcE(predData.a, predData.b, predData.L));

			fmt.Fprintf(con, fmt.Sprintf("%f", predLinear)) //Add new column

		}
	}	
}

func update(db [][4]string, r bufio.Reader){

	

	//fmt.Println(msg)
	//fmt.Println("received")
}

func csvToDb(filePath string){
    // Loading file.
    f, _ := os.Open(filePath)
    // Reader.
	r := csv.NewReader(f)	
    for {
		record, err := r.Read()
		// Stop at EOF.		
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
		}	
		muestra := [dim]string{}
        for i, val := range record {
			if i == 0 { // id
				continue;
			}
			muestra[i-1] = val;
		}
		db = append(db, muestra)
	}
}

func calcE(a, b, L float64) float64{

	a0Str := db[0][0];
	b0Str := db[0][0];
	L0Str := db[0][0];

	a0, _ := strconv.ParseFloat(a0Str, 64);
	b0, _ := strconv.ParseFloat(b0Str, 64);
	L0, _ := strconv.ParseFloat(L0Str, 64);

	E := math.Sqrt(math.Pow(a0-a,2)+math.Pow(b0-b,2)+math.Pow(L0-L,2));

	return E
}

func predecir(E float64) (float64, float64) {

	toPred := fmt.Sprintf("%f", E)

	//enviar comando de ejecucion python
	cmd := exec.Command("python", "LinearRegression.py", "makePrediction", toPred)
	out, _ := cmd.CombinedOutput()

	//limpiar output
	strOut := string(out)
	strOut = strings.TrimSpace(strOut)
	strOut = strings.Replace(strOut, "[", "", -1)
	strOut = strings.Replace(strOut, "]", "", -1)
	split := strings.Split(strOut, ",")

	fmt.Println(split)

	//conversion de resultados (split[0] = regresion lineal, split[1] = regresion polinomial)
	linearResult, _ := strconv.ParseFloat(split[0], 64)
	//print para probar
	
	polynomialResult, _ := strconv.ParseFloat(split[1], 64)

	return linearResult, polynomialResult
	
	//solo usar los valores de linearResult o polynomialResult
	//ambos tienen ventajas y desventajas
}