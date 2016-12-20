package car

type  Car struct {
	Name string
	Calor string
}

func (c Car) Start() string {
	return c.Name + " has been started"
}
