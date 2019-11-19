**Integrantes**
- Renzo Vargas        -----------------------Codigo: u201521485
- Luis Guillermo          ---------------------    Codigo: u201612216
- Mauricio Rodiguez       ----------------   Codigo: u201510971
- Fabricio Torrico        --------------------       Codigo: u201416750



# Detección de pérdida de color en la pulpa de fresas

### Resumen

En este trabajo evaluaremos la industria alimenticia, la cual a tenido un aumento desde el año pasado, aumento en la exportación y ventas. En dicha área se evidencio un problema con la pulpa de fresa, la cual posee características que se van perdiendo con el tiempo mientras no se encuentre fresca.
Por este motivo y en el siguiente trabajo con la ayuda de machine Lenin y programación concurrente, una plataforma donde se podrá calcular la probabilidad de que la pulpa de fruta pierda sus propiedades y ya no sea comercial según las características ingresadas.


**Índice**
1. Objetivo del Estudiante (Student Outcome)
2. Capítulo 1: Presentación
3. Capítulo 2: Marco
4. Capítulo 3: Gestión
5. Capítulo 4: Implementacie solución
6. Conclusiones
7. Recomendaciones
8. Glosario
9. Bibliografía




## Objetivo del Estudiante (Student Outcome)

ABET – CAC - Student Outcome 5: Funcionar eficazmente como miembro o líder de un equipo que participa en actividades apropiadas para la disciplina del programa.

## Capítulo 1: Presentación

La industria alimenticia en el Perú ha crecido un 17% durante el 2018, esto indica que el desempeño de los diversos sectores ha desarrollado de mejor manera sus productos mediante el uso de diversas tecnologías y aumento de relaciones comerciales. Tal es el caso del rubro de producción, venta y exportación de pulpa de fresa, la cual es usada en diversos productos alimenticios como mermeladas, yogures, dulces, entre otros. Esta pulpa es tratada con temperaturas de -20° C para mantener las propiedades biológicas, es decir, sin intervención de microorganismos que ocasionan que la pulpa pierda calidad. Sin embargo, aún ocurren cambios químicos al mantener congelado el producto, lo que provoca pérdida de características esenciales en la pulpa.

Se requiere que esta pulpa sea fresca, concentrada y además que posea un color rojizo. El problema radica cuando el color rojizo de la pulpa se empieza a perder dado a la oxidación generada en el transcurso del tiempo. Esto ocasiona que no sea comercial, debido a la pérdida de calidad de los componentes químicos del producto. Para esto se tiene que medir diariamente cuanto cambia el color y esto se les dificulta a muchos trabajadores ya que no tienen un estimado de en cuanto tiempo podría esta pulpa superar el límite de cambio de color permitido por convención.

Se propone realizar un programa que pueda predecir mediante técnicas de machine learning la fecha en la cual el color de la pulpa de fresa supere el limite establecido; utilizando un registro previo de propiedades específicas de la pulpa de fresa. Se planea poder registrar dicho historial utilizando blockchain para poder mantener los registros. Esto, va a garantizar que la información sea permanente e inmutable.



## Capítulo 2: Marco

Como ya fue mencionado previamente, para este trabajo utilizaremos ‘Blockchain’, este termino se puede explicar como un conjunto de tecnologías que emplean funciones criptográficas, además que facilitan la transferencia de un valor de un lugar a otro sin que aparezca alguna intervención de un tercero. Blockchain es un amplio registro de datos los cuales permiten que estén enlazados y cifrados de manera segura para mantener la confidencialidad y seguridad de las transacciones. (Dávila, 2019)

