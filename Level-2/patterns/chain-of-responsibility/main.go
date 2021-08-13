package main

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.SetNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
