package usecase

import "fmt"

func ErrorMessage(err error) string {
	return fmt.Sprintf("Произошла ошибка: %s", err.Error())
}
