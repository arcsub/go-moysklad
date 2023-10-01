package moysklad

// CommissionReportOutService
// Сервис для работы с выданными отчётами комиссионера.
type CommissionReportOutService struct {
	endpointGetList[CommissionReportOut]
	endpointCreate[CommissionReportOut]
	endpointCreateUpdateDeleteMany[CommissionReportOut]
	endpointDelete
	endpointGetById[CommissionReportOut]
	endpointUpdate[CommissionReportOut]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[CommissionReportOutPosition]
	endpointAttributes
	endpointSyncId[CommissionReportOut]
	endpointNamedFilter
	endpointTemplate[CommissionReportOut]
	endpointPublication
	endpointRemove
}

func NewCommissionReportOutService(client *Client) *CommissionReportOutService {
	e := NewEndpoint(client, "entity/commissionreportout")
	return &CommissionReportOutService{
		endpointGetList:                endpointGetList[CommissionReportOut]{e},
		endpointCreate:                 endpointCreate[CommissionReportOut]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[CommissionReportOut]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[CommissionReportOut]{e},
		endpointUpdate:                 endpointUpdate[CommissionReportOut]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[CommissionReportOutPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointSyncId:                 endpointSyncId[CommissionReportOut]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointTemplate:               endpointTemplate[CommissionReportOut]{e},
		endpointPublication:            endpointPublication{e},
		endpointRemove:                 endpointRemove{e},
	}
}
