package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	switch t := value.(type){
	case string:
		color.HiRed(t)
	case int:
		fmt.Printf("Код ошиюки: %v",t)
	case error:
		color.HiRed(t.Error())
	default:
		fmt.Println("Неизвестный тип ошибки")
	}
}