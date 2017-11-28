package rectangle

type Point struct {
	x int
	y int
}

type Rectangle struct {
	botLeftVtx Point
	topRightVtx Point
}

func getTheVertexPosition(v Point, r Rectangle) int {
	/* This method returns a value of
	--> 1 to indicates that the given vertex, 'v' is enclosed within the rectangle 'r'
	--> 0 to indicate that the given vertex, 'v' lies on one of the edges of the rectangle, 'r'
	--> -1 to indicate that the given vertex, 'v' lies completely outside the edges of the rectangle, 'r'
	*/
	if( v.x > r.botLeftVtx.x && v.x < r.topRightVtx.x && v.y > r.botLeftVtx.y && v.y < r.topRightVtx.y ) {
		return 1
	}
	if(((v.x==r.botLeftVtx.x || v.x == r.topRightVtx.x) && v.y >= r.botLeftVtx.y && v.y <= r.topRightVtx.y)) ||
		((v.y==r.botLeftVtx.y || v.y==r.topRightVtx.y) && v.x>=r.botLeftVtx.x && v.x<=r.topRightVtx.x) {
		return 0
	}
	return -1
}

func isTheRectangleValid(r Rectangle) bool {
	/* This method ensures that the vertices that define the rectangle, 'r' are the top right vertex and
	bottom left vertex. Providing other vertices, such as  top left vertex and bottom right vertex for
	the rectangle definition are not accepted.
	*/
	if(r.topRightVtx.x > r.botLeftVtx.x && r.topRightVtx.y > r.botLeftVtx.y) {
		return true
	}
	return false
}

