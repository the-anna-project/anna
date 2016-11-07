package text

import (
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/xh3b4sd/anna/object/networkresponse"
	objectspec "github.com/xh3b4sd/anna/object/spec"
	"github.com/xh3b4sd/anna/object/textoutput"
	"github.com/xh3b4sd/anna/service"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

// ClientConfig represents the configuration used to create a new text
// interface object.
type ClientConfig struct {
	// Dependencies.

	ServiceCollection servicespec.Collection

	// Settings.

	// GRPCAddr is the host:port representation based on the golang convention
	// for net.Listen to serve gRPC traffic.
	GRPCAddr string
}

// DefaultClientConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultClientConfig() ClientConfig {
	newConfig := ClientConfig{
		// Dependencies.

		ServiceCollection: service.MustNewCollection(),

		// Settings.

		GRPCAddr: "127.0.0.1:9119",
	}

	return newConfig
}

// NewClient creates a new configured text interface object.
func NewClient(config ClientConfig) (systemspec.TextInterfaceClient, error) {
	newClient := &client{
		ClientConfig: config,
	}

	// Dependencies.

	if newClient.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	// Settings.

	if newClient.GRPCAddr == "" {
		return nil, maskAnyf(invalidConfigError, "gRPC address must not be empty")
	}

	return newClient, nil
}

type client struct {
	ClientConfig
}

func (c *client) DecodeResponse(streamTextResponse *StreamTextResponse) (objectspec.TextOutput, error) {
	if streamTextResponse.Code != networkresponse.CodeData {
		return nil, maskAnyf(invalidAPIResponseError, "API response code must be %d", networkresponse.CodeData)
	}

	textOutputConfig := textoutput.DefaultConfig()
	textOutputConfig.Output = streamTextResponse.Data.Output
	textResponse, err := textoutput.New(textOutputConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return textResponse, nil
}

func (c *client) EncodeRequest(textInput objectspec.TextInput) *StreamTextRequest {
	streamTextRequest := &StreamTextRequest{
		Echo:      textInput.GetEcho(),
		Input:     textInput.GetInput(),
		SessionID: textInput.GetSessionID(),
	}

	return streamTextRequest
}

func (c *client) StreamText(ctx context.Context) error {
	done := make(chan struct{}, 1)
	fail := make(chan error, 1)

	conn, err := grpc.Dial(c.GRPCAddr, grpc.WithInsecure())
	if err != nil {
		return maskAny(err)
	}
	defer conn.Close()

	client := NewTextInterfaceClient(conn)
	stream, err := client.StreamText(ctx)
	if err != nil {
		return maskAny(err)
	}

	// Listen on the outout of the text interface stream send it back to the
	// client.
	go func() {
		for {
			streamTextResponse, err := stream.Recv()
			if err == io.EOF {
				// The stream ended. We broadcast to all goroutines by closing the done
				// channel.
				close(done)
				return
			} else if err != nil {
				fail <- maskAny(err)
				return
			}

			textResponse, err := c.DecodeResponse(streamTextResponse)
			if err != nil {
				fail <- maskAny(err)
				return
			}
			c.Service().TextOutput().GetChannel() <- textResponse
		}
	}()

	// Listen on the client input channel and forward it to the server stream.
	go func() {
		for {
			select {
			case <-done:
				return
			case textInput := <-c.Service().TextInput().GetChannel():
				streamTextRequest := c.EncodeRequest(textInput)
				err := stream.Send(streamTextRequest)
				if err != nil {
					fail <- maskAny(err)
					return
				}
			}
		}
	}()

	for {
		select {
		case <-stream.Context().Done():
			close(done)
			return maskAny(stream.Context().Err())
		case <-done:
			return nil
		case err := <-fail:
			return maskAny(err)
		}
	}
}
