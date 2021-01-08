package client

import "fmt"

func wrappError(customMsg string, originalErr error) error {
	return fmt.Errorf("%s : %v", customMsg, originalErr)
}
