package main

import "github.com/jestebanrods/edu-golang/experimental/internal/patterns/chainresponsibility"

func main() {
	cashier := &chainresponsibility.Cashier{}

	// Set next for medical department
	medical := &chainresponsibility.Medical{}
	medical.SetNext(cashier)

	// Set next for doctor department
	doctor := &chainresponsibility.Doctor{}
	doctor.SetNext(medical)

	// Set next for reception department
	reception := &chainresponsibility.Reception{}
	reception.SetNext(doctor)

	patient := &chainresponsibility.Patient{Name: "abc"}

	// Patient visiting
	reception.Execute(patient)
}
