module github.com/open-cluster-management/multicloud-operators-subscription

require (
	github.com/aws/aws-sdk-go-v2 v0.18.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/ghodss/yaml v1.0.1-0.20180820084758-c7ce16629ff4
	github.com/go-openapi/spec v0.19.0
	github.com/google/go-cmp v0.3.1
	github.com/google/go-github/v28 v28.1.1
	github.com/onsi/gomega v1.5.0
	github.com/open-cluster-management/multicloud-operators-channel v0.0.0-alpha
	github.com/open-cluster-management/multicloud-operators-deployable v0.0.0-20200305002029-e8093169e4ff
	github.com/open-cluster-management/multicloud-operators-placementrule v0.0.0-20200304214605-e52bfdcb25bb
	github.com/open-cluster-management/multicloud-operators-subscription-release v0.0.0-20200304203457-d846afd9e135
	github.com/operator-framework/operator-sdk v0.12.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/common v0.4.1
	github.com/sabhiram/go-gitignore v0.0.0-20180611051255-d3107576ba94
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0
	gopkg.in/src-d/go-git.v4 v4.13.1
	k8s.io/api v0.0.0
	k8s.io/apiextensions-apiserver v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/cli-runtime v0.0.0
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/helm v2.15.1+incompatible
	k8s.io/klog v0.3.3
	k8s.io/kube-openapi v0.0.0-20190918143330-0270cf2f1c1d
	sigs.k8s.io/controller-runtime v0.3.0
	sigs.k8s.io/kustomize v2.0.3+incompatible
)

// Pinned to kubernetes-1.15.4
replace (
	k8s.io/api => k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190918200908-1e17798da8c1
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190918202139-0b14c719ca62
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190918203125-ae665f80358a
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20190918202959-c340507a5d48
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190612205613-18da4a14b22b
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190918200425-ed2f0867c778
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190817025403-3ae76f584e79
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20190918203248-97c07dcbb623
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190918201136-c3a845f1fbb2
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20190918202837-c54ce30c680e
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20190918202429-08c8357f8e2d
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20190918202713-c34a54b3ec8e
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20190918202550-958285cf3eef
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190918203421-225f0541b3ea
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190918202012-3c1ca76f5bda
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20190918201353-5cc279503896
)

replace github.com/Masterminds/semver => github.com/Masterminds/semver v1.4.2

go 1.13
