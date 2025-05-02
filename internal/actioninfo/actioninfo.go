package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("mistake in Parse(): %v", err)
		}

		result, err := dp.ActionInfo()
		if err != nil {
			log.Printf("mistake in ActionInfo(): %v", err)
		}
		fmt.Println(result)
	}
}
