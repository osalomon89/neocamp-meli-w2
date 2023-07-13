# Deadlock

El error "all goroutines are asleep - deadlock!" en Go indica que se ha producido un bloqueo (deadlock) en el programa debido a un problema de sincronización incorrecta entre goroutines.
Un bloqueo (deadlock) ocurre cuando hay un punto en el programa en el que todas las goroutines están esperando a que algo suceda, pero no hay ninguna acción que pueda desencadenar ese algo. Como resultado, el programa se queda en un estado de bloqueo y no puede avanzar.