![](https://praxent.com/wp-content/uploads/2018/04/blockchain-1024x773.png)

Por tales motivos se aplicó Blockchain para la seguridad y registro de los datos que se van a almacenar mediante la función de el aplicativo.

Además, en el uso de machine learning se utilizó la regresión lineal y polinomial, para poder obtener la predicción de la clase deseada. Regresión lineal explica la relación que existe entre una variable dependiente y un grupo de variables independientes X1..., Xn. Para dicha técnica donde una variable Y expuesta a una variable X, se busca una función la cual sea más optima en la aproximación de un conjunto de puntos (xi,yi), mediante una curva. (Carollo Limeres, 2011-2012)
Se puede expresar mediante la siguiente expresión:

<img src="https://latex.codecogs.com/gif.latex?Y&space;=&space;\alpha&space;&plus;&space;\beta&space;X&space;&plus;&space;\varepsilon" title="Y = \alpha + \beta X + \varepsilon" />

 

Donde α es la ordenada en el origen, β es la pendiente de la recta y β una variable que incluye un conjunto grande de factores. X e Y son variables aleatorias, por ende, se pude generar una relación lineal entre estas variables. (Carollo Limeres, 2011-2012)

![](https://user-images.githubusercontent.com/40810772/69093084-c0026a80-0a1b-11ea-8e81-357f83d9d740.PNG)

Regresión polinomial es una forma de regresión lineal en la que la relación entre la variable independiente x y la variable dependiente y es modelada como un polinomio de grado n en x. A diferencia de la lineal, la polinomial puede llegar a captar mejor algunos modelos donde la curva puede ser utilizada mediante un exponente. (Agarwal, 2018)
 
<img src="https://latex.codecogs.com/gif.latex?Y&space;=&space;\theta&space;_{0}&space;&plus;&space;\theta&space;_{1}&space;x&space;&plus;&space;\theta&space;_{2}&space;x^{2}" title="Y = \theta _{0} + \theta _{1} x + \theta _{2} x^{2}" />

Con esto se puede ver que la curva puede llegar a pasar por mas puntos que la lineal.

![](https://user-images.githubusercontent.com/40810772/69093603-c34a2600-0a1c-11ea-9587-c47455e3dc87.png)


## Capítulo 3: Gestión
Para el desarrollo de este trabajo se utilizó la herramienta de github para poder llevar un control y el avance de todas las actividades a realizar, además de poder almacenar los archivos necesarios para dicha implementación.
![](https://user-images.githubusercontent.com/40810772/68733291-1dbe2f00-05a4-11ea-902e-687df6348ac7.PNG)
![](https://user-images.githubusercontent.com/40810772/68733081-9cff3300-05a3-11ea-86a1-c1609d6594a7.PNG)

Además, se utilizó la herramienta lucidchart para diseñar la vista de la aplicación y poder llevar un mejor enfoque a el producto terminado.
![](https://user-images.githubusercontent.com/40810772/68996639-dae1ad00-086a-11ea-8b2b-74c4b8b37e2c.png)
Como se ve en la imagen, se podrá visualizar los datos ingresados en una tabla con sus respectivos atributos detectando la clase "Días", además que se mostrará una gráfica comparando todos los datos ingresados.


## Capítulo 4: Implementación de solución
Se desarrollo un programa el cual obtuvo 2 partes, un código desde la vista del cliente y otro por parte del servidor.

En la implementación del cliente se enviará un mensaje avisando al que se agregara una nueva muestra, se le hará el aviso al servidor mediante un ADD y el servidor va a agregar una nueva fila a la base de datos regresando un ID al cliente, luego el cliente podrá mandar un dato el cual contendrá el ID y un dato extra, además de poder seleccionar donde se encontrara el dato, esto se podrá repetir varias veces. Por último, se podrá dar la opción a salir en donde se mandará un mensaje para avisar la salida, por lo cual el servidor agregará la muestra a un channel, dándole un valor y generando un valor al ID donde será agregado a el conjunto de muestras totales.

![](https://user-images.githubusercontent.com/40810772/69024037-107fb680-098f-11ea-8a34-e7013a4fbc7b.png)

Los datos que se obtendrán e ingresar son los siguientes:

| Tiempo   (días) | a*    | b*   | L*    | E*   |
|-----------------|-------|------|-------|------|
| 0               | 18.09 | 8.68 | 37.70 | 0.00 |
| 7               | 17.22 | 9.05 | 39.28 | 1.84 |

Donde la variable Días es el tiempo que tomara en malograrse y la clase a detectar, a*, b* y L* son los colores que va a ir tomando la muestra a lo largo del tiempo y E* es la distancia euclidiana la cual es la distancia euclidiana donde expresa el color desde el primer día en adelante.

Luego de obtener los datos agregados, se procede a hacer un cálculo para poder hallar la clase resultante, para esto se prueba regresión lineal y polinomial. Para esto se consideran las 3 dimensiones del espacio del color, además tiene el tiempo en días, además se utilizó la variable E * ya que es la distancia euclidiana y posee más relación lineal que las demás variables. Este algoritmo fue aplicado en Python, al ejecutar la relación polinomial se realizo con grado 3 ya que se obtienen mejores resultados, al final se ejecuta el algoritmo y retorna la predicción en días tanto lineal como polinomial.

## Conclusiones
Luego del desarrollo del proyecto, podemos concluir que los datos obtenidos de la pulpa de fresa tienen un patrón, en donde se pueden predecir sus datos mediante técnicas estadísticas sin una gran complejidad. Además, luego del desarrollo se evidencio que la regresión polinomial obtiene mejores resultados que la regresión lineal para esta predicción con los datos mencionados.

## Recomendaciones
Para futuros trabajos seria mas optimo mediante técnica de machine learning, reconocer el color de manera automática y no manual para obtener datos mas exactos y una mejor predicción. Ademas trabajar con un sistema de archivos que pueda guardar la nueva información en memoria principal para integrar los distintos modulos y poder exportar los archivos a otros medios de análisis.

## Glosario
- Blockchain
- Regresión lineal
- Regresión polinomial

## Bibliografía
Dávila, C. C. (30 de 10 de 2019). BBVA. Obtenido de https://www.bbva.com/es/pe/bbva-open-talks-lima-es-turno-de-hablar-de-blockchain/

Agarwal, A. (8 de Octubre de 2018). Towards Data Science. Obtenido de https://towardsdatascience.com/polynomial-regression-bbe8b9d97491

Carollo Limeres, M. C. (2011-2012). REGRESIÓN LINEAL SIMPLE. Estadística. FBA.





