package factorymethod

type musketFactory struct{}

func newMusketFactory() *musketFactory {
	return &musketFactory{}
}

func (f *musketFactory) getGun() (IGun, error) {
	return newMusket(), nil
}
