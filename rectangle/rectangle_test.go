package rectangle_test

import (
	"testing"
	
	"github.com/vishwanathj/golangexercises/rectangle"
)

func TestGetTheRectangleIntersectFunc(t *testing.T) {
	/* Note: This would ALSO test invalid rectangles as part of its test data*/
	testTable := [] struct {
		r1 rectangle.Rectangle
		r2 rectangle.Rectangle
		output rectangle.Rectangle
	} {
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1,}}, rectangle.Rectangle{rectangle.Point{2,2}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1}}},
		{rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{2,2}}},
		{rectangle.Rectangle{rectangle.Point{1,-1}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{1,0}, rectangle.Point{2,2}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{2,2}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,2}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,2}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{4,2}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{3,2}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{-1,1}, rectangle.Point{0,2}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,2}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,2}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{0,-1}, rectangle.Point{1,2}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,2}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1}}, rectangle.Rectangle{rectangle.Point{1,0}, rectangle.Point{2,1}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,2}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1}}, rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{1,-1}, rectangle.Point{2,3}}, rectangle.Rectangle{rectangle.Point{1,0}, rectangle.Point{2,3}}},
		{rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{4,2}}, rectangle.Rectangle{rectangle.Point{1,0}, rectangle.Point{4,3}}, rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{4,2}}},
		{rectangle.Rectangle{rectangle.Point{4,2}, rectangle.Point{0,1}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,2}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
	}
	for _, tdr := range testTable {
		res:=rectangle.GetTheIntersectingRectangle(tdr.r1, tdr.r2)
		if (res != tdr.output) {
			t.Errorf("Calculated result %v incorrect, correct result is %v", res, tdr.output)
		}
	}

}

func TestIsTheRectangleValidFunc(t *testing.T) {
	testTable := [] struct {
		input rectangle.Rectangle
		output bool
	} {
		{ rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1}}, true},
		{rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{0,0}}, false},
		{rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,0}}, false},
		{rectangle.Rectangle{rectangle.Point{1,0}, rectangle.Point{0,1}}, false},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,0}}, false},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,1}}, false},
	}
	for _, tdr := range testTable {
		res:=rectangle.IsTheRectangleValid(tdr.input)
		if (res != tdr.output) {
			t.Errorf("Calculated result %t incorrect, correct result is %t", res, tdr.output )
		}
	}
}

func TestGetTheVertexPositionFunc(t *testing.T) {
	testTable := [] struct {
		v rectangle.Point
		r rectangle.Rectangle
		output int
	} {
		{rectangle.Point{2,2}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, 1},
		{rectangle.Point{-1,-1}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, -1},
		{rectangle.Point{0,3}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, 0},
	}
	for _, tdr := range testTable {
		res:=rectangle.GetTheVertexPosition(tdr.v, tdr.r)
		if (res != tdr.output) {
			t.Errorf("Calculated result %d incorrect, correct result is %d", res, tdr.output )
		}
	}
}

func TestCalculateIntersectingRectangleFunc(t *testing.T) {
	/*Note: also tests cases where the ordering of the rectangles being passed as parameter to the method matters,
	there are cases where passing r2 as first parameter and r1 as second parameter may give different result as
	compared to when passing r1 as first parameter and r2 as second parameter.*/
	testTable := [] struct {
		r1 rectangle.Rectangle
		r2 rectangle.Rectangle
		output rectangle.Rectangle
	} {
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1,}}, rectangle.Rectangle{rectangle.Point{2,2}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{2,2}}},
		{rectangle.Rectangle{rectangle.Point{1,1}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{1,-1}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
		{rectangle.Rectangle{rectangle.Point{1,-1}, rectangle.Point{3,3}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{1,0}, rectangle.Point{2,2}}},
		{rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,2}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{0,0}}},
		{rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{2,2}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,2}}, rectangle.Rectangle{rectangle.Point{0,1}, rectangle.Point{1,2}}},
	}
	for _, tdr := range testTable {
		res:=rectangle.CalcIntersectingRectangle(tdr.r1, tdr.r2)
		if (res != tdr.output) {
			t.Errorf("Calculated result %v incorrect, correct result is %v", res, tdr.output)
		}
	}
}
