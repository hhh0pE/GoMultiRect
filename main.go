package GoMultiRect

import "image"

type MultiRect struct {
	Rects []image.Rectangle
}

func (mr MultiRect) Area() float64 {
	var area float64
	for _, r := range mr.Rects {
		area += float64(r.Dx()) * float64(r.Dy())
	}
	return area
}

func (mr *MultiRect) AddRect(r image.Rectangle) {
	mr.Rects = append(mr.Rects, r)
}

func (mr *MultiRect) Sub(rin image.Rectangle) {
	var newRects []image.Rectangle
	for _, r := range mr.Rects {
		i := rin.Intersect(r)
		if i.Max.X == 0 && i.Max.Y == 0 {
			newRects = append(newRects, r)
			continue
		}

		// make 8 rects

		// left-top
		if nr := image.Rect(r.Min.X, r.Min.Y, i.Min.X, i.Min.Y); !nr.Empty() {
			newRects = append(newRects, nr)
		}
		// center-top
		if nr := image.Rect(i.Min.X, r.Min.Y, i.Max.X, i.Min.Y); !nr.Empty() {
			newRects = append(newRects, nr)
		}
		// right-top
		if nr := image.Rect(i.Max.X, r.Min.Y, r.Max.X, i.Min.Y); !nr.Empty() {
			newRects = append(newRects, nr)
		}
		// left-center
		if nr := image.Rect(r.Min.X, i.Min.Y, i.Min.X, i.Max.Y); !nr.Empty() {
			newRects = append(newRects, nr)
		}
		// right-center
		if nr := image.Rect(i.Max.X, i.Min.Y, r.Max.X, i.Max.Y); !nr.Empty() {
			newRects = append(newRects, nr)
		}
		// left-bottom
		if nr := image.Rect(r.Min.X, i.Max.Y, i.Min.X, r.Max.Y); !nr.Empty() {
			newRects = append(newRects, nr)
		}
		// center-bottom
		if nr := image.Rect(i.Min.X, i.Max.Y, i.Max.X, r.Max.Y); !nr.Empty() {
			newRects = append(newRects, nr)
		}
		// right-bottom
		if nr := image.Rect(i.Max.X, i.Max.Y, r.Max.X, r.Max.Y); !nr.Empty() {
			newRects = append(newRects, nr)
		}

	}
	mr.Rects = newRects
}

