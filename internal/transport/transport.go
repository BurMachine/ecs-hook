package transport

import (
	"context"
	meteringadapterapi "ecs-hook/pkg/api/git.sbercloud.tech/cp/go/billing/metering-adapter/pkg/metering-adapter-api/v1"
	"errors"
	any1 "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

type RequestGRPC struct {
	state             string
	EventType         string
	EventRequestId    string
	EventMessage      interface{}
	EventTime         *timestamp.Timestamp
	UserId            string
	Source            string
	ProductInstanceID string
}

type BillingData struct {
	CustomerId          []uint8
	ResourceID          []uint8
	ResourceName        []uint8
	ResourceTag         []uint8
	ResourceTypeName    []uint8
	EnterpriseProjectId []uint8
	IamProjectId        []uint8
	UsageMeasureName    []uint8
	UsageTypeCode       []uint8
	OfferingName        []uint8
	UsageValue          float64
	Amount              float64
}

func Connect(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             time.Second,
			PermitWithoutStream: true,
		}),
	)
	return conn, err
}

func (data *RequestGRPC) SendResponse(conn *grpc.ClientConn) (*meteringadapterapi.SendMeteringEventResponse, error) {
	client := meteringadapterapi.NewMeteringAdapterServiceClient(conn)
	anyDat, ok := data.EventMessage.([]byte)
	if ok != true {
		return nil, errors.New("failed convert input interface to []byte")
	}

	req := meteringadapterapi.SendMeteringEventRequest{
		EventRequestId:    data.EventRequestId,
		EventMessage:      &any1.Any{Value: anyDat},
		EventType:         data.EventType,
		EventTime:         data.EventTime,
		ProductInstanceId: data.ProductInstanceID,
		UserId:            data.UserId,
		Source:            nil,
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("key", "value")) // можно добавить метаданные
	resp, err := client.SendMeteringEvent(ctx, &req)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}
	return resp, err
}
