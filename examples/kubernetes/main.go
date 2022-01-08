package main

import (
	"github.com/jgrancell/logger"
	"github.com/jgrancell/logger/formats"
)

func main() {
	// The Standard Out Handler is the default, and has no required configuration
	k8s := logger.Logger{
		Format: &formats.KubernetesFormat{
			LogParameters: map[string]string{"pod": "example-pod-1", "node": "appnode-na-example"},
		},
	}

	// We can catch any errors the logger has and use the FallbackHandler to ensure they arent lost.
	if err := k8s.Init(); err != nil {
		// All logging commands give access to the convenience returns ReturnWithInt, ReturnWithError, and Return
		//   to provide int, error, and bool returns respectively.
		_ = k8s.Fallback.Handle(err.Error())
	}

	_ = k8s.Info("example kubernetes standard format message")

	// Batching log messages can be useful if your application (such as a k8s controller) may have several log lines
	//   related to the same ongoing end goal. This lets you cherrypick out log lines related only to a specific action
	//   happening in your application.
	//
	// Batches are not committed to the log until you either call `.Parse()` or one of the accessor functions, such as `.Warning()` or `.Info()`
	batch := k8s.Batch("aaaa-bbbb-cccc-dddd")

	_ = batch.WithMessage(logger.LOG_DEBUG, "detected example CRD change for namespace-foo/name-bar. starting update process for k8s resources.")
	batch.WithMessage(logger.LOG_WARNING, "missing example value foobar in updated CDR. using value for foobar from previous version")
	_ = k8s.Batch("111-222-333-444").Warning("missing example value fizzbuzz in updated CDR. using value for foobar from previous version")
	_ = batch.Info("completed reconciliation for namespace-foo/name-bar resource. resource is up to date.")
}
