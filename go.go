package tml

// Go runs a function in the background and returns a chan error for deferred
// pickup of an error response
func Go(f func() error) <-chan error {
	c := make(chan error, 1)
	go func(c chan error) {
		c <- f()
		close(c)
	}(c)
	return c
}

// Monitor monitors a Go Routine, it provides a channel, on which an error will
// be sent if the Gorouting panics, or nil, when it terminates normally
func Monitor(f func()) <-chan Panic {
	c := make(chan Panic, 1)
	go monitor(c, f)
	return c
}

type (
	Panic interface{}
)

func monitor(c chan Panic, f func()) {
	defer func() {
		if err := recover(); err != nil {
			c <- Panic(err)
		} else {
			c <- nil
		}
		close(c)
	}()
	f()
}
