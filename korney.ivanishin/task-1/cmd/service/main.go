package main

import "fmt"

func scan_operand(operand_ptr *float64) {
        for {
                _, err := fmt.Scan(operand_ptr)

                if err == nil {
                        break
                } else {
                        var flush_str string = "default"
                        for n_flushed := 1; n_flushed != 0; n_flushed, _ = fmt.Scanln(&flush_str) {}
                        fmt.Print("nope, please enter a real number: ")
                }
        }
}

func main() {
        var (
                operand_1 float64
                operand_2 float64
                operator string
                result float64
        )

        fmt.Print("please enter the first operand: ")
        scan_operand(&operand_1)

        fmt.Print("please enter the second operand: ")
        scan_operand(&operand_2)

        fmt.Print("please enter the operator ('+', '-', '*' or '/' sign): ")
        
        for {
                fmt.Scan(&operator)

                var stop bool = true

                switch operator {
                case "+":
                        result = operand_1 + operand_2
                case "-":
                        result = operand_1 - operand_2
                case "*":
                        result = operand_1 * operand_2
                case "/":
                        if operand_2 == 0 {
                                fmt.Print("division by zero is forbidden, " +
                                          "please enter another second operand: ")
                                scan_operand(&operand_2)
                                fmt.Print("please reenter the operator " +
                                          "('+', '-', '*' or '/' sign): ")
                                stop = false
                        } else {
                                result = operand_1 / operand_2
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
