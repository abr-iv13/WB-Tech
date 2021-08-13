package main

import (
	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/chain-of-responsibility/hospital"
)

func main() {

	cashier := &hospital.Cashier{}

	//Set next for medical department
	medical := &hospital.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &hospital.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &hospital.Reception{}
	reception.SetNext(doctor)

	patient := &hospital.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
