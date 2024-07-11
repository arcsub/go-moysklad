package moysklad

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"reflect"

	"io"
	"os"
	"path/filepath"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Uint is a helper routine that allocates a new uint64 value
// to store v and returns a pointer to it.
func Uint(v uint64) *uint64 { return &v }

// Float is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Stringify attempts to create a reasonable string representation of types in
// the Moysklad library. It does things like resolve pointers to their values
// and omits struct fields with nil values.
func Stringify(message any) string {
	var buf bytes.Buffer
	v := reflect.ValueOf(message)
	stringifyValue(&buf, v)
	return buf.String()
}

func stringifyValue(w io.Writer, val reflect.Value) {
	if val.Kind() == reflect.Ptr && val.IsNil() {
		w.Write([]byte("<nil>"))
		return
	}

	v := reflect.Indirect(val)

	switch v.Kind() {
	case reflect.String:
		fmt.Fprintf(w, `"%s"`, v)
	case reflect.Slice:
		w.Write([]byte{'['})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte{' '})
			}

			stringifyValue(w, v.Index(i))
		}

		w.Write([]byte{']'})
		return
	case reflect.Struct:
		if v.Type().Name() != "" {
			w.Write([]byte(v.Type().String()))
		}

		// special handling of Timestamp values
		if v.Type() == reflect.TypeOf(Timestamp{}) {
			fmt.Fprintf(w, "{%s}", v.Interface())
			return
		}

		w.Write([]byte{'{'})

		var sep bool
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			if fv.Kind() == reflect.Ptr && fv.IsNil() {
				continue
			}
			if fv.Kind() == reflect.Slice && fv.IsNil() {
				continue
			}
			if fv.Kind() == reflect.Map && fv.IsNil() {
				continue
			}

			if sep {
				w.Write([]byte(", "))
			} else {
				sep = true
			}

			w.Write([]byte(v.Type().Field(i).Name))
			w.Write([]byte{':'})
			stringifyValue(w, fv)
		}

		w.Write([]byte{'}'})
	default:
		if v.CanInterface() {
			fmt.Fprint(w, v.Interface())
		}
	}
}

// Clamp задаёт значение в диапазоне между указанными нижней и верхней границами.
func Clamp(val, min, max int) int {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

func getFilenameContent(filePath string) (string, string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", "", err
	}
	return filepath.Base(filePath), base64.StdEncoding.EncodeToString(b), nil
}

func getContentFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bodyBytes), nil
}

// RawMetaTyper описывает методы, необходимые для преобразования одного типа в другой.
type RawMetaTyper interface {
	MetaTyper
	Raw() []byte
}

// filterType преобразует слайс элементов типа D в слайс элементов типа M
func filterType[M MetaTyper, D RawMetaTyper](elements []D) Slice[M] {
	var slice = Slice[M]{}
	for _, el := range elements {
		if e := UnmarshalAsType[M](el); e != nil {
			slice.Push(e)
		}
	}
	return slice
}

// UnmarshalAsType принимает объект D, реализующий интерфейс [RawMetaTyper] и приводит его к типу M.
//
// Возвращает nil в случае неудачи или при несоответствии типов [MetaType].
func UnmarshalAsType[M MetaTyper, D RawMetaTyper](element D) *M {
	var t = *new(M)

	if t.MetaType() != element.MetaType() {
		return nil
	}

	data := element.Raw()
	if data == nil {
		return nil
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return nil
	}

	return &t
}

// UnmarshallAny принимает любой тип data, сериализует и пытается десериализовать в тип T.
func UnmarshallAny[T any](data any) (T, error) {
	var t T
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return t, err
	}

	if err = json.Unmarshal(b, &t); err != nil {
		log.Println(err)
		return t, err
	}

	return t, nil
}

// Deref разыменовывает указатель
func Deref[T any](ptr *T) T {
	if ptr == nil {
		var v T
		return v
	}
	return *ptr
}

// IsEqualPtr сравнивает значения указателей типа T
func IsEqualPtr[T comparable](l *T, r *T) bool {
	return l != nil && r != nil && Deref(l) == Deref(r)
}

// IsMetaEqual сравнивает `meta.href` двух сущностей типа *T
func IsMetaEqual[T MetaOwner](l *T, r *T) bool {
	lMeta := Deref(l).GetMeta()
	rMeta := Deref(r).GetMeta()
	return l != nil && r != nil && lMeta.IsEqual(&rMeta)
}

