package helm

import (
	keptnevents "github.com/keptn/go-utils/pkg/events"
	corev1 "k8s.io/api/core/v1"
)

type CanaryOnNamespaceGenerator struct {
}

func NewCanaryOnNamespaceGenerator() *CanaryOnNamespaceGenerator {
	return &CanaryOnNamespaceGenerator{}
}

func (c *CanaryOnNamespaceGenerator) GetCanaryService(originalSvc corev1.Service, event *keptnevents.ServiceCreateEventData, stageName string) (canaryService *corev1.Service) {

	canaryService = &corev1.Service{}

	canaryService.Kind = "Service"
	canaryService.APIVersion = "v1"
	canaryService.Name = originalSvc.Name + "-canary"
	canaryService.Spec.Type = "ExternalName"
	canaryService.Spec.ExternalName = originalSvc.Name + "." + c.GetNamespace(event.Project, stageName, false) + ".svc.cluster.local"
	canaryService.Spec.Ports = originalSvc.Spec.Ports
	return
}

func (*CanaryOnNamespaceGenerator) IsK8sResourceDuplicated() bool {

	return true
}

// GetNamespace returns the namespace for a specific project and stage
func (*CanaryOnNamespaceGenerator) GetNamespace(project string, stage string, generated bool) string {
	suffix := ""
	if generated {
		suffix = "-generated"
	}
	return project + "-" + stage + suffix
}
