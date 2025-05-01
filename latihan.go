package main
import "fmt"
const NMAX int = 100
func maximumProduct(nums []int) int {
    var n, i, hasil int
    hasil = 1
    n = len(nums)
    for i = 0; i < n; i++{
        hasil *= nums[i]
    }
    return hasil
}
func main(){
	var n, i int
	var A[NMAX]int
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	hasil := maximumProduct(A[:n])
	fmt.Print(hasil)
}