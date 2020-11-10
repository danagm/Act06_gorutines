package main

import (
	"container/list"
	"fmt"
	"time"
)

var boolean bool
var deletedList list.List

type Process struct {
	id   int
	done bool
}

func (p *Process) start(c chan string) {
	go printProcess(c)
	i := 0
	for {
		if deletedList.Len() != 0 && deletedList.Front().Value == p.id {
			endP := deletedList.Front()
			deletedList.Remove(endP)
			return
		}
		if boolean == true {
			s := fmt.Sprint("id ", p.id, ": ", i)
			c <- s
		}
		i++
		time.Sleep(time.Millisecond * 500)
	}
}

func printProcess(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	opc := 1
	var count = 0
	var processList list.List

	boolean = false
	c := make(chan string)

	for opc != 0 {
		fmt.Println("1.Agregar proceso\n2.Mostrar proceso\n3.Eliminar proceso\n0.Salir")
		fmt.Scan(&opc)

		if opc == 0 {
			return
		} else if opc == 1 {
			p := Process{
				id:   count,
				done: false,
			}

			go p.start(c)
			processList.PushBack(p)

			fmt.Printf("Se agregó el proceso %d\n", count)
			count++
		} else if opc == 2 {
			boolean = !boolean
		} else if opc == 3 {
			var input int
			fmt.Print("Proceso a eliminar: ")
			fmt.Scan(&input)

			for process := processList.Front(); process != nil; process = process.Next() {
				pID := process.Value.(Process).id
				if pID == input {
					fmt.Printf("Se eliminó el proceso %d\n", pID)
					deletedList.PushBack(input)
					// Delete
					processList.Remove(process)
					break
				}
			}
		} else {
			return
		}
	}
}
