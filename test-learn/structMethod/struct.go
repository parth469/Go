package structMethod

func Perimeter(r Rectangle) float64 {

	if r.Width < 0 || r.Height < 0 {
		return 0.0
	}

	return 2 * (r.Width + r.Height)
}

