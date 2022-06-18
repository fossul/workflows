module github.com/fossul/workflows

go 1.16

require (
	github.com/fossul/fossul v0.5.0
	github.com/golang/mock v1.6.0
	github.com/stretchr/testify v1.7.1
	go.temporal.io/api v1.8.0
	go.temporal.io/sdk v1.15.0
	go.temporal.io/sdk/contrib/tools/workflowcheck v0.0.0-20220616213014-706af9738b9c
)

replace (
	github.com/fossul/fossul/src/client/k8s => ./src/client/k8s
	github.com/fossul/fossul/src/client/k8s/snapshotctrl/client/versioned => ../fossul/src/client/k8s/snapshotctrl/client/versioned
	github.com/fossul/fossul/src/client/k8s/virtctrl/client/versioned => ../fossul/src/client/k8s/virtctrl/client/versioned
	k8s.io/api => k8s.io/api v0.22.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.22.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.4
	k8s.io/apiserver => k8s.io/apiserver v0.22.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.22.4
	k8s.io/client-go => k8s.io/client-go v0.22.4
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.22.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.22.4
	k8s.io/code-generator => k8s.io/code-generator v0.22.4
	k8s.io/component-base => k8s.io/component-base v0.22.4
	k8s.io/cri-api => k8s.io/cri-api v0.22.4
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.22.4
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.22.4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.22.4
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20201113171705-d219536bb9fd
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.22.4
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.22.4
	k8s.io/kubectl => k8s.io/kubectl v0.22.4
	k8s.io/kubelet => k8s.io/kubelet v0.22.4
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.22.4
	k8s.io/metrics => k8s.io/metrics v0.22.4
	k8s.io/node-api => k8s.io/node-api v0.22.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.22.4
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.22.4
	k8s.io/sample-controller => k8s.io/sample-controller v0.22.4
)
