package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	// Импортируем официальный SDK под псевдонимом b24 для удобства
	b24 "github.com/bitrix24/b24gosdk"
)

func main() {
	// 1. Создаем контекст выполнения с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// 2. Вставьте ваш URL входящего вебхука из Битрикс24
	webhookURL := "https://b24-c74vn3.bitrix24.ru/rest/1/73x59n2yr79kfv4n/"

	// 3. Инициализируем клиент SDK
	client := b24.NewClient(webhookURL)

	// 4. Формируем параметры товара. В SDK используется тип b24.Params
	// (это удобный аналог стандартного map[string]any)
	productParams := b24.Params{
		"fields": b24.Params{
			"NAME":        "Товар 33", // через Go SDK",
			"PRICE":       2450.00,
			"CURRENCY_ID": "RUB",
			"XML_ID":      "1C_CODE_77771", // Наш будущий мостик для 1С
		},
	}

	fmt.Println("⏳ Отправка запроса через b24gosdk...")

	// 5. Вызываем метод. Сервис Core() выполняет любой стандартный метод REST API.
	// Метод crm.product.add является безопасным для повтора, если сеть моргнет,
	// но SDK рекомендует явно указывать WithIdempotent() для чтения или точечной записи.
	res, err := client.Core().Call(ctx, "crm.product.add", productParams)
	if err != nil {
		fmt.Printf("❌ Ошибка выполнения метода: %v\n", err)
		return
	}

	// 6. SDK возвращает сырой результат в res.Result.
	// Метод crm.product.add возвращает просто число (ID нового товара).
	// Декодируем его в специальный тип b24.ID, который защищает от путаницы строк и чисел в API.
	var productID b24.ID
	if err := json.Unmarshal(res.Result, &productID); err != nil {
		fmt.Printf("❌ Ошибка декодирования ID товара: %v\n", err)
		return
	}

	// Победа!
	fmt.Println("✅ Успех!")
	fmt.Printf("📦 Товар успешно создан в облаке Битрикс24. ID товара: %v\n", productID)

	// // Полный путь к методу создания товара в CRM
	// apiURL := webhookURL + "catalog.product.add.json"

	// res, err := client.Core().Call(ctx, "catalog.product.add", b24.Params{
	// 	"fields": b24.Params{
	// 		"iblockId": 23,
	// 		"name":     "Товар",
	// 	},
	// })
	// if err != nil {
	// 	return fmt.Errorf("catalog.product.add: %w", err)
	// }

	// // Метод заворачивает ответ в объект с ключом "element".
	// raw, ok := b24.Unwrap(res.Result, "element")
	// if !ok {
	// 	return fmt.Errorf("в ответе нет ключа element")
	// }

	// var item struct {
	// 	Active     string `json:"active"`
	// 	Available  string `json:"available"`
	// 	Bundle     string `json:"bundle"`
	// 	CanBuyZero string `json:"canBuyZero"`
	// 	Code       string `json:"code"`
	// 	CreatedBy  int    `json:"createdBy"`
	// }
	// if err := json.Unmarshal(raw, &item); err != nil {
	// 	return fmt.Errorf("разбор ответа: %w", err)
	// }
	// fmt.Println(item.Active, item.Available)
}
