package loan

func Calculate(p float32, r float32, t int) (intrest float32) {
	intrest = p * r * float32(t) / 100
	return
}
