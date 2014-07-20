package tml

// Go runs a function in the background and returns a chan error for deferred
// pickup of an error response
func Go(f func() error) <-chan error {
	c := make(chan error)
	go func(c chan error) {
		c <- f()
	}(c)
	return c
}
