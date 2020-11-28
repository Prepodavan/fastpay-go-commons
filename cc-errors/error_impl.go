package cc_errors

import (
	"encoding/json"
	"fmt"
)

func (err *BaseError) Error() string {
	const jsonErrFormat = "{" +
		"\"code\": %d," +
		" \"message\": \"Ошибка при формирование структуры ошибки. %s\"," +
		" \"data\": \"\"" +
		"}"
	byteError, jsonErr := json.Marshal(err)
	if jsonErr != nil {
		return fmt.Sprintf(jsonErrFormat, ErrorDefault, jsonErr.Error())
	}
	return string(byteError)
}
