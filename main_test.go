package GoMultiRect

import (
	"fmt"
	"image"
	"testing"
)

func TestMultiRect(t *testing.T) {
	var mr MultiRect
	mr.AddRect(image.Rect(0, 0, 500, 500))

	if mr.Area() != 500*500 || len(mr.Rects) != 1 {
		t.Error("check1 failed")
	}
	fmt.Println(mr.Area(), len(mr.Rects)) // must be

	mr.Sub(image.Rect(150, 150, 200, 200))
	if len(mr.Rects) != 8 {
		t.Error("len(Rects) != 8")
		return
	}
	if !mr.Rects[0].Eq(image.Rect(0, 0, 150, 150)) {
		t.Error("rec1 error")
	}
	if !mr.Rects[1].Eq(image.Rect(150, 0, 200, 150)) {
		t.Error("rect2 error")
	}
	if !mr.Rects[2].Eq(image.Rect(200, 0, 500, 150)) {
		t.Error("rect3 error")
	}
	if mr.Area() != (500*500-50*50) || len(mr.Rects) != 8 {
		t.Error("check2 failed")
	}

	r2 := image.Rect(0, 200, 250, 300)
	mr.Sub(r2)
	if mr.Area() != (500*500 - 250*100 - 50*50) {
		t.Error("check3 failed")
	}

	r3 := image.Rect(150, 150, 500, 200)
	mr.Sub(r3)
	if mr.Area() != (500*500 - 250*100 - 50*50 - 300*50) {
		t.Error("check4 failed")
	}

}

func TestMultiRectSub(t *testing.T) {
	var mr MultiRect
	mr.AddRect(image.Rect(0, 0, 500, 500))

	subMultiRect := mr.Intersects(image.Rect(100, 100, 150, 150))

	mr.Sub(image.Rect(100, 100, 150, 150))

	subMultiRect2 := mr.Intersects(image.Rect(50, 50, 200, 200))
	fmt.Println(subMultiRect)
	fmt.Println(subMultiRect2)
}