# Generador de poesías

Es un pequeño programa que se entrena con poesías y luego puede decir frases que podría haber dicho esa persona

## Modo de uso

- Escribir en el archivo alfonsina.txt las poesías con las que se quiere entrenar al sistema
- Ejecutar go run generador.go

Cuantas más poesías se agreguen mejor será el texto generado

## ¿Cómo funciona?

El libro de frases se alimenta de las poesías, para cada palabra de la poesía registra cuales son las palabras que le siguen.

El predictor toma una de las palabras de la poesía, busca todas las palabras que le siguen y elige una al azar, luego repite el proceso empezando por la palabra que eligió al azar.

Está basado en la idea de las cadenas de Markov http://setosa.io/ev/markov-chains/
