![](https://dev.moysklad.ru/doc/api/remap/1.2/images/logo-e9f672b5.svg)

[![](https://godoc.org/github.com/arcsub/go-moysklad?status.svg)](http://godoc.org/github.com/arcsub/go-moysklad)
[![Go Report Card](https://goreportcard.com/badge/github.com/arcsub/go-moysklad)](https://goreportcard.com/report/github.com/arcsub/go-moysklad)
![GitHub Tag](https://img.shields.io/github/v/tag/arcsub/go-moysklad?style=flat-square)
![GitHub License](https://img.shields.io/github/license/arcsub/go-moysklad?style=flat-square)
![GitHub commit activity](https://img.shields.io/github/commit-activity/t/arcsub/go-moysklad?style=flat-square)
![GitHub last commit](https://img.shields.io/github/last-commit/arcsub/go-moysklad?style=flat-square)

<div align="center">
  <h2>go-moysklad (МойСклад)</h2>
  <div>SDK для работы с <a href="https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api">МойСклад JSON API 1.2</a></div>
  <br>
  <img src="https://github.com/arcsub/go-moysklad/assets/47686389/6bec5834-6eb9-442f-b1ee-efeaa85bb946" width="200px">
</div>

>
> **SDK находится в стадии разработки!**
>
> Некоторые методы могут отсутствовать или работать неправильно!
>
> **Подробная документация в процессе написания.**

## Установка

> Требуемая версия go >= 1.21

```
go get -u github.com/arcsub/go-moysklad@HEAD
```

## Особенности

### Используется [resty](https://pkg.go.dev/github.com/go-resty/resty/v2) как стандартный клиент

Можно самостоятельно настроить resty-клиент, изучив его [документацию](https://pkg.go.dev/github.com/go-resty/resty/v2).

### Возвращаемые аргументы
Каждый запрос на создание/изменение/удаление возвращает 3 аргумента.
Рассмотрим объявление функции
```go
func (s *endpointCreate[T]) Create(ctx context.Context, entity *T, params ...func (*Params)) (*T, *resty.Response, error)
```
В примере выше нас интересуют возвращаемые аргументы: `(*T, *resty.Response, error)`
1. `*T` – указатель на сущность/документ, например *Product при вызове `Create()` (возвращает `bool` при вызове метода `Delete()`).
2. `*resty.Response` – ответ на запрос, содержащий *http.Response и некоторую другую информацию.
3. `error` – ошибки, если они были. При возникновении ошибок от API МойСклад в качестве ошибки будет заполненная структура `ApiErrors`

### Указатели
Поля структур сущностей и документов являются указателями.

- Чтобы получить значение по указателю необходимо вызвать метод структуры `GetFieldName()`.
  - `FieldName` - наименование поля.

Например:
```go
name := product.GetName()
id := product.GetID()
```

- Чтобы установить значение необходимо передать значение в соответствующий метод `SetFieldName(value)`
  - `FieldName` - наименование поля
  - `value` - передаваемое значение.

> [!NOTE]
> Методы `SetFieldName()` возвращают указатель на объект, что позволяет вызывать методы по цепочке.


Например:
```go
product := new(moysklad.Product)
product.SetName("iPhone 16 Pro Max").SetCode("APPL16PM")
```

- ~~Для безопасного разыменовывания указателя необходимо передать указатель в метод `Deref()`~~
- ~~Чтобы установить указатель на примитивное значение поля также существуют вспомогательные методы:~~
  - ~~`Bool()` возвращает *bool~~
  - ~~`Int()` возвращает *int~~
  - ~~`Uint()` возвращает *uint64~~
  - ~~`Float()` возвращает *float64~~
  - ~~`String()` возвращает *string~~

## Использование
### Создание экземпляра клиента
```go
// создание стандартного клиента (на базе resty.New())
client := moysklad.New(moysklad.Config{
  // с использованием токена
  Token: os.Getenv("MOYSKLAD_TOKEN"),
  
  // или с использованием логина и пароля
  Username: os.Getenv("MOYSKLAD_USERNAME"),
  Password: os.Getenv("MOYSKLAD_PASSWORD"),
})
```

### Создание экземпляра клиента со своим http клиентом

```go
httpClient := &http.Client{Timeout: 5 * time.Minute}

client := moysklad.New(moysklad.Config{
  Token: os.Getenv("MOYSKLAD_TOKEN"),
  HTTPClient: httpClient,
})
```

### Создание экземпляра клиента с resty клиентом

```go
restyClient := resty.New()
restyClient.SetRetryCount(10) // количество повторных попыток

client := moysklad.New(moysklad.Config{
  Token: os.Getenv("MOYSKLAD_TOKEN"),
  RestyClient: restyClient,
})
```
### Параметры запроса

#### Пример передачи параметров запроса в метод

Функциональные параметры запроса можно передавать в методы, сигнатура которых это предусматривают.

```go
list, _, err := client.Entity().Product().GetList(context.Background(),
  moysklad.WithExpand("country"),
  moysklad.WithOrder("name"),
  moysklad.WithLimit(10),
)
```

---
#### Количество элементов на странице `limit=val`
Пример:
```go
moysklad.WithLimit(100)
```

#### Смещение от первого элемента `offset=val`
Пример:
```go
moysklad.WithOffset(100)
```

#### Контекстный поиск `search=val`
Пример:
```go
moysklad.WithSearch("iPhone 16 Pro Max")
```
#### Замена ссылок объектами
Пример:
```go
moysklad.WithExpand("positions").WithExpand("group")
```

#### Фильтрация по значению `key=value`
Пример:
```go
moysklad.WithFilterEquals("name", "Яблоко")
```

#### Строго больше `key>value`
Пример:
```go
moysklad.WithFilterGreater("sum", "100")
```

#### Строго меньше `key<value`
Пример:
```go
moysklad.WithFilterLesser("sum", "1000")
```

#### Больше или равно `key=>value`
Пример:
```go
moysklad.WithFilterGreaterOrEquals("moment", "2023-06-01")
```

#### Меньше или равно `key<=value`
Пример:
```go
moysklad.WithFilterLesserOrEquals("moment", "2023-06-01")
```

#### Не равно `key!=value`
Пример:
```go
moysklad.WithFilterNotEquals("name", "0001")
```

#### Частичное совпадение (обычное подобие) `key~value`
Пример:
```go
moysklad.WithFilterEquivalence("code", "ms")
```

#### Полное совпадение в начале значения (левое подобие) `key~=value`
Пример:
```go
moysklad.WithFilterEquivalenceLeft("code", "ms")
```

#### Полное совпадение в конце значения (правое подобие) `key=~value`
Пример:
```go
moysklad.WithFilterEquivalenceRight("code", "ms")
```

#### Частичное совпадение не выводится `key!~value`
Пример:
```go
moysklad.WithFilterNotEquivalence("code", "ms")
```

#### Фильтрация по удалённым документам `isDeleted=val`
Пример:
```go
moysklad.WithFilterDeleted(true)
```

#### Фильтрация по напечатанным документам `printed=val`
Пример:
```go
moysklad.WithFilterPrinted(true)
```

#### Фильтрация по опубликованным документам `published=val`
Пример:
```go
moysklad.WithFilterPublished(true)
```

#### Фильтрация по архивным сущностям `archived=val`
Пример:
```go
moysklad.WithFilterArchived(true)
```

#### Группировка выдачи `groupBy=val`
Пример:
```go
moysklad.WithGroupBy(moysklad.GroupByProduct)
```

#### Применение сохранённого фильтра `namedFilter=href`
Метод принимает указатель на сохранённый фильтр.

Пример:
```go
moysklad.WithNamedFilter(&NamedFilter{...})
```

#### Сортировка по умолчанию `order=fieldName`
Пример:
```go
moysklad.WithOrder("name")
```

#### Сортировка по возрастанию `order=fieldName,asc`
Пример:
```go
moysklad.WithOrderAsc("name")
```

#### Сортировка по убыванию `order=fieldName,desc`
Пример:
```go
moysklad.WithOrderDesc("name")
```

#### Остатки и себестоимость в позициях документов `fields=stock`
Метод устанавливает лимит позиций в 100 единиц, а также применяет замену ссылок объектами для поля `positions`

Пример:
```go
moysklad.WithStockFiled()
```

#### Тип остатка `stockType=val`

Используется в отчёте "Остатки"

Пример:
```go
moysklad.WithStockType(moysklad.StockDefault)
```

#### Интервал, с которым будет построен отчет `interval=val`
Используется в отчётах

Пример:
```go
moysklad.WithInterval(moysklad.IntervalMonth)
```

#### Начало периода `momentFrom=val`
Метод принимает `time.Time`
Пример:
```go
moysklad.WithMomentFrom(time.Now())
```

#### Конец периода `momentTo=val`
Метод принимает `time.Time`
Пример:
```go
moysklad.WithMomentTo(time.Now())
```

### Сервисы
Для перехода к определённому сервису необходимо вызвать цепочку методов, аналогично пути запроса.

#### Пример: ProductService
Сервис для работы с товарами.

Относительный путь: `/entity/product`
Цепочка вызовов от клиента будет выглядеть следующим образом:
```go
// `/entity/product`
_ = client.Entity().Product()

// `/entity/variant`
_ = client.Entity().Variant()

// `/context/companysettings`
_ = client.Context().CompanySettings()

// `/report/dashboard`
_ = client.Report().Dashboard()
```

### Запрос по объекту `Meta`

Если возникает необходимость точечно запросить информацию о сущности, имея только её `Meta`, можно использовать
метод `FetchMeta`.

Чтобы использовать данный функционал необходимо точно знать, какой тип данных мы ожидаем получить в ответ.

Метод имеет следующую сигнатуру:

```go
func FetchMeta[T any](ctx context.Context, client *Client, meta Meta, params ...func (*Params)) (*T, *resty.Response, error)
```

Пример:

```go
product, _, _ := moysklad.FetchMeta[moysklad.Product](ctx, client, product.GetMeta())
```
### Пример работы
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
  client := moysklad.New(moysklad.Config{
    Username:               os.Getenv("MOYSKLAD_USERNAME"),
    Password:               os.Getenv("MOYSKLAD_PASSWORD"),
    DisabledWebhookContent: true,
  })

  // сервис для работы с товарами
  productService := client.Entity().Product()

  // выполняем запрос на получение списка товаров без дополнительных параметров
  products, _, err := productService.GetList(context.Background())
  if err != nil {
    panic(err)
  }

  // выводим названия полученных товаров
  for _, product := range products.Rows {
    fmt.Println(product.GetName())
  }

  // создадим новый товар
  product := new(moysklad.Product)
  product.SetName("Created Product")

  // отправим запрос на создание товара
  // в качестве аргументов передадим контекст и товар
  productCreated, _, err := productService.Create(context.Background(), product, 
    moysklad.WithExpand("country", "owner", "productFolder"), // пример передачи параметров запроса
  )
  if err != nil {
    panic(err)
  }

  // выведем название созданного товара
  fmt.Println(productCreated.GetName())

  // изменим название товара
  productCreated.SetName("Updated Product")

  // отправим запрос на изменение товара
  // в качестве аргументов передадим контекст, ID изменяемой сущности и изменённый товар
  productUpdated, _, err := productService.Update(context.Background(), productCreated.GetID(), productCreated)
  if err != nil {
    panic(err)
  }

  // выведем название изменённого товара
  fmt.Println(productUpdated.GetName())

  // отправим запрос на удаление товара
  // в качестве аргументов передадим контекст и удаляемую сущность, в которой содержится идентификатор
  success, _, err := productService.Delete(context.Background(), productUpdated)
  if err != nil {
    panic(err)
  }

  // выведем состояние выполненного запроса, где true - успешно удалено, false – не удалено, см ошибки
  fmt.Println("Deleted", success)
}
```

## Обратная связь

Буду рад видеть ваши идеи и предложения в [Issues](https://github.com/arcsub/go-moysklad/issues).

<!--div align="center">
  <a href="https://pay.cloudtips.ru/p/eac3797c" target="_blank">
  <img src="https://github.com/arcsub/go-moysklad/assets/47686389/6431baa5-28e4-48b6-8d97-0e50bb6646d2" width="150px">
<div align="center">Поддержать проект</div>
</a>
</div-->
<br/>