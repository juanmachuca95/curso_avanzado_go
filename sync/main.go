package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

// Muchos Despositos -- un solo escritor
// Un balance -- muchos lectores

/*
	IMPORTANTE
	Solo para aclarar las diferencias entre RWLock y no usar nada:

	Lock bloquea lecturas (con RLock) y escrituras (con Lock) de otras goroutines
	Unlock permite nuevas lecturas (con Rlock) y/o otra escritura (con Lock)
	RLock bloquea escrituras (Lock) pero no bloquea lecturas (RLock)
	RUnlock permite nuevas escrituras (y también lecturas, pero por la naturaleza de RLock, estas no se vieron bloqueadas nunca)
	En esencia, RLock de RWLock garantiza una secuencia de lecturas en donde el valor que lees no se verá alterado por nuevos escritores, a diferencia de no usar nada.
*/

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))
}
