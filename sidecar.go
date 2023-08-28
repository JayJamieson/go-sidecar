package gosidecar

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

type LambdaClient struct {
	*lambda.Client
}

type LambdaFunction interface {
	Name() string
	Handler() string
	Package() string
}

// Execute will synchronously invoke and wait for response
func (c *LambdaClient) Execute(function LambdaFunction, payload any) ([]byte, error) {
	request, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	result, err := c.Invoke(nil, &lambda.InvokeInput{
		FunctionName: aws.String(function.Name()),
		LogType:      types.LogTypeNone,
		Payload:      request,
	})

	return result.Payload, err
}

// ExecuteAsync is fire and forget where returned result is not cared about
func (c *LambdaClient) ExecuteAsync(function LambdaFunction, payload any) error {
	// TODO: Implement async invocation
	panic("Not yet implemented")
	return nil
}

// ExecuteHTTP will invoke function using configured API Gateway or function url
func (c *LambdaClient) ExecuteHTTP(function LambdaFunction, payload any) error {
	// TODO: Implement execute HTTP invocation
	panic("Not yet implemented")
	return nil
}
