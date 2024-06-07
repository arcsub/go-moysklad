package moysklad

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"reflect"

	"image/color"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

// Clamp задаёт значение в диапазоне между указанными нижней и верхней границами
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

type FileTypes interface {
	File | Image
}

// NewFileFromFilepath возвращает *File, на основе переданного пути до файла
// и ошибку, если файл не удалось найти
func NewFileFromFilepath(filePath string) (*File, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(filePath)
	content := base64.StdEncoding.EncodeToString(b)
	f := &File{
		Title:    String(fileName),
		Filename: String(fileName),
		Content:  String(content),
	}
	return f, nil
}

// NewImageFromFilepath возвращает *Image, на основе переданного пути до файла
// и ошибку, если файл не удалось найти
func NewImageFromFilepath(filePath string) (*Image, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(filePath)
	content := base64.StdEncoding.EncodeToString(b)
	f := &Image{
		Title:    String(fileName),
		Filename: String(fileName),
		Content:  String(content),
	}
	return f, nil
}

type RawMetaTyper interface {
	MetaTyper
	Raw() json.RawMessage
}

func filterType[M MetaTyper, D RawMetaTyper](elements []D) Slice[M] {
	var slice = Slice[M]{}
	for _, el := range elements {
		if e := unmarshalAsType[M](el); e != nil {
			slice.Push(e)
		}
	}
	return slice
}

// unmarshalAsType принимает объект, удовлетворяющий интерфейсу RawMetaTyper
// и структурирует в тип T
func unmarshalAsType[M MetaTyper, D RawMetaTyper](element D) *M {
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

type Interval string

func (i Interval) String() string {
	return string(i)
}

const (
	IntervalHour  Interval = "hour"
	IntervalDay   Interval = "day"
	IntervalMonth Interval = "month"
)

// RGBtoUint64 конвертирует код цвета из формата RRGGBB и RGB в uint64
// Example
// RGBtoUint64("#E6E6E6")
// RGBtoUint64("e3e3e3")
// RGBtoUint64("FFF")
func RGBtoUint64(str string) (uint64, error) {
	var errInvalid = errors.New("invalid format")
	var err error

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}

		err = errInvalid

		return 0
	}

	n := strings.Replace(str, "0x", "", -1)
	n = strings.Replace(n, "0X", "", -1)
	n = strings.Replace(n, "#", "", -1)

	c := color.RGBA{A: 0xff}

	switch len(n) {
	case 6: // RRGGBB
		c.R = hexToByte(n[0])<<4 + hexToByte(n[1])
		c.G = hexToByte(n[2])<<4 + hexToByte(n[3])
		c.B = hexToByte(n[4])<<4 + hexToByte(n[5])
	case 3: // RGB
		c.R = hexToByte(n[0]) * 17
		c.G = hexToByte(n[1]) * 17
		c.B = hexToByte(n[2]) * 17
	default: // invalid format
		err = errInvalid
	}

	if err != nil {
		return 0, err
	}

	hex := fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B)
	output, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	return output, nil
}

var reContentDisposition = regexp.MustCompile(`filename="(.*)"`)

func GetFileFromResponse(resp *resty.Response) (*PrintFile, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.RawBody()); err != nil {
		return nil, err
	}

	var fileName string
	headerStr := resp.Header().Get(headerContentDisposition)
	if match := reContentDisposition.FindStringSubmatch(headerStr); len(match) > 1 {
		fileName = match[1]
	}
	file := &PrintFile{buf, fileName}
	return file, nil
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

type DeleteManyRequest []MetaWrapper

func NewDeleteManyRequest() DeleteManyRequest {
	return make(DeleteManyRequest, 0)
}

// Push добавляет элементы MetaOwner в конец среза.
func (s *DeleteManyRequest) Push(entities ...MetaOwner) *DeleteManyRequest {
	for _, entity := range entities {
		*s = append(*s, MetaWrapper{entity.GetMeta()})
	}
	return s
}

// PushMeta добавляет элементы Meta в конец среза.
func (s *DeleteManyRequest) PushMeta(entities ...Meta) *DeleteManyRequest {
	for _, entity := range entities {
		*s = append(*s, MetaWrapper{entity})
	}
	return s
}
