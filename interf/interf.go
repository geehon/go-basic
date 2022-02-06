package interf

func DemoInterface() {
	var car Car
	car = Tesla{"特斯拉"}
	car.getSpeed()
	car = Wow{"沃尔沃"}
	car.getSpeed()
}

type Car interface {
	getSpeed()
}

type Tesla struct {
	name string
}

type Wow struct {
	name string
}

func (te Tesla) getSpeed() {
	println(te.name, "时速：100km/h")
}

func (wow Wow) getSpeed() {
	println(wow.name, "时速：120km/h")
}
