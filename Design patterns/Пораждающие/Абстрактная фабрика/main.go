package main

import "fmt"

func main() {
	asusFactory, _ := GetTechFactory("Asus")
	lenovoFactory, _ := GetTechFactory("Lenovo")

	asusLaptop := asusFactory.makeLaptop()
	asusMonitor := asusFactory.makeMonitor()

	lenovoLaptop := lenovoFactory.makeLaptop()
	lenovoMonitor := lenovoFactory.makeMonitor()

	printMonitorDetails(asusMonitor)
	printLaptopDetails(asusLaptop)

	printMonitorDetails(lenovoMonitor)
	printLaptopDetails(lenovoMonitor)
}

func printLaptopDetails(m ILaptop) {
	fmt.Printf("Logo: %s\n", m.getLogo())
	fmt.Printf("Model: %s\n", m.getModel())
}

func printMonitorDetails(m IMonitor) {
	fmt.Printf("Logo: %s\n", m.getLogo())
	fmt.Printf("Model: %s\n", m.getModel())
}
