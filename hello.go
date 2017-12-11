package main
import "fmt"
func main() {
	fmt.Printf("Testing Go!\nHello, World!\n")
	fmt.Println("****************testing outputs**************")
	fmt.Println("highest ODI scores for")
	var player1 string = "Sachin"
	var score int = 200
	player2 := "Dhoni"
	score2 := 183
	fmt.Println(player1 + " Tendulkar", score)
	fmt.Println("MS " + player2, score2)
	fmt.Println("************mathematical operations****************")
	a1 := 10
	b1 := 20
	fmt.Println("a1 = ", a1, "b1 = ", b1)
	/*fmt.Println("sum = ", a1 + b1)
	fmt.Println("diff = ", a1 - b1)
	fmt.Println("product = ", a1 * b1)
	fmt.Println("div = ", a1 / b1)
	fmt.Println("mod = ", a1 % b1)*/
	fmt.Printf("sum = %d\ndiff = %d\nproduct = %d\ndiv = %d\nmod = %d\n", a1 + b1, a1 - b1, a1 * b1, a1 / b1, a1 % b1)
	const pi float64 = 3.146412 //cannnot use := here as const qualifier is being used
	fmt.Println("pi = ", pi)
	fmt.Printf("%f has datatype = %T\n", pi, pi)
	yes := true
	no := false
	fmt.Println("*****************logical operators*****************")
	fmt.Printf("true = %b\nfalse = %b\n", yes, no)
	fmt.Println("yes && no is", yes && no)
	fmt.Println("!yes is", no)
	fmt.Println("****************represeting datatypes**************")
	fmt.Printf("int for char 74 is %c\n", 74)
	fmt.Printf("representing 100 in binary = %b\n", 100)
	fmt.Printf("representing 100 in hex = %x\n", 100)
	fmt.Printf("representing PI is e = %e\n", pi)

	fmt.Println("*******************for-loop********************")
	i := 1
	for i <= 10 {
		fmt.Printf("%d ",i)
		i++
	}
	fmt.Println()
	for i >= 0 {
		fmt.Printf("%d ",i)
		i--
	}
	fmt.Println()
	for j := 2; j < 20; j *= 2 {
		fmt.Printf("%d\n", j)
	}
	fmt.Println("******************testing if-else*****************")
	year := 2017
	if year % 4 == 0 {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Normal Year")
	}
	fmt.Println("testing switch block")
	month := "march"
	switch month {
		case "march" : fmt.Println("March")
		default : fmt.Println("not March")
	}
	fmt.Println("**********using switch inside for loop***********")
	x := 2
	for x <= 10 {
		fmt.Printf("%d ", x)
		switch x % 3 {
			case 0 : fmt.Println("divisible by 3")
			case 1 : fmt.Println("div by 3 remainder is 1")
			default : fmt.Println("div by 3 remainder is 2")
		}
		x++
	}
	fmt.Println("************using arrays********************")
	players := [5]string {"fav players","sachin","dravid","dhoni","kohli"}
	for _, value := range players {
		fmt.Println(value)
	}
	fmt.Println("****************using slices****************")
	nums := []int {1,2,3,4,5,6}
	numsSlice2 := nums[2:4] //acts as a substring second arg gives the length
	fmt.Println("slicing [2:4] from nums {1,2,3,4,5,6}")
	for _,values := range numsSlice2 {
		fmt.Println(values)
	} //without _ it will print the indices
	fmt.Println("************printing the array elements************")
	fmt.Println(nums[:5])
	fmt.Println(players[1:])
	fmt.Println("********making slices and copying******************")
	numSlice := []int {5,4,3,2,1}
	fmt.Println(numSlice[0:])
	fmt.Println("make a slice, give 0 to first 8, max length = 10")
	numSlice2 := make([]int, 8, 10)
	fmt.Println(numSlice2[0:])
	copy(numSlice2, numSlice)
	fmt.Println(numSlice2[0:])
	fmt.Println("*******appending to a slice****************")
	numSlice2 = append(numSlice2,-1,-2,-3)//increases the size of slice by 3 and appends new items at the end, should output [5 4 3 2 1 0 0 0 -1 -2 -3]
	fmt.Println(numSlice2[0:])
	fmt.Println("**************using Maps********************")
	captains := make(map[string] string)
	captains["india"] = "kohli"
	captains["aus"] = "smith"
	captains["eng"] = "root"
	captains["nz"] = "kane"
	fmt.Println(captains)
	fmt.Println("for aus captain is", captains["aus"])
	fmt.Println("length of map is", len(captains))
	delete(captains, "aus")
	fmt.Println("key for aus deleted, length of map =", len(captains))
	fmt.Println(captains)
	fmt.Println("************using function addThemUp for nums******")
	fmt.Println(numSlice2[0:])
	fmt.Println(addThemUp(numSlice2))
	fmt.Println("*******returning multiple values**************")
	mySum, myProduct := sumProduct(numSlice)
	fmt.Println("mySum = ",mySum,"myProduct =",myProduct)
}
//function returning the sum of an array
func addThemUp(nums []int) int {
	x := 0
	for _,values := range nums {
		x += values
	}
	return x
}
//function returning two values, one for sum, other for product
func sumProduct(nums []int) (int, int) {
	fmt.Println(nums)
	x := 0
	p := 1
	for _, values := range nums {
		x += values
		p *= values
	}
	return x, p
}
