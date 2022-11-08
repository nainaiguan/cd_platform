package sit

import (
	"cd_platform/util"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (s *Service) CreateDeployment(ctx context.Context, project string, raw []byte) error {
	uns, err := util.RawJsonToUnstructured(raw)

	if err != nil {
		util.Logger.Errorf("exec.CreateDeployment err: %s", err)
		return err
	}

	if _, err := s.Mid.K8sclient.DynamicClient.Resource(schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}).Namespace(util.ProjectToSitNS(project)).Create(context.TODO(), uns, metav1.CreateOptions{}); err != nil {
		util.Logger.Errorf("exec.CreateDeployment err: %s", err)
		return err
	}
	
	return nil
}