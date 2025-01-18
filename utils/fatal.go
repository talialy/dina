package utils

import "log"

// A wraper for log.Fatal to show
// errors that are NOT expected
func Fatal(err error) {
	log.Fatalf("%s\nYou may report this incident with a pull request to help us improve\nsorry for the inconvenience!", err)
}
