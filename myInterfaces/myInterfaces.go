package myinterfaces

type Interface1 interface {
	ShowFirstItem()
}

type Peripheral struct {
	Brand string
	Model string
}

type Dog struct {
	Name string
	Size string
}

func (objectPeripheral Peripheral) ShowFirstItem() string {
	return objectPeripheral.Brand
}

func (objectDog Dog) ShowFirstItem() string {
	return objectDog.Name
}
