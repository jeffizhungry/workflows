package workflows

import (
	"context"
	"fmt"
	"time"

	"github.com/jeffizhungry/workflows/exampleapp/worker/workflows"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// TaskListName is the task list for this sample
const (
	TaskListName = "helloWorldGroup"
	signalName   = "helloWorldSignal"
)

// This is registration process where you register all your workflows
// and activity function handlers.
func init() {
	workflow.Register(HelloWorkflow)
	activity.Register(helloActivity)
}

var activityOptions = workflow.ActivityOptions{
	ScheduleToStartTimeout: 1 * time.Minute,
	StartToCloseTimeout:    1 * time.Minute,
	HeartbeatTimeout:       20 * time.Second,
}

func helloActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("helloworld activity started")
	return "Hello " + name + "! How old are you!", nil
}

// HelloWorkflow is a hello workflow
func HelloWorkflow(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	logger := workflow.GetLogger(ctx)
	logger.Info("helloworld workflow started")
	var activityResult string
	err := workflow.ExecuteActivity(ctx, helloActivity, name).Get(ctx, &activityResult)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return "", err
	}

	// After saying hello, the workflow will wait for you to inform it of your age!
	selector := workflow.NewSelector(ctx)
	var ageResult int
	for {
		signalChan := workflow.GetSignalChannel(ctx, signalName)
		selector.AddReceive(signalChan, func(c workflow.Channel, more bool) {
			c.Receive(ctx, &ageResult)
			workflow.GetLogger(ctx).Info("Received age results from signal!", zap.String("signal", signalName), zap.Int("value", ageResult))
		})
		workflow.GetLogger(ctx).Info("Waiting for signal on channel.. " + signalName)
		// Wait for signal
		selector.Select(ctx)

		// We can check the age and return an appropriate response
		if ageResult > 0 && ageResult < 150 {
			logger.Info("Workflow completed.", zap.String("NameResult", activityResult), zap.Int("AgeResult", ageResult))

			return fmt.Sprintf("Hello "+name+"! You are %v years old!", ageResult), nil
		} else {
			return "You can't be that old!", nil
		}
	}
}

// HelloWorkflowClient is used to interact with the HelloWorkflow
type HelloWorkflowClient struct {
	cadenceClient client.Client
}

// NewHelloWorkflowClient initializes a new HelloWorkflowClient
func NewHelloWorkflowClient(cadenceClient client.Client) *HelloWorkflowClient {
	return &HelloWorkflowClient{cadenceClient}
}

// Start starts HelloWorkflow
func (wc *HelloWorkflowClient) Start(ctx context.Context, accountID string) (workflowID string, err error) {
	options := client.StartWorkflowOptions{
		TaskList:                     workflows.TaskListName,
		ExecutionStartToCloseTimeout: time.Hour * 24,
	}
	execution, err := wc.cadenceClient.StartWorkflow(ctx, options, HelloWorkflow, accountID)
	if err != nil {
		return "", err
	}
	return execution.ID, nil
}

// Continue continues HelloWorkflow with an age param.
func (wc *HelloWorkflowClient) Continue(ctx context.Context, workflowID string, age int) error {
	return wc.cadenceClient.SignalWorkflow(
		ctx,
		workflowID,
		"",
		signalName,
		age,
	)
}
