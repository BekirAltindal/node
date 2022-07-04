/** Common error types
 *
 * @author Steffen Vogel <svogel2@eonerc.rwth-aachen.de>
 * @copyright 2014-2022, Institute for Automation of Complex Power Systems, EONERC
 * @license Apache 2.0
 *********************************************************************************/

package errors

import (
	"C"
)
import "fmt"

var (
	ErrEndOfFile = fmt.Errorf("end-of-file")
)

func ErrorToInt(e error) int {
	if e == nil {
		return 0
	} else {
		return -1
	}
}

func IntToError(ret int) error {
	if ret == 0 {
		return nil
	} else {
		return fmt.Errorf("ret=%d", ret)
	}
}
