package moysklad

// File Файл.
// Ключевое слово: files
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly
type File struct {
	Created   *Timestamp `json:"created,omitempty"`   // Время загрузки Файла на сервер
	CreatedBy *Employee  `json:"createdBy,omitempty"` // Метаданные сотрудника, загрузившего Файл
	Content   *string    `json:"content,omitempty"`   // Файл, закодированный в формате Base64.
	Filename  *string    `json:"filename,omitempty"`  // Имя Файла
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные объекта
	Miniature *Meta      `json:"miniature,omitempty"` // Метаданные миниатюры изображения (поле передается только для Файлов изображений)
	Size      *int       `json:"size,omitempty"`      // Размер Файла в байтах
	Tiny      *Meta      `json:"tiny,omitempty"`      // Метаданные уменьшенного изображения (поле передается только для Файлов изображений)
	Title     *string    `json:"title,omitempty"`     // Название Файла
}

func (file File) GetCreated() Timestamp {
	return Deref(file.Created)
}

func (file File) GetCreatedBy() Employee {
	return Deref(file.CreatedBy)
}

func (file File) GetContent() string {
	return Deref(file.Content)
}

func (file File) GetFilename() string {
	return Deref(file.Filename)
}

func (file File) GetMeta() Meta {
	return Deref(file.Meta)
}

func (file File) GetMiniature() Meta {
	return Deref(file.Miniature)
}

func (file File) GetSize() int {
	return Deref(file.Size)
}

func (file File) GetTiny() Meta {
	return Deref(file.Tiny)
}

func (file File) GetTitle() string {
	return Deref(file.Title)
}

func (file *File) SetContent(content string) *File {
	file.Content = &content
	return file
}

func (file *File) SetFilename(filename string) *File {
	file.Filename = &filename
	return file
}

func (file *File) SetMeta(meta *Meta) *File {
	file.Meta = meta
	return file
}

func (file File) String() string {
	return Stringify(file)
}

func (file File) MetaType() MetaType {
	return MetaTypeFiles
}