// NullValue тип для поля, которое может быть указано как null.
// Имеет обобщённый тип T.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-podderzhka-null
type NullValue[T any] struct {
	value *T   // значение поля
	null  bool // признак null
}

// NewNullValue устанавливает значение поля value.
//
// Устанавливает null при передаче nil в качестве аргумента.
func NewNullValue[T any](value *T) *NullValue[T] {
	if value == nil {
		return &NullValue[T]{null: true}
	}

	// Проверяем, есть ли у объекта метод .Clean(), и вызываем его
	if method := reflect.ValueOf(value).MethodByName("Clean"); method.IsValid() {
		if cleanFn, ok := method.Interface().(func() *T); ok {
			return &NullValue[T]{value: cleanFn()}
		}
	}

	return &NullValue[T]{value: value}
}

// IsNull возвращает true, если объект не инициализирован или поле null.
func (nullValue *NullValue[T]) IsNull() bool {
	return nullValue == nil || nullValue.isNull()
}

// isNull возвращает true, если поле null.
func (nullValue NullValue[T]) isNull() bool {
	return nullValue.null || nullValue.value == nil
}

// getValue возвращает значение поля.
func (nullValue NullValue[T]) getValue() T {
	return Deref(nullValue.value)
}

// setValue устанавливает значение поля.
func (nullValue *NullValue[T]) setValue(value *T) *NullValue[T] {
	if value == nil {
		nullValue.setNull()
	} else {
		nullValue.value = value
		nullValue.null = false
	}
	return nullValue
}

// setNull устанавливает значение поля null.
func (nullValue *NullValue[T]) setNull() *NullValue[T] {
	nullValue.value = nil
	nullValue.null = true
	return nullValue
}

// String реализует интерфейс [fmt.Stringer].
func (nullValue NullValue[T]) String() string {
	return fmt.Sprintf("%v", nullValue.value)
}

// MarshalJSON реализует интерфейс [json.Marshaler].
func (nullValue NullValue[T]) MarshalJSON() ([]byte, error) {
	if nullValue.isNull() {
		return json.Marshal(nil)
	}
	return json.Marshal(nullValue.getValue())
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
func (nullValue *NullValue[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &nullValue.value)
}

// NullValueAny тип для поля Value структуры [Attribute].
type NullValueAny struct {
	value any  // значение поля
	null  bool // признак null
}

// NewNullValueAny устанавливает значение поля null.
func NewNullValueAny() *NullValueAny {
	return &NullValueAny{null: true}
}

// NewNullValueAnyFrom устанавливает значение value.
func NewNullValueAnyFrom(value any) *NullValueAny {
	return &NullValueAny{value: value}
}

// IsNull возвращает true, если поле null.
func (nullValueAny NullValueAny) IsNull() bool {
	return nullValueAny.null
}

// Get возвращает значение поля.
func (nullValueAny NullValueAny) Get() any {
	return nullValueAny.value
}

// Set устанавливает значение поля.
func (nullValueAny *NullValueAny) Set(value any) *NullValueAny {
	nullValueAny.value = value
	return nullValueAny
}

// SetNull устанавливает значение поля null.
func (nullValueAny *NullValueAny) SetNull() *NullValueAny {
	nullValueAny.null = true
	return nullValueAny
}

// String реализует интерфейс [fmt.Stringer].
func (nullValueAny NullValueAny) String() string {
	return fmt.Sprintf("%v", nullValueAny.value)
}

// MarshalJSON реализует интерфейс [json.Marshaler].
func (nullValueAny NullValueAny) MarshalJSON() ([]byte, error) {
	if nullValueAny.IsNull() {
		return json.Marshal(nil)
	}
	return json.Marshal(nullValueAny.Get())
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
func (nullValueAny *NullValueAny) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &nullValueAny.value)
}

// AsMetaWrapperSlice оборачивает каждый элемент среза в объект [MetaWrapper] и возвращает срез объектов [MetaWrapper].
func AsMetaWrapperSlice[T MetaOwner](entities []*T) []MetaWrapper {
	var o = make([]MetaWrapper, 0, len(entities))
	for _, entity := range entities {
		if entity != nil {
			o = append(o, (*entity).GetMeta().Wrap())
		}
	}
	return o
}

// CheckType сопоставляет код сущности entity со значением [MetaType].
func CheckType[T MetaOwner](entity T, metaType MetaType) bool {
	return entity.GetMeta().GetType() == metaType
}
