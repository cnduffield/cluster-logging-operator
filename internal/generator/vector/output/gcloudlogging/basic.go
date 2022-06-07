package gcloudlogging

import (
	"github.com/openshift/cluster-logging-operator/internal/generator"
)

type BasicAuthConf generator.ConfLiteral

func (t BasicAuthConf) Name() string {
	return "gcloudloggingBasicAuthConf"
}

func (t BasicAuthConf) Template() string {
	return `
{{define "gcloudloggingBasicAuthConf" -}}
# {{.Desc}}
[sinks.{{.ComponentID}}.auth]
{{- end}}`
}
