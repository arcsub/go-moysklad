// Пример удаления договора из документа Отгрузка
package main

import (
	"context"
	"fmt"
	"github.com/arcsub/go-moysklad/moysklad"
	"github.com/google/uuid"
)

func main() {
	client := moysklad.NewClient().WithTokenAuth("MS_TOKEN")

	// Получаем некий документ отгрузки, у которого заполнено поле contract
	demand, _, err := client.Entity().Demand().GetByID(context.Background(), uuid.MustParse("00017575-4179-11ee-0a80-08d300045970"))
	if err != nil {
		panic(err)
	}

	fmt.Println(demand.Contract) // ptr

	// Чтобы удалить договор документа передаём nil в качестве аргумента
	demand.SetContract(nil)

	// Обновляем документ
	// { "contract": null }
	demandUpdated, _, err := demand.Update(context.Background(), client)
	if err != nil {
		panic(err)
	}

	fmt.Println(demandUpdated.Contract) // nil
}
