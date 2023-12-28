package data

import (
	"fmt"
	"strconv"
)

type Runtinme int32

func (r Runtinme) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
