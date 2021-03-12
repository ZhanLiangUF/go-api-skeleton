package errdefs

func getImplementer(err error) error {
	switch e := err.(type) {
		case
	}
}

func isNotFound(err error) bool {
	_, ok := getImplementer(err).(Err)
}