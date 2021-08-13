package hospital

import "fmt"

type Doctor struct {
	Next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.Next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.Next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
	d.Next = next
}
