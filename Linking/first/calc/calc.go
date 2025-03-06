// calc/calc.go
package calc

import (
	"first/models" // Импортируем пакет models
	"fmt"
)

// PrintDataExpression - функция для вывода данных выражения
func PrintDataExpression(expr models.Expression) {
	fmt.Printf("Operand1: %f, Operand2: %f, Operator: %s\n", expr.Operand1, expr.Operand2, expr.Operator)
}
