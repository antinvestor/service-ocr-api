package ocr_v1

import (
	"context"
	apic "github.com/antinvestor/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"

	"math"
)

const ctxKeyService = "ocrClientKey"

func defaultProfileClientOptions() []apic.ClientOption {
	return []apic.ClientOption{
		apic.WithEndpoint("ocr.api.antinvestor.com:443"),
		apic.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		apic.WithGRPCDialOption(grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func ToContext(ctx context.Context, ocrClient *OCRClient) context.Context {
	return context.WithValue(ctx, ctxKeyService, ocrClient)
}

func FromContext(ctx context.Context) *OCRClient {
	ocrClient, ok := ctx.Value(ctxKeyService).(*OCRClient)
	if !ok {
		return nil
	}

	return ocrClient
}

// OCRClient is a client for interacting with the profile service API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type OCRClient struct {
	// gRPC connection to the service.
	clientConn *grpc.ClientConn

	// The gRPC API client.
	ocrClient OCRServiceClient

	// The x-ant-* metadata to be sent with each request.
	xMetadata metadata.MD
}

// NewOCRClient creates a new ocr client.
// The service that an application uses to perform ocr requests
func NewOCRClient(ctx context.Context, opts ...apic.ClientOption) (*OCRClient, error) {
	clientOpts := defaultProfileClientOptions()

	connPool, err := apic.DialConnection(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	cl := &OCRClient{
		clientConn: connPool,
		ocrClient:  NewOCRServiceClient(connPool),
	}

	cl.setClientInfo()

	return cl, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (pc *OCRClient) Close() error {
	return pc.clientConn.Close()
}

// setClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (pc *OCRClient) setClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", apic.VersionGo()}, keyval...)
	kv = append(kv, "grpc", grpc.Version)
	pc.xMetadata = metadata.Pairs("x-ai-api-client", apic.XAntHeader(kv...))
}

func (pc *OCRClient) Recognize(ctx context.Context, id string, language string, properties map[string]string, fileId ...string) (*OcrResponse, error) {

	ctx2, cancel := context.WithTimeout(ctx, time.Second*300)
	defer cancel()

	ocrService := NewOCRServiceClient(pc.clientConn)

	ocrRequest := OcrRequest{
		ReferenceId: id,
		LanguageId:  language,
		FileId:      fileId,
		Properties:  properties,
	}

	return ocrService.Recognize(ctx2, &ocrRequest)
}

func (pc *OCRClient) StatusCheck(ctx context.Context, id string) (*OcrResponse, error) {

	ctx2, cancel := context.WithTimeout(ctx, time.Second*300)
	defer cancel()

	ocrService := NewOCRServiceClient(pc.clientConn)

	statusCheckRequest := StatusRequest{
		ReferenceId: id,
	}

	return ocrService.Status(ctx2, &statusCheckRequest)
}
