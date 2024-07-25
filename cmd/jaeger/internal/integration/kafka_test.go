// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"testing"

	"github.com/jaegertracing/jaeger/plugin/storage/integration"
)

func TestKafkaStorage(t *testing.T) {
	integration.SkipUnlessEnv(t, "kafka")

	collectorConfig := "../../collector-with-kafka.yaml"
	ingesterConfig := "../../ingester-remote-storage.yaml"

	collector := &E2EStorageIntegration{
		SkipStorageCleaner: true,
		ConfigFile:         collectorConfig,
	}

	// Initialize and start the collector
	collector.e2eInitialize(t, "kafka")

	ingester := &E2EStorageIntegration{
		ConfigFile: ingesterConfig,
		StorageIntegration: integration.StorageIntegration{
			CleanUp:                      purge,
			GetDependenciesReturnsSource: true,
			SkipArchiveTest:              true,
		},
	}

	// Initialize and start the ingester
	ingester.e2eInitialize(t, "kafka")

	// Run the span store tests
	ingester.RunSpanStoreTests(t)
}