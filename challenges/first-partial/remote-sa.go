package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)
//Struct for a Point
type Point struct {
	x, y float64
}
//Return the X value of the Point
func (p *Point) X() float64 {
	return p.x
}
//Return the Y value of the Point
func (p *Point) Y() float64 {
	return p.y
}
// traditional function
func getDistance(p, q Point) float64 {
	return math.Hypot(q.X()-p.X(), q.Y()-p.Y())
}

// same thing, but as a method of the Point type
func (p Point) getDistance(q Point) float64 {
	return math.Hypot(q.X()-p.X(), q.Y()-p.Y())
}
//Struct for a Vector
type Vector struct {
	x, y float64
}
//Return the X value of the Point
func (v *Vector) X() float64 {
	return v.x
}
//Return the Y value of the Point
func (v *Vector) Y() float64 {
	return v.y
}
//Point Method that transforms two points into a single vector
func (a Point) toVector(b Point) Vector{
	return Vector{(b.X() - a.X()), (b.Y() - a.Y())}
}
//Vector Method that gets the cross product between two Vectors 2X2
func (a Vector) cross(b Vector) float64 {
	return a.X() * b.Y() - a.Y() * b.X()
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8005", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}
	
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}
//Function that checks if two vectors are on the same segment
func onSegment(p, q, r Point) bool {
	if q.X() <= math.Max(p.X(), r.X()) && q.X() >= math.Min(p.X(), r.X()) && q.Y() <= math.Max(p.Y(), r.Y()) && q.Y() >= math.Min(p.Y(), r.Y()) {
		return true
	}
	return false
}
//Checks the orientation of vector 0 = Collinear, 1 = CCW, 2 = CW
func orientation(p, q, r Point) int {
	var val float64 = (q.Y()-p.Y())*(r.X()-q.X()) - (q.X()-p.X())*(r.Y()-q.Y())
	switch {
	case val == 0:
		return 0
	case val > 0:
		return 1
	default:
		return 2
	}
}
//Checks if 4 points intersects
func doIntersect(p1, q1, p2, q2 Point) bool {
	var o1, o2, o3, o4 int = orientation(p1, q1, p2), orientation(p1, q1, q2), orientation(p2, q2, p1), orientation(p2, q2, q1)
	if o1 != o2 && o3 != o4 {
		return true
	}
	if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	}
	if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	}
	if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	}
	if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	}
	return false

}
// getArea gets the area inside from a given shape
func getArea(p []Point) float64 {
	var a float64
	for i := 0; i < len(p); i++ {
		j := (i + 1) % len(p)
		a+=(p[i].X() * p[j].Y()) - (p[j].X() * p[i].Y())
	}
	return math.Abs(a/2.0)
}
// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(p []Point) float64 {
	var per float64
	for i := 0; i < len(p); i++ {
		per += getDistance(p[i], p[(i + 1) % len(p)])
	}
	return per
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}
	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome to the Remote Shapes Analyzer\n")
	response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
	response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
	response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
	response += fmt.Sprintf(" - Area            : %v\n", area)

	// Send response to client
	fmt.Fprintf(w, response)
}
