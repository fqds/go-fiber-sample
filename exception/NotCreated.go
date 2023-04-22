package exception

type NotCreatedObject struct {
}

func (err NotCreatedObject) Error() string {
	return "Object is not created"
}
