# API tasks management

## Run

run this command from project root:

```sh
make run
```
This will bring up the server on the indicated port.

## Run test
run this command from project root:

```sh
make test
```
this command will execute all the tests of the project.

## Collection postman

In the collection-postman folder there is a collection to test the API from postman, download and import it.

## interviewgo
A simple exercise for a golang techincal interview


## Challenge
Completar el CRUD del TODO Task, agregando funcionalidad para que el usuario pueda agregar, modificar, buscar mediante el id y eliminar una tarea.

como ayuda están los tests que se encuentran dentro del archivo task_test.go alojado en la dirección pkg/handlers/task

para correr los tests simplemente con ejecutar

`make test`

Como dato adicional dentro de los tests están fixtures como ejemplos de lo que se espera que haga la aplicación.