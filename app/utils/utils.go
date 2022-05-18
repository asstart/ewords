package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data interface{}) string {
	res, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf("can't encode, raw data: %v", data)
	}
	return fmt.Sprintf("%v", string(res))
}
