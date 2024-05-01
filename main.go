package lambda_base

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	// Initialize a new Zap logger in production mode
	l, _ := zap.NewProduction()
	logger = l
	defer logger.Sync() // Flushes logger buffer, if any, when this function exits
}

// Event represents the structure of the event data expected by MyHandler
type Event struct {
	// Define the structure of the event data here if needed
}

// MyHandler is the AWS Lambda handler function
func MyHandler(ctx context.Context, e Event) error {
	// Log an informational message using Zap logger
	logger.Info("in my handler", zap.Any("event", e))
	return nil
}

func main() {
	// Start the Lambda execution using the lambda.Start() function
	lambda.Start(MyHandler)
}
