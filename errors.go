package decorators

import "fmt"

var ErrNotFunc = fmt.Errorf("received value is not a function")
var ErrCannotSetFunc = fmt.Errorf("programming error: cannot set decorated function")
var ErrCannotCastFunc = fmt.Errorf("programming error: cannot cast decorated function")
