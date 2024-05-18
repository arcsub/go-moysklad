package moysklad

// Image Изображение.
// Ключевое слово: image
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie
type Image struct {
	Content   *string    `json:"content,omitempty"`   // изображение, закодированное в Base64
	Filename  *string    `json:"filename,omitempty"`  // Имя файла
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные
	Miniature *Meta      `json:"miniature,omitempty"` // Метаданные миниатюры изображения
	Size      *int       `json:"size,omitempty"`      // Размер файла в байтах
	Tiny      *Meta      `json:"tiny,omitempty"`      // Метаданные уменьшенного изображения
	Title     *string    `json:"title,omitempty"`     // Название Изображения
	Updated   *Timestamp `json:"updated,omitempty"`   // Время загрузки файла на сервер
}

func (i Image) String() string {
	return Stringify(i)
}

func (i Image) MetaType() MetaType {
	return MetaTypeImage
}

type Images MetaArray[Image]

// Push добавляет элементы в срез.
// Элементы, превышающее максимальное значение MaxImages, игнорируются
func (i *Images) Push(elements ...*Image) *Images {
	i.Rows = append(i.Rows, elements...)
	if len(i.Rows) > MaxImages {
		i.Rows = i.Rows[:MaxImages]
	}
	return i
}
