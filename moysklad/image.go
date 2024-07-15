package moysklad

import "time"

// Image Изображение.
//
// Код сущности: image
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie
type Image struct {
	Content   *string    `json:"content,omitempty"`   // изображение, закодированное в Base64
	Filename  *string    `json:"filename,omitempty"`  // Имя файла
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные объекта
	Miniature *Meta      `json:"miniature,omitempty"` // Метаданные миниатюры изображения
	Size      *int       `json:"size,omitempty"`      // Размер файла в байтах
	Tiny      *Meta      `json:"tiny,omitempty"`      // Метаданные уменьшенного изображения
	Title     *string    `json:"title,omitempty"`     // Название Изображения
	Updated   *Timestamp `json:"updated,omitempty"`   // Время загрузки файла на сервер
}

// GetContent возвращает изображение, закодированное в Base64.
func (image Image) GetContent() string {
	return Deref(image.Content)
}

// GetFilename возвращает Имя файла.
func (image Image) GetFilename() string {
	return Deref(image.Filename)
}

// GetMeta возвращает Метаданные объекта.
func (image Image) GetMeta() Meta {
	return Deref(image.Meta)
}

// GetMiniature возвращает Метаданные миниатюры изображения.
func (image Image) GetMiniature() Meta {
	return Deref(image.Miniature)
}

// GetSize возвращает Размер файла в байтах.
func (image Image) GetSize() int {
	return Deref(image.Size)
}

// GetTiny возвращает Метаданные уменьшенного изображения.
func (image Image) GetTiny() Meta {
	return Deref(image.Tiny)
}

// GetTitle возвращает Название Изображения.
func (image Image) GetTitle() string {
	return Deref(image.Title)
}

// GetUpdated возвращает Время загрузки файла на сервер.
func (image Image) GetUpdated() time.Time {
	return Deref(image.Updated).Time()
}

// SetContent устанавливает изображение, закодированное в Base64.
func (image *Image) SetContent(content string) *Image {
	image.Content = &content
	return image
}

// SetFilename устанавливает Имя файла.
func (image *Image) SetFilename(filename string) *Image {
	image.Filename = &filename
	return image
}

// SetMeta устанавливает Метаданные объекта.
func (image *Image) SetMeta(meta *Meta) *Image {
	image.Meta = meta
	return image
}

// String реализует интерфейс [fmt.Stringer].
func (image Image) String() string {
	return Stringify(image)
}

// MetaType возвращает код сущности.
func (Image) MetaType() MetaType {
	return MetaTypeImage
}

// NewImageFromURL принимает URL путь до изображения и возвращает [Image].
func NewImageFromURL(url string) (*Image, error) {
	content, err := getContentFromURL(url)
	if err != nil {
		return nil, err
	}
	return &Image{Content: String(content)}, nil
}

// NewImageFromFilepath принимает путь до изображения и возвращает [Image].
func NewImageFromFilepath(filePath string) (*Image, error) {
	fileName, content, err := getFilenameContent(filePath)
	if err != nil {
		return nil, err
	}

	image := &Image{
		Title:    String(fileName),
		Filename: String(fileName),
		Content:  String(content),
	}
	return image, nil
}
