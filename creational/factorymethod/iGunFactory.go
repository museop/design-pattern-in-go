package factorymethod

type IGunFactory interface {
	getGun(IGun, error)
}
