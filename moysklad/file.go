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

func (f File) String() string {
	return Stringify(f)
}

func (f File) MetaType() MetaType {
	return MetaTypeFiles
}

type Files MetaArray[File]

// Push добавляет элементы в срез.
// Элементы, превышающее максимальное значение MaxImages, игнорируются
func (f *Files) Push(elements ...*File) *Files {
	limit := min(MaxFiles, MaxImages-len(f.Rows), len(elements))
	f.Rows = append(f.Rows, elements[:limit]...)
	return f
}
