package meta

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/logicmonitor/terraform-provider-helm/pkg/helm"
	"github.com/logicmonitor/terraform-provider-helm/pkg/helm/repo"
	"github.com/logicmonitor/terraform-provider-helm/pkg/kubernetes"
)

// Meta is the meta struct used by Terraform.
type Meta struct {
	helm.Helm
	kubernetes.Kubernetes
	repo.Repo
	ExplicitPath string
	Data         *schema.ResourceData
}
