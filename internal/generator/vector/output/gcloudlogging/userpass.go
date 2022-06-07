package gcloudlogging

import (
	"github.com/openshift/cluster-logging-operator/internal/generator/vector/output/security"
)

type UserNamePass security.UserNamePass

func (up UserNamePass) Name() string {
	return "gcloudloggingUsernamePasswordTemplate"
}

func (up UserNamePass) Template() string {
	return `{{define "` + up.Name() + `" -}}
strategy = "basic"
user = "{{.Username}}"
password = "{{.Password}}"
{{- end}}`
}
