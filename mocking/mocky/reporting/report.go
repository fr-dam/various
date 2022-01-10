package reporting

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"google.golang.org/api/iterator"

)

type ReportService struct {
	ctx       context.Context
	projectID string
	gcs       StorageClienter
}

type ReportWriter interface {
	Write(ctx context.Context, report []byte) error
	Read(ctx context.Context, name string) ([]byte, error)
	List(ctx context.Context) map[string][]byte
}

func (s ReportService) WriteReport(ctx context.Context, metadata map[string]string, report []byte) error {
	reportName := fmt.Sprintf("%d", time.Now().Unix())
	writer := s.gcs.GetBucketWriter(ctx, reportName)
	writer.ContentType = "application/json"
	writer.Metadata = metadata
	_, err := io.Copy(writer, bytes.NewReader(report))
	return err
}

func (s ReportService) ReadReport(ctx context.Context, name string) ([]byte, error) {
	reader, err := s.gcs.GetBucketReader(ctx, name)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}

func (s ReportService) ListReports(ctx context.Context) (map[string]map[string]string, error) {
	reports := make(map[string]map[string]string)
	it := s.gcs.GetBucketObjectIterator(ctx)
	for {
		next, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		reports[next.Name] = next.Metadata
	}
	return reports, nil
}
