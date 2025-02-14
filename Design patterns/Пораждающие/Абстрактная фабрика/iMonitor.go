package main

type IMonitor interface {
	setLogo(logo string)
	setModel(model string)
	getLogo() string
	getModel() string
}

type Monitor struct {
	logo  string
	model string
}

func (m *Monitor) setLogo(logo string) {
	m.logo = logo
}

func (m *Monitor) getLogo() string {
	return m.logo
}

func (m *Monitor) setModel(model string) {
	m.model = model
}

func (m *Monitor) getModel() string {
	return m.model
}
