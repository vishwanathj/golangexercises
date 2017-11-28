package main

import (
	"fmt"

	"github.com/vishwanathj/golangexercises/rectangle"

)

/*Below array holds the static test data.
Each row has EIGHT integer entries.
The FIRST four entries in each row represents the bottom left vertex and
top right vertex of rectangle R1. The LAST four entries in each row represents
the bottom left vertex and top right vertex of rectangle R2.
Feel free to insert or delete test data between the "Begin" & "End" comments.
*/
var testData = [][]int{
	//Begin
	{0,0,1,1,2,2,3,3} ,
	{0,0,1,1,0,0,1,1},
	{-1,1,0,2,0,0,3,3},
	{0,0,2,2,0,1,1,2},
	//End
	}


func main() {

	for _, h := range testData {
		r1:= rectangle.Rectangle{rectangle.Point{h[0], h[1]}, rectangle.Point{h[2],h[3]}}
		r2:= rectangle.Rectangle{rectangle.Point{h[4], h[5]}, rectangle.Point{h[6],h[7]}}
		disp:= fmt.Sprintf("The intersecting rectangle of r1: %v and r2: %v is %v", r1, r2, rectangle.GetTheIntersectingRectangle(r1,r2))
		fmt.Println(disp)
		fmt.Println()
	}
}
