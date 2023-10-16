![](https://dev.moysklad.ru/doc/api/remap/1.2/images/logo.svg)

# go-moysklad (МойСклад)

SDK для работы с [МойСклад JSON API 1.2](https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api)

> **Внимание!** SDK находится в стадии разработки!
> 
> Некоторые методы могут отсутствовать или работать неправильно!

# Установка

> Требуемая версия go >= 1.9
> 
```
go get github.com/arcsub/go-moysklad
```

# Особенности
## [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp/)
- В данном решении используется go-retryablehttp с корректировкой в необходимом для работы API МойСклад заголовке.
- При получении ошибки 429 запрос будет повторяться N количество раз (по умолчанию 5). Изменить количество попыток можно с помощью метода клиента `WithMaxRetries()`

## Возвращаемые аргументы
Каждый запрос на создание/изменение/удаление возвращает 3 аргумента.
Рассмотрим объявление функции 
```go
func (s *endpointCreate[T]) Create(ctx context.Context, entity *T, params *Params) (*T, *Response, error)
```
В примере выше нас интересуют возвращаемые аргументы: `(*T, *Response, error)`
1. `*T` – указатель на сущность/документ, например *Product при вызове `Create()` (возвращает `bool` при вызове метода `Delete()`).
2. `*Response` – сырой ответ на запрос, *http.Response, обёрнутый в *Response.
3. `error` – ошибки, если они были. При возникновении ошибок от API МойСклад в качестве ошибки будет заполненная структура `ApiErrors`

## Указатели
Поля структур сущностей и документов являются указателями.
- Для безопасного разыменовывания указателя необходимо передать указатель в метод `Deref()`
```go
  name := moysklad.Deref(product.Name)
```
- Чтобы установить указатель на примитивное значение поля также существуют вспомогательные методы:
  - `Bool()` возвращает *bool
  - `Int()` возвращает *int
  - `Uint()` возвращает *uint64
  - `Float()` возвращает *float64
  - `String()` возвращает *string

Пример:
```go
  product := &moysklad.Product{
    Name: moysklad.String("Apple iPhone 15"),
    Active: moysklad.Bool(false),
  }
```
# Использование
## Создание экземпляра клиента
```go
  client := moysklad.NewClient()
```

## Аутентификация
Имеется два способа аутентификации.
- С помощью токена. Метод клиента `WithTokenAuth()`
```go
  client := moysklad.NewClient().WithTokenAuth(os.Getenv("MOYSKLAD_TOKEN"))
```

- С помощью пары логин/пароль. Метод клиента `WithBasicAuth()`
```go
  client := moysklad.NewClient().WithBasicAuth(os.Getenv("MOYSKLAD_USERNAME"), os.Getenv("MOYSKLAD_PASSWORD"))
```

## Методы клиента
### WithTokenAuth(token)
Получить простой клиент с авторизацией через токен.
```go
  client := moysklad.NewClient().WithTokenAuth(os.Getenv("MOYSKLAD_TOKEN"))
```
### WithBasicAuth(username, password)
Получить простой клиент с авторизацией через пару логин/пароль.
```go
  client := moysklad.NewClient().
	  WithBasicAuth(os.Getenv("MOYSKLAD_USERNAME"), os.Getenv("MOYSKLAD_PASSWORD"))
```
### WithMaxRetries(retries)
Устанавливает максимальное кол-во попыток для одного запроса и возвращает простой клиент.
По умолчанию у запроса 5 попыток.
```go
  // установим 10 попыток для каждого запроса
  client := moysklad.NewClient().WithMaxRetries(10)
```
### WithDisabledWebhookContent(value)
Временное отключение уведомлений вебхуков
```go
  // отключим уведомления вебхуков на данном клиенте
  client := moysklad.NewClient().WithDisabledWebhookContent(true)
```

### Async()
Методы для работы с асинхронными задачами.

Относительный путь: `/async`
```go
  asyncService := clientExt.Async
```
### Audit()
Методы для работы с аудитом.

Относительный путь: `/audit`
```go
  auditService := clientExt.Audit
```
### Context()
Методы для работы с контекстом.

Относительный путь: `/context`
```go
  contextService := clientExt.Context
```
### Entity()
Методы для работы с сущностями и документами.

Относительный путь: `/entity`
```go
  entityService := clientExt.Entity
```
### Report()
Методы для работы с отчётами.

Относительный путь: `/report`
```go
  reportService := clientExt.Report
```
### Security()
Относительный путь: `/security`

Методы для получения токена.
```go
  securityService := clientExt.Security
```

## Параметры запроса
### Создать экземпляр для работы с параметрами запроса
```go
params := new(moysklad.Params)
```

### Методы для работы с параметрами запроса

#### Количество элементов на странице `limit=val`
Пример:
```go
params.WithLimit(100)
```

#### Смещение от первого элемента `offset=val`
Пример:
```go
params.WithOffset(100)
```

#### Контекстный поиск `search=val`
Пример:
```go
params.WithSearch("iPhone 15")
```
#### Замена ссылок объектами
Пример:
```go
params.WithExpand("positions").WithExpand("group")
```

#### Фильтрация по значению `key=value`
Пример:
```go
params.WithFilterEquals("name", "Яблоко")
```

#### Строго больше `key>value`
Пример:
```go
params.WithFilterGreater("sum", "100")
```

#### Строго меньше `key<value`
Пример:
```go
params.WithFilterLesser("sum", "1000")
```

#### Больше или равно `key>=value`
Пример:
```go
params.WithFilterGreaterOrEquals("moment", "2023-06-01")
```

#### Меньше или равно `key<=value`
Пример:
```go
params.WithFilterLesserOrEquals("moment", "2023-06-01")
```

#### Не равно `key!=value`
Пример:
```go
params.WithFilterNotEquals("name", "0001")
```

#### Частичное совпадение (обычное подобие) `key~value`
Пример:
```go
params.WithFilterEquivalence("code", "ms")
```

#### Полное совпадение в начале значения (левое подобие) `key~=value`
Пример:
```go
params.WithFilterEquivalenceLeft("code", "ms")
```

#### Полное совпадение в конце значения (правое подобие) `key=~value`
Пример:
```go
params.WithFilterEquivalenceRight("code", "ms")
```

#### Частичное совпадение не выводится `key!~value`
Пример:
```go
params.WithFilterNotEquivalence("code", "ms")
```

#### Фильтрация по удалённым документам `isDeleted=val`
Пример:
```go
params.WithFilterDeleted(true)
```

#### Фильтрация по напечатанным документам `printed=val`
Пример:
```go
params.WithFilterPrinted(true)
```

#### Фильтрация по опубликованным документам `published=val`
Пример:
```go
params.WithFilterPublished(true)
```

#### Фильтрация по архивным сущностям `archived=val`
Пример:
```go
params.WithFilterArchived(true)
```

#### Группировка выдачи `groupBy=val`
Пример:
```go
params.WithGroupBy(moysklad.GroupByProduct)
```

#### Применение сохранённого фильтра `namedFilter=href`
Метод принимает указатель на сохранённый фильтр.

Пример:
```go
params.WithNamedFilter(&NamedFilter{...})
```

#### Сортировка по умолчанию `order=fieldName`
Пример:
```go
params.WithOrder("name")
```

#### Сортировка по возрастанию `order=fieldName,asc`
Пример:
```go
params.WithOrderAsc("name")
```

#### Сортировка по убыванию `order=fieldName,desc`
Пример:
```go
params.WithOrderDesc("name")
```

#### Остатки и себестоимость в позициях документов `fields=stock`
Метод устанавливает лимит позиций в 100 единиц, а также применяет замену ссылок объектами для поля `positions`

Пример:
```go
params.WithStockFiled()
```

#### Тип остатка `stockType=val`
Используется в отчёте "Остатки" 

Пример:
```go
params.WithStockType(moysklad.StockDefault)
```

#### Интервал, с которым будет построен отчет `interval=val`
Используется в отчётах

Пример:
```go
params.WithInterval(moysklad.IntervalMonth)
```

#### Начало периода `momentFrom=val`
Метод принимает `time.Time`
Пример:
```go
params.WithMomentFrom(time.Now())
```

#### Конец периода `momentTo=val`
Метод принимает `time.Time`
Пример:
```go
params.WithMomentTo(time.Now())
```

## Сервисы
Для перехода к определённому сервису необходимо вызвать цепочку методов, аналогично пути запроса.

### Пример: ProductService
Сервис для работы с товарами.

Относительный путь: `/entity/product`
Цепочка вызовов от клиента будет выглядеть следующим образом:
```go
client := moysklad.NewClient()

// `/entity/product`
_ = client.Entity().Product()

// `/entity/variant`
_ = client.Entity().Variant()

// `/context/companysettings`
_ = client.Context().CompanySettings()

// `/report/dashboard`
_ = client.Report().Dashboard()
```
## Пример работы
```go
package main

import (
  "context"
  "fmt"
  "github.com/arcsub/go-moysklad/moysklad"
  "os"
)

func main() {
  // инициализируем простой клиент с аутентификацией по паре логин/пароль
  client := moysklad.NewClient().
	  WithBasicAuth(os.Getenv("MOYSKLAD_USERNAME"), os.Getenv("MOYSKLAD_PASSWORD")).
	  WithDisabledWebhookContent(true)

  // сервис для работы с товарами
  productService := client.Entity().Product()

  // выполняем запрос на получение списка товаров без дополнительных параметров (nil)
  products, _, err := productService.Get(context.Background(), nil)
  if err != nil {
    panic(err)
  }

  // выводим названия полученных товаров
  for _, product := range products.Rows {
    name := moysklad.Deref(product.Name) // Deref безопасно разыменовывает указатель
    fmt.Println(name)
  }

  // создадим новый товар
  product := new(moysklad.Product)

  // придумаем ему название (обязательное поле)
  product.Name = moysklad.String("Created Product")

  // отправим запрос на создание товара
  // в качестве аргументов передадим контекст, указатель на товар и nil в качестве параметров
  productCreated, _, err := productService.Create(context.Background(), product, nil)
  if err != nil {
    panic(err)
  }

  // выведем название созданного товара
  fmt.Println(moysklad.Deref(productCreated.Name))

  // изменим название товара
  productCreated.Name = moysklad.String("Updated Product")

  // отправим запрос на изменение товара
  // в качестве аргументов передадим контекст, указатель на Id изменяемой сущности, указатель на изменённый товар и nil в качестве параметров
  productUpdated, _, err := productService.Update(context.Background(), productCreated.Id, productCreated, nil)
  if err != nil {
    panic(err)
  }

  // выведем название изменённого товара
  fmt.Println(moysklad.Deref(productUpdated.Name))

  // отправим запрос на удаление товара
  // в качестве аргументов передадим контекст и указатель на Id удаляемой сущности
  success, _, err := productService.Delete(context.Background(), productUpdated.Id)
  if err != nil {
    panic(err)
  }

  // выведем состояние выполненного запроса, где true - успешно удалено, false – не удалено, см ошибки
  fmt.Println("Deleted", success)
}
```