package main

type ILaptop interface {
	setLogo(logo string)
	setModel(model string)
	getLogo() string
	getModel() string
}

type Laptop {
	logo string
	model string
}

func (l *Laptop) setLogo(logo string) {
	l.logo = logo
}

func (l *Laptop) getLogo() string {
	return l.logo
}

func (l *Laptop) setModel(model string) {
	l.model = model
}

func (l *Laptop) getModel() string{
	return l.model
}