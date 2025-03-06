// main.go
package main

import (
	"first/calc" // Импортируем пакет calc
	"first/models"
)

func main() {
	// Создаем новое выражение
	expr := models.NewExpression(5.0, 3.0, "+")

	// Выводим данные выражения
	calc.PrintDataExpression(expr)
}
