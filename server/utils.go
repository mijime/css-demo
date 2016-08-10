package server

import (
	"fmt"

	"github.com/go-errors/errors"
)

func Must(err error) {
	if err != nil {
		fmt.Println(err.(*errors.Error).ErrorStack())
		panic(err)
	}
}
