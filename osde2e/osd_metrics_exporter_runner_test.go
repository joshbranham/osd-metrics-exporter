// THIS FILE IS GENERATED BY BOILERPLATE. DO NOT EDIT.
//go:build osde2e
// +build osde2e

package osde2etests

import (
	"path/filepath"
	"testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	testResultsDirectory = "/test-run-results"
	jUnitOutputFilename  = "junit-osd-metrics-exporter.xml"
)

// Test entrypoint. osde2e runs this as a test suite on test pod.
func TestOsdMetricsExporter(t *testing.T) {
	RegisterFailHandler(Fail)

	suiteConfig, reporterConfig := GinkgoConfiguration()
	reporterConfig.JUnitReport = filepath.Join(testResultsDirectory, jUnitOutputFilename)
	RunSpecs(t, "Osd Metrics Exporter", suiteConfig, reporterConfig)

}

