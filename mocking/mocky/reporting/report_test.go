package reporting

import (
	"context"
	"github.com/fr-dam/various/mocking/mocky/reporting/mocks"
	"testing"

	//"cloud.google.com/go/storage"
	"github.com/bxcodec/faker"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type reportTestSuite struct {
	suite.Suite
	ctx   context.Context
	log   *logrus.Entry
	lorem faker.Lorem
}

func TestReportSuite(t *testing.T) {
	r := &reportTestSuite{
		ctx:   context.Background(),
		log:   logrus.WithField("package", "reporting").WithField("suite", "reportTestSuite"),
		lorem: faker.Lorem{},
	}
	suite.Run(t, r)
}

func (r *reportTestSuite) TestWriteReports() {
	type args struct {
		metadata map[string]string
		report   []byte
	}
	type MockOp struct {
		method       string
		args         []interface{}
		params       []interface{}
		returnValues []interface{}
	}
	var tests = []struct {
		name           string
		bucketName     string
		args           args
		reportOps      []MockOp
		expectedResult bool
		expectedError  error
	}{
		{
			name:       "Successful write",
			bucketName: "test-bucket",
			args: args{
				metadata: map[string]string{
					r.lorem.Word(): r.lorem.Word(),
				},
				report: []byte(r.lorem.Word()),
			},
			reportOps: []MockOp{
				{
					method: "GetBucketWriter",
					params: []interface{}{
						mock.AnythingOfType("*context.emptyCtx"),
						mock.AnythingOfType("string"),
					},
					returnValues: []interface{}{
						mock.AnythingOfType("*storage.Writer"),
					},
				},
			},
			expectedResult: true,
			expectedError:  nil,
		},
	}

	for _, t := range tests {
		r.T().Log(t.name)
		storageClienter := &mocks.StorageClienter{}
		mocks := []*mock.Mock{&storageClienter.Mock}
		mockOps := [][]MockOp{t.reportOps}
		r.T().Logf("Found %d mocks", len(mocks))
		for i := range mocks {
			for _, op := range mockOps[i] {
				r.T().Logf("method: %s, params: %+v", op.method, op.params)
				mocks[i].On(op.method, op.params...).Return(op.returnValues...)
			}
		}

		r.T().Log("HERE 1")
		svc := ReportService{ctx: r.ctx, projectID: "fr-1234", gcs: storageClienter}

		r.T().Log("HERE 2")
		actualErr := svc.WriteReport(r.ctx, t.args.metadata, t.args.report)
		r.T().Log(actualErr)
		r.T().Log("HERE 3")
	}
}
