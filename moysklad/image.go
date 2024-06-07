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

func (image Image) GetContent() string {
	return Deref(image.Content)
}

func (image Image) GetFilename() string {
	return Deref(image.Filename)
}

func (image Image) GetMeta() Meta {
	return Deref(image.Meta)
}

func (image Image) GetMiniature() Meta {
	return Deref(image.Miniature)
}

func (image Image) GetSize() int {
	return Deref(image.Size)
}

func (image Image) GetTiny() Meta {
	return Deref(image.Tiny)
}

func (image Image) GetTitle() string {
	return Deref(image.Title)
}

func (image Image) GetUpdated() Timestamp {
	return Deref(image.Updated)
}

func (image *Image) SetContent(content string) *Image {
	image.Content = &content
	return image
}

func (image *Image) SetFilename(filename string) *Image {
	image.Filename = &filename
	return image
}

func (image *Image) SetMeta(meta *Meta) *Image {
	image.Meta = meta
	return image
}

func (image Image) String() string {
	return Stringify(image)
}

func (image Image) MetaType() MetaType {
	return MetaTypeImage
}

type Images MetaArray[Image]

// Push добавляет элементы в срез.
// Элементы, превышающее максимальное значение MaxImages, игнорируются
func (i *Images) Push(elements ...*Image) *Images {
	i.Rows.Push(elements...)
	if i.Rows.Len() > MaxImages {
		i.Rows = i.Rows[:MaxImages]
	}
	return i
}
