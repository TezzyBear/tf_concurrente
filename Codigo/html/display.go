package main

import (
	//"html/template"
	"net/http"
	//"path/filepath"
	"strings"
	"fmt"
)

const dim = 4;
var db [][dim]string; //Slice de muestras
var htmlStr string;

func handle(w http.ResponseWriter, r *http.Request) {

	db = append(db, [4]string{"h1","a2","3","4"})
	db = append(db, [4]string{"h5","6","7d","8"})
	db = append(db, [4]string{"h9","10","11","12"})

	tableStr := ""

	for i:= 0; i < len(db); i++ {
		tableStr += "<tr>"
		for j:= 0 ; j < dim; j ++ {
			tableStr += "<td>"
			tableStr += db[i][j]
			tableStr += "</td>"
		}
		tableStr += "</tr>"
	}
	
	htmlStr = strings.Replace(htmlStr, "REMPLAZAMOSACA", tableStr, 1)

	fmt.Println(htmlStr)

	fmt.Fprintf(w, htmlStr)

}

func main() {
	html()
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}

func html(){
	htmlStr = `
<!DOCTYPE html>
<html>
<head>
<style>
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
</head>  
<body>
<div>
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
<canvas id="myCanvas" width="200" height="100"></canvas>
<script>
var c = document.getElementById("myCanvas");
var ctx = c.getContext("2d");
ctx.moveTo(100,50);
ctx.lineTo(200,100);
ctx.stroke();
</script>
</div>
</body>
</html>
	`
}