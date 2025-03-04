package main
import "fmt"

func main(){
	var intNum int16 = 32767
	fmt.Println(intNum)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var floatNum2 float64 = 12345678.9
	fmt.Println(floatNum2)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)
}