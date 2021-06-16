// Code generated by "register.go" DO NOT EDIT.
package main

import (
	_ "github.com/flant/addon-operator/sdk"

	_ "github.com/deckhouse/deckhouse/global-hooks"
	_ "github.com/deckhouse/deckhouse/global-hooks/discovery"
	_ "github.com/deckhouse/deckhouse/modules/000-common/hooks"
	_ "github.com/deckhouse/deckhouse/modules/020-deckhouse/hooks"
	_ "github.com/deckhouse/deckhouse/modules/030-cloud-provider-aws/hooks"
	_ "github.com/deckhouse/deckhouse/modules/030-cloud-provider-azure/hooks"
	_ "github.com/deckhouse/deckhouse/modules/030-cloud-provider-gcp/hooks"
	_ "github.com/deckhouse/deckhouse/modules/030-cloud-provider-openstack/hooks"
	_ "github.com/deckhouse/deckhouse/modules/030-cloud-provider-vsphere/hooks"
	_ "github.com/deckhouse/deckhouse/modules/030-cloud-provider-yandex/hooks"
	_ "github.com/deckhouse/deckhouse/modules/035-cni-flannel/hooks"
	_ "github.com/deckhouse/deckhouse/modules/040-control-plane-manager/hooks"
	_ "github.com/deckhouse/deckhouse/modules/040-node-manager/hooks"
	_ "github.com/deckhouse/deckhouse/modules/041-kube-proxy/hooks"
	_ "github.com/deckhouse/deckhouse/modules/150-user-authn/hooks"
	_ "github.com/deckhouse/deckhouse/modules/300-prometheus/hooks"
	_ "github.com/deckhouse/deckhouse/modules/302-vertical-pod-autoscaler/hooks"
	_ "github.com/deckhouse/deckhouse/modules/340-monitoring-applications/hooks"
	_ "github.com/deckhouse/deckhouse/modules/340-monitoring-ping/hooks"
	_ "github.com/deckhouse/deckhouse/modules/340-prometheus-madison-integration/hooks"
	_ "github.com/deckhouse/deckhouse/modules/600-secret-copier/hooks"
	_ "github.com/deckhouse/deckhouse/modules/998-flant-pricing/hooks"
)
