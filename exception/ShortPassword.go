package exception

type ShortPassword struct {
}

func (err ShortPassword) Error() string {
	return "Short Password"
}
