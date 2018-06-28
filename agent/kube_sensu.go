package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sensu/sensu-go/transport"
	"github.com/sensu/sensu-go/types"

	"github.com/bitnami-labs/kubewatch/config"
	"github.com/bitnami-labs/kubewatch/pkg/controller"
	"github.com/bitnami-labs/kubewatch/pkg/utils"
	"github.com/sirupsen/logrus"
)

type KubeSensu struct {
	agent *Agent
}

type KubeWatchHandler struct {
	agent *Agent
}

func NewKubeSensu(a *Agent) *KubeSensu {
	return &KubeSensu{
		agent: a,
	}
}

func (ks *KubeSensu) Run(ctx context.Context) error {
	// ks.RunKubeStateMetrics(ctx)
	ks.RunKubeWatch(ctx)

	return nil
}

func (ks *KubeSensu) RunKubeWatch(ctx context.Context) error {
	config := &config.Config{}

	// TODO(schu): configure resources to watch
	config.Resource.Services = true

	handler := &KubeWatchHandler{
		agent: ks.agent,
	}

	controller.Start(config, handler)

	return nil
}

func (ks *KubeSensu) RunKubeStateMetrics(ctx context.Context) error {
	go func() {
		// TODO(schu): should be configurable
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-ctx.Done():
				logger.Warn("kube-sensu stopping")
				return
			case <-ticker.C:
				if err := ks.SendKubeStateMetrics(ctx); err != nil {
					logger.WithError(err).Error("failed to send kube state metrics")
				}
			}
		}
	}()
	return nil
}

func (ks *KubeSensu) SendKubeStateMetrics(ctx context.Context) error {

	// TODO: figure out which metrics make sense for a general
	// Sensu use-case and what should be configurable for users.
	// Also, metrics need to define a handler; what should handle
	// kube state metrics?

	return fmt.Errorf("no implemented yet")
}

func (handler *KubeWatchHandler) Init(config *config.Config) error {
	return nil
}

func (handler *KubeWatchHandler) ObjectCreated(obj interface{}) {
	logger.WithField("obj", obj).Info("object created")

	objectMeta := utils.GetObjectMetaData(obj)

	namespace := objectMeta.Namespace
	name := objectMeta.Name

	keepalive := &types.Event{}

	keepalive.Entity = &types.Entity{
		ID:               fmt.Sprintf("%s-%s", namespace, name),
		Class:            types.EntityProxyClass,
		Environment:      "default",
		Organization:     "default",
		KeepaliveTimeout: 30,
	}

	keepalive.Timestamp = time.Now().Unix()

	msg, err := json.Marshal(keepalive)
	if err != nil {
		logger.WithError(err).Info("failed to marshal keepalive message")
		return
	}

	handler.agent.sendMessage(transport.MessageTypeKeepalive, msg)
}

func (handler *KubeWatchHandler) ObjectDeleted(obj interface{}) {
	logger.WithField("obj", obj).Info("object deleted")
}

func (handler *KubeWatchHandler) ObjectUpdated(oldObj, newObj interface{}) {
	logger.WithFields(logrus.Fields{
		"oldObj": oldObj,
		"newObj": newObj,
	}).Info("object updated")
}
