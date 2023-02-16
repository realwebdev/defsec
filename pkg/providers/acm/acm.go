package acm

import "github.com/aquasecurity/defsec/pkg/types"

type acm struct {
	CertificateSummaryList []CertificateSummaryList
}

type CertificateSummaryList struct {
	Metadata types.Metadata
	CertificateArn    types.StringValue
	Name   types.StringValue
	Active types.BoolValue
}