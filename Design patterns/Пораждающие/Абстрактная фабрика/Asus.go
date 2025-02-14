package main

type Asus struct {
}

func (a *Asus) makeLaptop(ILaptop) {
	return &AsusLaptop{
		Laptop: Laptop{
			logo:  "Asus",
			model: "VivoBook",
		},
	}
}

func (a *Asus) makeMonitor(IMonitor) {
	return &AsusMonitor{
		Monitor: Monitor{
			logo:  "Asus",
			model: "TUF Gaming",
		},
	}
}
