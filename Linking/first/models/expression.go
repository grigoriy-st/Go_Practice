// models/expression.go
package models

// Expression - структура, представляющая математическое выражение
type Expression struct {
	Operand1 float64
	Operand2 float64
	Operator string
}

// NewExpression - функция для создания нового выражения
func NewExpression(op1, op2 float64, operator string) Expression {
	return Expression{Operand1: op1, Operand2: op2, Operator: operator}
}
