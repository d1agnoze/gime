package internal

import "fmt"

var ERR_MISSING_ARG error = fmt.Errorf("Missing Arguments")
var ERR_UNSUPPORT_MIME error = fmt.Errorf("mime type not supported")
var ERR_INVALID_URL error = fmt.Errorf("invalid url")
