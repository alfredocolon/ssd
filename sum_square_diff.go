/*
  By:    alfredocolon,   alfredo.j.colon@gmail.com
  Date:  23-Aug-2016

  Solution to projectEuler.net problem # 6
  
  Sum square difference

  Problem 6
  The sum of the squares of the first ten natural numbers is,
  
  12 + 22 + ... + 102 = 385

  The square of the sum of the first ten natural numbers is,
  
  (1 + 2 + ... + 10)2 = 552 = 3025

  Hence the difference between the sum of the squares of the 
  first ten natural numbers and the square of the sum is 3025 - 385 = 2640.
  
  Find the difference between the sum of the squares of the 
  first one hundred natural numbers and the square of the sum.
  
*/

package main

import (
	"fmt";
	"math"
       )

func main() {

  var sum_of_the_squares float64
  var square_of_the_sum  float64
  var i float64

  for i = 1; i<=100; i++ {
    sum_of_the_squares += math.Pow(i,2)
    square_of_the_sum += i
  }

  fmt.Println("The answer is: ", int64(math.Pow(square_of_the_sum,2) - sum_of_the_squares))
}