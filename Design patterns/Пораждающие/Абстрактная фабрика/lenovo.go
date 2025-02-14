package main

type Lenovo struct {
}

func (l *Lenovo) makeLaptop() ILaptop {
	return &LenovoLaptop{
		Laptop: Laptop{
			logo:  "Lenono",
			model: "ThinkPad",
		},
	}
}

func (l *Lenovo) makeMonitor() IMonitor {
	return &LenovoMonitor{
		Monitor: Monitor{
			logo:  "Lenovo",
			model: "Legion",
		},
	}
}
