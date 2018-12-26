package main

import "errors"

func Coverage(conditon bool) error {
	if conditon {
		return errors.New("condition was set")
	}
	return nil
}
