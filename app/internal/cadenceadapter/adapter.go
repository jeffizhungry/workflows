package cadenceadapter

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/encoded"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CadenceConfig contains cadence config information
type CadenceConfig struct {
	Domain   string
	Service  string
	HostPort string
}

// CadenceAdapter class for workflow helper.
type CadenceAdapter struct {
	CadenceClient  client.Client
	ServiceClient  workflowserviceclient.Interface
	Scope          tally.Scope
	Logger         *zap.Logger
	Config         CadenceConfig
	Builder        *WorkflowClientBuilder
	DataConverter  encoded.DataConverter
	CtxPropagators []workflow.ContextPropagator
}

// Setup setup the config for the code run
func (h *CadenceAdapter) Setup(config CadenceConfig) {
	// Return early if already initialized
	if h.CadenceClient != nil {
		return
	}

	logger, _ := zap.NewDevelopment()
	h.Logger = logger
	h.Config = config
	hostPort := h.Config.HostPort
	domainName := h.Config.Domain
	ctxLog := logrus.WithFields(logrus.Fields{
		"domain":   domainName,
		"hostPort": hostPort,
	})

	// Create builder
	h.Builder = NewBuilder(logger, hostPort, domainName)

	// Cadence client used to perform CRUD operation.
	cadenceClient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to build cadence client")
	}
	h.CadenceClient = cadenceClient

	// Service client that communicates with cadence using the rpc.
	serviceClient, err := h.Builder.BuildServiceClient()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to build service client")
	}
	h.ServiceClient = serviceClient

	// What is this?
	domainClient, _ := h.Builder.BuildCadenceDomainClient()
	_, err = domainClient.Describe(context.Background(), domainName)
	if err != nil {
		ctxLog.WithError(err).Error("Domain doesn't exist")
		return
	}
	ctxLog.Info("Domain successfully registered")
}
