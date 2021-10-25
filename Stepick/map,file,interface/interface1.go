package main

import "fmt"

func printError(value interface{}) {

	fmt.Printf("value=%v: %T", value, value)
}
func readTask() (value1, value2, operation interface{}) {
	return value1, value2, operation
}


func main() {
	value1, value2, operation := readTask()
	typeValue1 := fmt.Sprintf("%T", value1)
	typeValue2 := fmt.Sprintf("%T", value2)
	if num1, ok := value1.(float64); ok {
		if num2, ok := value2.(float64); ok{
			switch operation {
			case "+" : fmt.Printf("%.4f", num1+num2 )
			case "-" : fmt.Printf("%.4f", num1-num2 )
			case "*" : fmt.Printf("%.4f", num2*num1 )
			case "/" : fmt.Printf("%.4f", num1/num2 )
			default:  fmt.Printf("неизвестная операция")
				break
			}
		}

	}
	if typeValue1 != "float64" {
		printError(value1)
	}
	if typeValue2 != "float64" {
		printError(value2)
	}
}
/*func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
                                            // все полученные значения имеют тип пустого интерфейса
    var v1,v2 float64
    var ok bool
    if v1, ok = value1.(float64); !ok {
        fmt.Printf("value=%v: %T", value1, value1)
        return
    }
    if v2, ok = value2.(float64); !ok {
        fmt.Printf("value=%v: %T", value2, value2)
        return
    }
    switch op, ok := operation.(string); {
        case !ok:
            fmt.Printf("неизвестная операция")
        case op == "+":
            fmt.Printf("%.4f", v1 + v2)
        case op == "-":
            fmt.Printf("%.4f", v1 - v2)
        case op == "/":
            fmt.Printf("%.4f", v1 / v2)
        case op == "*":
            fmt.Printf("%.4f", v1 * v2)
        default:
            fmt.Printf("неизвестная операция")
    }
}*/