func calcIntersectingRectangle(r1, r2 Rectangle) (Rectangle) {
	/* This method calculates if there is an intersection/overlap of r2 with r1.
	The method returns a Rectangle struct defined by the bottomLeftVertex and topRightVertex if there
	exists an intersection or overlap between the rectangles.

	The method returns a Rectangle struct that has (0,0) values for the bottomLeftVertex and topRightVertex
	if there DOES NOT exist an intersection nor overlap between the rectangles.

	The method accomplishes the below
	Step 1: Get the vertex positions of all the vertices in the rectangle, 'r2' with respect to rectangle, 'r1'.
	Step 2: Determine from the vertex positons,
			--> a. if the rectangle. 'r2' is completely disjoint from the rectangle, r1
			--> b. if the rectangle. 'r2' is completely enveloped inside of rectangle, r1
			--> c. intersection exists because one or two vertices of rectangle, r2 are inside of rectangle, r1.
			--> d. intersection MAY or MAY NOT exist as an edge/side of rectangle, r2 overlaps an edge/side of
					rectangle, r1.
	*/

	//figure out how many vertices of second rectangle lie inside, on the edge or outside the first rectangle.
	var r2VtxPosition  = [4]int{-1,-1,-1,-1}
	r2VtxPosition[0] = getTheVertexPosition(r2.botLeftVtx, r1)
	r2VtxPosition[1] = getTheVertexPosition(Point{r2.botLeftVtx.x, r2.topRightVtx.y}, r1)
	r2VtxPosition[2] = getTheVertexPosition(r2.topRightVtx, r1)
	r2VtxPosition[3] = getTheVertexPosition(Point{r2.topRightVtx.x, r2.botLeftVtx.y}, r1)

	var verticesOnEdge = 0
	var verticesInsideRect = 0
	var verticesOutsideRect = 0
	for i:=0; i<4; i++ {
		if (r2VtxPosition[i] == 1) {
			verticesInsideRect++
		} else if(r2VtxPosition[i] == 0) {
			verticesOnEdge++
		} else {
			verticesOutsideRect++
		}
	}

	if (verticesOutsideRect == 4) {
		return Rectangle{Point{0,0},Point{0,0}}
	}

	/* Below code is equivalent to checking if
	i) all vertices of the second rectangle are inside of the first rectangle
	ii) all vertices of second rectangle are on the edges of the first rectangle
	iii) three vertices of second rectangle are on edges of first rectangle and one vertex of second rectangle is
	 		inside the first rectangle
	iv) two vertices of second rectangle are on edges of first rectangle and two vertices of second rectangle are
			inside the first rectangle */
	if (verticesOnEdge + verticesInsideRect == 4) {
		return r2
	}

	//Calculate the intersection rectangle when a single vertex is inside the first rectangle
	if (verticesInsideRect == 1) {
		var botLeftVtxIntersect, topRightVtxIntersect Point
		if(r2VtxPosition[0] == 1) {
			botLeftVtxIntersect =r2.botLeftVtx
			topRightVtxIntersect =r1.topRightVtx
		} else if(r2VtxPosition[1] == 1) {
			botLeftVtxIntersect=Point{r2.botLeftVtx.x,r1.botLeftVtx.y}
			topRightVtxIntersect=Point{r1.topRightVtx.x, r2.topRightVtx.y}
		} else if(r2VtxPosition[2] == 1) {
			botLeftVtxIntersect=r1.botLeftVtx
			topRightVtxIntersect=r2.topRightVtx
		} else if (r2VtxPosition[3] == 1) {
			botLeftVtxIntersect=Point{r1.botLeftVtx.x, r2.botLeftVtx.y}
			topRightVtxIntersect=Point{r2.topRightVtx.x, r1.topRightVtx.y}
		}
		return Rectangle{botLeftVtxIntersect, topRightVtxIntersect}
	}

	//Calculate the intersection rectangle when two vertices are inside the first rectangle
	if (verticesInsideRect == 2) {
		var botLeftVtxIntersect, topRightVtxIntersect Point
		if(r2VtxPosition[0]==1 && r2VtxPosition[1]==1){
			botLeftVtxIntersect = r2.botLeftVtx
			topRightVtxIntersect = Point{r1.topRightVtx.x, r2.topRightVtx.y}
		} else if (r2VtxPosition[1]==1 && r2VtxPosition[2] == 1) {
			botLeftVtxIntersect = Point{r2.botLeftVtx.x, r1.botLeftVtx.y}
			topRightVtxIntersect = r2.topRightVtx
		} else if (r2VtxPosition[2]==1 && r2VtxPosition[3]==1) {
			botLeftVtxIntersect = Point{r1.botLeftVtx.x, r2.botLeftVtx.y}
			topRightVtxIntersect = r2.topRightVtx
		} else if (r2VtxPosition[3]==1 && r2VtxPosition[0]==1) {
			botLeftVtxIntersect = r2.botLeftVtx
			topRightVtxIntersect =Point{r2.topRightVtx.x, r1.topRightVtx.y}
		}
		return Rectangle{botLeftVtxIntersect, topRightVtxIntersect}
	}

	// Calculate the intersection rectangle when two vertices are on the edge of the first rectangle
	// Next determine if the calculated intersection rectangle is a valid rectangle.
	if (verticesOnEdge == 2) {
		var botLeftVtxIntersect, topRightVtxIntersect Point
		if(r2VtxPosition[0]==0 && r2VtxPosition[1]==0){
			botLeftVtxIntersect = r2.botLeftVtx
			topRightVtxIntersect = Point{r1.topRightVtx.x, r2.topRightVtx.y}
		} else if (r2VtxPosition[1]==0 && r2VtxPosition[2] == 0) {
			botLeftVtxIntersect = Point{r2.botLeftVtx.x, r1.botLeftVtx.y}
			topRightVtxIntersect = r2.topRightVtx
		} else if (r2VtxPosition[2]==0 && r2VtxPosition[3]==0) {
			botLeftVtxIntersect = Point{r1.botLeftVtx.x, r2.botLeftVtx.y}
			topRightVtxIntersect = r2.topRightVtx
		} else if (r2VtxPosition[3]==0 && r2VtxPosition[0]==0) {
			botLeftVtxIntersect = r2.botLeftVtx
			topRightVtxIntersect =Point{r2.topRightVtx.x, r1.topRightVtx.y}
		}

		calculatedRectangle:=Rectangle{botLeftVtxIntersect, topRightVtxIntersect}
		if(isTheRectangleValid(calculatedRectangle)){
			return calculatedRectangle
		}
	}

	return Rectangle{Point{0,0},Point{0,0}}

}

func getTheIntersectingRectangle(r1, r2 Rectangle) Rectangle {
	/* This method returns {{0,0},{0,0}} if there DOES NOT exist an intersection rectangle between the 2 given rectangles,
	else, it returns the calculated intersection values, {{x1, y1}, {x2, y2}} where x1 < x2 and y1 < y2.
	Step 1: Ensure that the rectangles are valid.
	Step 2: caculate if an intersection of rectangle, r2 exists with rectangle, r1.
	Step 3: if step 2 does not return a valid intersection, calculate if intersection of r1 exists with r2.
	*/
	nonIntersectRect:=Rectangle{Point{0,0},Point{0,0}}

	if (! isTheRectangleValid(r1) || ! isTheRectangleValid(r2)) {
		return nonIntersectRect
	}

	intersectRect := calcIntersectingRectangle(r1, r2)
	if (intersectRect != nonIntersectRect) {
		return intersectRect
	}

	/* Note that in the below call r2 is being passed as the first parameter*/
	return calcIntersectingRectangle(r2,r1)
}

