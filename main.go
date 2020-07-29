package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func main() {
	var x typev1.EndpointsInterface
	x.List(context.Background(), metav1.ListOptions{})
}
