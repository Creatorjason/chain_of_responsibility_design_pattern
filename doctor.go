package main

import "fmt"

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor check up done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking up on patient now")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}
