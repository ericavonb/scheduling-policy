package main

import (
	"context"
	"runtime"

	stub "github.com/ericavonb/scheduling-policy/pkg/stub"
	sdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	k8sutil "github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"

	"github.com/sirupsen/logrus"
	//_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()

	sdk.ExposeMetricsPort()

	resource := "policy.experiments.k8s.io/v1alpha1"
	kind := "SchedulingPolicy"
	resyncPeriod := 5
	logrus.Infof("Watching %s, %s, %d", resource, kind, resyncPeriod)
	sdk.Watch(resource, kind, "", resyncPeriod)
	sdk.Handle(stub.NewHandler())
	sdk.Run(context.TODO())
}
