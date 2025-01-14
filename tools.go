//go:build tools
// +build tools

// Official workaround to track tool dependencies with go modules:
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	// Makefile
	"github.com/openshift/build-machinery-go/make/lib"
	"github.com/openshift/build-machinery-go/make/targets/openshift"
	"github.com/openshift/build-machinery-go/make/targets/openshift/operator"
)
