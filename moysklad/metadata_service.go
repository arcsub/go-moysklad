package moysklad

// MetadataService
// Сервис для работы с метаданными.
type MetadataService struct {
	endpointGetOne[Metadata]
}

func NewMetadataService(client *Client) *MetadataService {
	e := NewEndpoint(client, "entity/metadata")
	return &MetadataService{
		endpointGetOne: endpointGetOne[Metadata]{e},
	}
}
