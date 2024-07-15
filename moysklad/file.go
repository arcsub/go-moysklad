package moysklad

import "time"

// File Файл.
//
// Код сущности: files
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly
type File struct {
	Created   *Timestamp `json:"created,omitempty"`   // Время загрузки Файла на сервер
	CreatedBy *Employee  `json:"createdBy,omitempty"` // Метаданные сотрудника, загрузившего Файл
	Content   *string    `json:"content,omitempty"`   // Файл, закодированный в формате Base64
	Filename  *string    `json:"filename,omitempty"`  // Имя Файла
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные файла
	Miniature *Meta      `json:"miniature,omitempty"` // Метаданные миниатюры изображения (поле передается только для Файлов изображений)
	Size      *int       `json:"size,omitempty"`      // Размер Файла в байтах
	Tiny      *Meta      `json:"tiny,omitempty"`      // Метаданные уменьшенного изображения (поле передается только для Файлов изображений)
	Title     *string    `json:"title,omitempty"`     // Название Файла
}

// GetCreated возвращает Время загрузки Файла на сервер.
func (file File) GetCreated() time.Time {
	return Deref(file.Created).Time()
}

// GetCreatedBy возвращает Метаданные сотрудника, загрузившего Файл.
func (file File) GetCreatedBy() Employee {
	return Deref(file.CreatedBy)
}

// GetContent возвращает Файл, закодированный в формате Base64.
func (file File) GetContent() string {
	return Deref(file.Content)
}

// GetFilename возвращает Имя Файла.
func (file File) GetFilename() string {
	return Deref(file.Filename)
}

// GetMeta возвращает Метаданные файла.
func (file File) GetMeta() Meta {
	return Deref(file.Meta)
}

// GetMiniature возвращает Метаданные миниатюры изображения (поле передается только для Файлов изображений).
func (file File) GetMiniature() Meta {
	return Deref(file.Miniature)
}

// GetSize возвращает Размер Файла в байтах.
func (file File) GetSize() int {
	return Deref(file.Size)
}

// GetTiny возвращает Метаданные уменьшенного изображения (поле передается только для Файлов изображений).
func (file File) GetTiny() Meta {
	return Deref(file.Tiny)
}

// GetTitle возвращает Название Файла.
func (file File) GetTitle() string {
	return Deref(file.Title)
}

// SetContent устанавливает Файл, закодированный в формате Base64.
func (file *File) SetContent(content string) *File {
	file.Content = &content
	return file
}

// SetFilename устанавливает Имя Файла.
func (file *File) SetFilename(filename string) *File {
	file.Filename = &filename
	return file
}

// SetMeta устанавливает Метаданные Файла.
func (file *File) SetMeta(meta *Meta) *File {
	file.Meta = meta
	return file
}

// String реализует интерфейс [fmt.Stringer].
func (file File) String() string {
	return Stringify(file)
}

// MetaType возвращает код сущности.
func (File) MetaType() MetaType {
	return MetaTypeFiles
}

// NewFileFromURL принимает URL путь до файла и возвращает [File].
func NewFileFromURL(url string) (*File, error) {
	content, err := getContentFromURL(url)
	if err != nil {
		return nil, err
	}
	return &File{Content: String(content)}, nil
}

// NewFileFromFilepath принимает путь до файла и возвращает [File].
func NewFileFromFilepath(filePath string) (*File, error) {
	fileName, content, err := getFilenameContent(filePath)
	if err != nil {
		return nil, err
	}

	file := &File{
		Title:    String(fileName),
		Filename: String(fileName),
		Content:  String(content),
	}
	return file, nil
}
