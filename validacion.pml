#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte consultores = 0
byte mutex = 1
byte datosCompletos = 1

active proctype IngresoDato() {
	do
	::
		wait(datosCompletos)
		printf("Ingreso de dato 1\n")
		printf("Ingreso de dato 2\n")
		printf("Ingreso de dato 3\n")
		signal(datosCompletos)
	od
}

active[2] proctype Consultores() {
	do
	::
		wait(mutex)
		consultores++
		if
		:: (consultores == 1) -> wait(datosCompletos)
		:: else ->
		fi
		signal(mutex)
		
		printf("Prediccion\n")
		
		wait(mutex)
		consultores--
		if
		:: (consultores == 0) -> signal(datosCompletos)
		:: else ->
		fi
		signal(mutex)
	od
}