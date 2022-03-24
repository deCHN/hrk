package hrk

import (
	"fmt"
	"testing"
)

/* https://betterprogramming.pub/how-to-crack-the-top-25-golang-interview-questions-a94396d6c808
 * Implement the SumOfSquares function, which takes an integer c and returns the sum of all squares between 1 and c. You'll need to use select statements, goroutines, and channels.
For example, entering 5 would return 55 because $1^2 + 2^2 + 3^2 + 4^2 + 5^2 = 55$
*/
func SumOfSquares(c, quit chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}
}

func TestSumOfSquares(t *testing.T) {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0

	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 0
	}()

	SumOfSquares(mychannel, quitchannel)
}
