package chainofresponsibility

func TestChainOfResponsibility() {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}

	//Patient visiting
	reception.execute(patient)
}
