package main

import (
	//"html/template"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	//"path/filepath"

	"strings"
)

const dim = 5

var db [][dim]string //Slice de muestras
var htmlStr string

func readcsv() {
	csvfile, err := os.Open("../machine_learning/strawberriesData.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	lines, _ := r.ReadAll()
	for i, line := range lines {
		if i == 0 {
			continue
		}
		db = append(db, [5]string{line[2], line[3], line[4], line[5], line[1]})
	}
}

func handle(w http.ResponseWriter, r *http.Request) {

	tableStr := ""

	for i := 0; i < len(db); i++ {
		tableStr += "<tr>"
		for j := 0; j < dim; j++ {
			tableStr += "<td>"
			tableStr += db[i][j]
			tableStr += "</td>"
		}
		tableStr += "</tr>"
	}

	arrdatos := "["

	for i := 0; i < len(db); i++ {
		arrdatos += "["
		for j := 0; j < dim; j++ {
			arrdatos += "'"
			arrdatos += db[i][j]
			if j == dim-1 {
				arrdatos += "'"
			} else {
				arrdatos += "',"
			}
		}
		if i == len(db)-1 {
			arrdatos += "]"
		} else {
			arrdatos += "],"
		}
	}
	arrdatos += "]"

	htmlStr = strings.Replace(htmlStr, "REMPLAZAMOSACA", tableStr, 1)
	htmlStr = strings.Replace(htmlStr, "datosaqui", arrdatos, 1)
	//fmt.Println(arrdatos)
	fmt.Fprintf(w, htmlStr)

}

func main() {
	readcsv()
	html()
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}

func html() {
	htmlStr = `
<!DOCTYPE html>
<html>
<head>
<script src="https://cdn.jsdelivr.net/npm/@tensorflow/tfjs/dist/tf.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/p5.js/0.7.3/p5.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/p5.js/0.7.3/addons/p5.dom.js"></script>
<style>
#arriba{
	float:top;
}
#abajo{
	float:bottom;
}

#container{
	float:left;
	padding: 30px;
	height: 200px;
	widht:400px;
}
#aderecha{
	float:right;
	padding: 50px;
}

table.scroll {
width: 650px; /* Optional */
border-collapse: collapse;
border: 2px solid #000;	 
}
table.scroll tbody, table.scroll thead {
display: block; 
}
table.scroll tbody {
height: 300px;
overflow-y: auto;
overflow-x: hidden;
}
tbody td, thead th {
width: 124px;
padding: 10px;
}
thead tr{
background-color: rgb(21, 6, 105) ;
}
tbody tr{
background-color: white;
}
body{
background: #4e54c8;
background: -webkit-linear-gradient(to left, #8f94fb, #4e54c8);
background: linear-gradient(to left, #8f94fb, #4e54c8); 
}
td {
border-bottom: solid 1px rgb(21, 6, 105);
}
</style>

<script src="https://cdn.anychart.com/releases/v8/js/anychart-base.min.js"></script>
<script src="https://cdn.anychart.com/releases/v8/js/anychart-ui.min.js"></script>
<script src="https://cdn.anychart.com/releases/v8/js/anychart-exports.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/regression/1.4.0/regression.min.js"></script>
</head>  
<body>
<div id="contenido">
<div id="arriba">
<h1 align="center" style="color: black;">Trabajo final Concurrente</h1>
<table class="scroll" align="center" border="0">
<thead>
<tr style="color: white">
<th>a*</th>
<th>b*</th>
<th>L*</th>
<th>E*</th>
<th>Tiempo (días)</th>
</tr>
</thead>
<tbody align="center">
REMPLAZAMOSACA
</tbody>
</table>
</div>
<div id="abajo">
<div id="container">
<script>
var datos = datosaqui;
var rawData = [];

for (i=0;i<datos.length;i++){
	var temp = [];
	temp.push(datos[i][4]);
	temp.push(datos[i][3]);
	rawData.push(temp);
}

console.log(rawData);

var result = regression('linear', rawData);

//get coefficients from the calculated formula
var coeff = result.equation;

anychart.onDocumentReady(function () {

  var data_1 = rawData;
  var data_2 = setTheoryData(rawData);

  chart = anychart.scatter();

  // creating the first series (marker) and setting the experimental data
  var series1 = chart.marker(data_1);

  // creating the second series (line) and setting the theoretical data
  var series2 = chart.line(data_2);


  chart.container("container");
  chart.draw();
});

function formula(coeff, x) {
  var result = null;
  for (var i = 0, j = coeff.length - 1; i < coeff.length; i++, j--) {
    result += coeff[i] * Math.pow(x, j);
  }
  return result;
}

function setTheoryData(rawData) {
  var theoryData = [];
  for (var i = 0; i < rawData.length; i++) {
    theoryData[i] = [rawData[i][0], formula(coeff, rawData[i][0])];
  }
  return theoryData;
}

</script>
</div>
<div id="aderecha">
</div>
</div>
</div>
</body>
</html>
	`
}
