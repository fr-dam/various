package reporting

import (
	"context"
)

type ReportWriter interface {
	Write(ctx context.Context) error
}
//
//type ReportService struct {
//	ctx       context.Context
//	projectID string
//	client *storage.Client
//	gcs    ReportWriter
//}
//
//func NewReportService(ctx context.Context, projectID string) *ReportService {
//	gcs, err := storage.NewClient(ctx)
//	if err != nil {
//		panic(err)
//	}
//	return &ReportService{
//		ctx: ctx,
//		projectID: projectID,
//		//gcs: gcs,
//		client: gcs,
//	}
//}

func Write(ctx context.Context, report map[string]interface{}) error {
	_ = report
	return nil
}
