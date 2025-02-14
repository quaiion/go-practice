package main

import "fmt"

func scanOperand() float64 {
        var scannedOperand float64

        for {
                _, err := fmt.Scan(&scannedOperand)

                if err == nil {
                        break
                } else {
                        var flushStr string = "default"
                        for nFlushed := 1; nFlushed != 0; nFlushed, _ = fmt.Scanln(&flushStr) {}
                        fmt.Print("nope, please enter a real number: ")
                }
        }

        return scannedOperand
}

func main() {
        var (
                operand1 float64
                operand2 float64
                operator string
                result float64
        )

        fmt.Print("please enter the first operand: ")
        operand1 = scanOperand()

        fmt.Print("please enter the second operand: ")
        operand2 = scanOperand()

        fmt.Print("please enter the operator ('+', '-', '*' or '/' sign): ")
        
        for {
                fmt.Scan(&operator)

                var stop bool = true

                switch operator {
                case "+":
                        result = operand1 + operand2
                case "-":
                        result = operand1 - operand2
                case "*":
                        result = operand1 * operand2
                case "/":
                        if operand2 == 0 {
                                fmt.Print("division by zero is forbidden, " +
                                          "please enter another second operand: ")
                                operand2 = scanOperand()
                                fmt.Print("please reenter the operator " +
                                          "('+', '-', '*' or '/' sign): ")
                                stop = false
                        } else {
                                result = operand1 / operand2
                        }
                default:
                        fmt.Print("nope, please try another operator sign. " +
                                  "remember, only '+', '-', '*' and '/' are " +
                                  "an option: ")
                        stop = false
                }
                
                if stop {
                        break
                }
        }

        fmt.Printf("result: %.3f\n", result)
}
