package main
import "fmt"
func main() {
	fmt.Println("Yoo man")
	fmt.Println(add(3,2))
	a:=5
	b:=8
	fmt.Println("Value: ",a,",",b," swapped :", fmt.Sprint(swap(a,b)))
	fmt.Println("Value: ", a," ",b, " Decremented: ", fmt.Sprint(decrement_values(a,b)) )
	//fmt.Println(fmt.Sprintf("Values:%d,%d, Decremented Values: %d, %d ", a, b, decrement_values(a,b)))
}

func decrement_values(a,b int) (x,y int){
	x = a-2
	y = b-2
	return
}

func swap(a,b int ) (int, int) {
	return b,a
}

func add(a,b int ) int {
	return a+b
}
