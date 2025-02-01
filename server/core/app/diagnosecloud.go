package app

import "github.com/fredyk/diagnose-cloud/server/core/ports"

type DiagnoseCloud struct {
	ExternalDiagnoseService ports.DiagnosePort
}

type Options struct {
	ExternalDiagnoseService ports.DiagnosePort
}

func NewDiagnoseCloud(options Options) *DiagnoseCloud {
	return &DiagnoseCloud{
		ExternalDiagnoseService: options.ExternalDiagnoseService,
	}
}
