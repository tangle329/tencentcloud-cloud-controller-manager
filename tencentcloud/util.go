package tencentcloud

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	nodeutil "k8s.io/kubernetes/pkg/util/node"
)

func (cloud *Cloud) getNodeHostIP(name string) (string, error) {
	node, err := cloud.kubeClient.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	ip, err := nodeutil.GetNodeHostIP(node)
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}
