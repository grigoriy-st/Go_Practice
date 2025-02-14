package main

import "fmt"

type ITechFactory interface {
	makeLaptop() ILaptop
	makeMonitor() IMonitor
}

func GetTechFactory(brand string) (ITechFactory, error) {
	if brand == "Lenovo" {
		return &Lenovo(), nil
	}

	if brand == "Asus" {
		return &Asus{}, nil
	}

	return nil, fmt.Errorf("Wrong brand")
}
