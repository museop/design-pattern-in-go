package factorymethod

type AK47Factory struct{}

func newAK47Factory() *AK47Factory {
	return &AK47Factory{}
}

func (f *AK47Factory) getGun() (IGun, error) {
	return newAk47(), nil
}
