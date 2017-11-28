package rectangle

import "testing"

func TestGetTheRectangleIntersectFunc(t *testing.T) {
	/* Note: This would ALSO test invalid rectangles as part of its test data*/
	testTable := [] struct {
		r1 Rectangle
		r2 Rectangle
		output Rectangle
	} {
		{Rectangle{Point{0,0}, Point{1,1,}}, Rectangle{Point{2,2}, Point{3,3}}, Rectangle{Point{0,0}, Point{0,0}}},
		{Rectangle{Point{0,0}, Point{1,1}}, Rectangle{Point{0,0}, Point{1,1}}, Rectangle{Point{0,0}, Point{1,1}}},
		{Rectangle{Point{1,1}, Point{2,2}}, Rectangle{Point{0,0}, Point{3,3}}, Rectangle{Point{1,1}, Point{2,2}}},
		{Rectangle{Point{1,-1}, Point{3,3}}, Rectangle{Point{0,0}, Point{2,2}}, Rectangle{Point{1,0}, Point{2,2}}},
		{Rectangle{Point{0,0}, Point{2,2}}, Rectangle{Point{1,1}, Point{3,3}}, Rectangle{Point{1,1}, Point{2,2}}},
		{Rectangle{Point{0,0}, Point{3,3}}, Rectangle{Point{0,1}, Point{1,2}}, Rectangle{Point{0,1}, Point{1,2}}},
		{Rectangle{Point{0,0}, Point{3,3}}, Rectangle{Point{0,1}, Point{4,2}}, Rectangle{Point{0,1}, Point{3,2}}},
		{Rectangle{Point{0,0}, Point{3,3}}, Rectangle{Point{-1,1}, Point{0,2}}, Rectangle{Point{0,0}, Point{0,0}}},
		{Rectangle{Point{0,0}, Point{2,2}}, Rectangle{Point{0,1}, Point{1,2}}, Rectangle{Point{0,1}, Point{1,2}}},
		{Rectangle{Point{0,0}, Point{2,2}}, Rectangle{Point{0,-1}, Point{1,2}}, Rectangle{Point{0,0}, Point{1,2}}},
		{Rectangle{Point{0,0}, Point{1,1}}, Rectangle{Point{1,0}, Point{2,1}}, Rectangle{Point{0,0}, Point{0,0}}},
		{Rectangle{Point{0,0}, Point{1,1}}, Rectangle{Point{0,1}, Point{1,2}}, Rectangle{Point{0,0}, Point{0,0}}},
		{Rectangle{Point{0,0}, Point{1,1}}, Rectangle{Point{1,1}, Point{2,2}}, Rectangle{Point{0,0}, Point{0,0}}},
		{Rectangle{Point{0,0}, Point{3,3}}, Rectangle{Point{1,-1}, Point{2,3}}, Rectangle{Point{1,0}, Point{2,3}}},
		{Rectangle{Point{0,1}, Point{4,2}}, Rectangle{Point{1,0}, Point{4,3}}, Rectangle{Point{1,1}, Point{4,2}}},
		{Rectangle{Point{4,2}, Point{0,1}}, Rectangle{Point{0,1}, Point{1,2}}, Rectangle{Point{0,0}, Point{0,0}}},
	}
	for _, tdr := range testTable {
		res:=getTheIntersectingRectangle(tdr.r1, tdr.r2)
		if (res != tdr.output) {
			t.Errorf("Calculated result %v incorrect, correct result is %v", res, tdr.output)
		}
	}

}

func TestIsTheRectangleValidFunc(t *testing.T) {
	testTable := [] struct {
		input Rectangle
		output bool
	} {
		{ Rectangle{Point{0,0}, Point{1,1}}, true},
		{Rectangle{Point{1,1}, Point{0,0}}, false},
		{Rectangle{Point{0,1}, Point{1,0}}, false},
		{Rectangle{Point{1,0}, Point{0,1}}, false},
		{Rectangle{Point{0,0}, Point{1,0}}, false},
		{Rectangle{Point{0,0}, Point{0,1}}, false},
	}
	for _, tdr := range testTable {
		res:=isTheRectangleValid(tdr.input)
		if (res != tdr.output) {
			t.Errorf("Calculated result %t incorrect, correct result is %t", res, tdr.output )
		}
	}
}

func TestGetTheVertexPositionFunc(t *testing.T) {
	testTable := [] struct {
		v Point
		r Rectangle
		output int
	} {
		{Point{2,2}, Rectangle{Point{0,0}, Point{3,3}}, 1},
		{Point{-1,-1}, Rectangle{Point{0,0}, Point{3,3}}, -1},
		{Point{0,3}, Rectangle{Point{0,0}, Point{3,3}}, 0},
	}
	for _, tdr := range testTable {
		res:=getTheVertexPosition(tdr.v, tdr.r)
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
		r1 Rectangle
		r2 Rectangle
		output Rectangle
	} {
		{Rectangle{Point{0,0}, Point{1,1,}}, Rectangle{Point{2,2}, Point{3,3}}, Rectangle{Point{0,0}, Point{0,0}}},
		{Rectangle{Point{0,0}, Point{3,3}}, Rectangle{Point{1,1}, Point{2,2}}, Rectangle{Point{1,1}, Point{2,2}}},
		{Rectangle{Point{1,1}, Point{2,2}}, Rectangle{Point{0,0}, Point{3,3}}, Rectangle{Point{0,0}, Point{0,0}}},
		{Rectangle{Point{0,0}, Point{2,2}}, Rectangle{Point{1,-1}, Point{3,3}}, Rectangle{Point{0,0}, Point{0,0}}},
		{Rectangle{Point{1,-1}, Point{3,3}}, Rectangle{Point{0,0}, Point{2,2}}, Rectangle{Point{1,0}, Point{2,2}}},
		{Rectangle{Point{0,1}, Point{1,2}}, Rectangle{Point{0,0}, Point{2,2}}, Rectangle{Point{0,0}, Point{0,0}}},
		{Rectangle{Point{0,0}, Point{2,2}}, Rectangle{Point{0,1}, Point{1,2}}, Rectangle{Point{0,1}, Point{1,2}}},
	}
	for _, tdr := range testTable {
		res:=calcIntersectingRectangle(tdr.r1, tdr.r2)
		if (res != tdr.output) {
			t.Errorf("Calculated result %v incorrect, correct result is %v", res, tdr.output)
		}
	}
}