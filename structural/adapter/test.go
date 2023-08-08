package adapter

func TestAdapter() {
	client := &Client{}

	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
