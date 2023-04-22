package exception

type NotAuthenticated struct {
}

func (err NotAuthenticated) Error() string {
	return "Not authenticated"
}
