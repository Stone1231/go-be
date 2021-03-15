package models

import (
	"strconv"
	"strings"
)

// Uint uint
type Uint uint

func (t *Uint) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	dataStr = strings.ReplaceAll(dataStr, "\"", "")
	if data == nil || dataStr == `""` {
		*t = 0
		return nil
	}
	sec, err := strconv.ParseUint(dataStr, 10, 32)
	if err != nil {
		return err
	}
	*t = Uint(sec)
	return nil
}
