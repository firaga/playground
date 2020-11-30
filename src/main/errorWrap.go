package main

import (
	"errors"
	"fmt"
	"strings"

	pkgerr "github.com/pkg/errors"
)

func main() {
	if err := methodA(false); err != nil {
		//fmt.Printf("%+v", err)
	}
	fmt.Println(strings.Repeat(string('~'), 47))
	if err := methodA(true); err != nil {
		fmt.Printf("%+v", err)
	}
}

func methodA(wrap bool) error {
	if err := methodB(wrap); err != nil {
		if wrap {
			return pkgerr.Wrap(err, "methodA call methodB error")
		}
		return err
	}
	return nil
}

func methodB(wrap bool) error {
	if err := methodC(); err != nil {
		if wrap {
			return pkgerr.Wrap(err, "methodB call methodC error")
		}
		return err
	}
	return nil
}

func methodC() error {
	return errors.New("test error stack")
}
