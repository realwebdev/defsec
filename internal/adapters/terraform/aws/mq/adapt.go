package mq

import (
	"github.com/aquasecurity/defsec/pkg/providers/aws/mq"
	"github.com/aquasecurity/defsec/pkg/terraform"
	"github.com/aquasecurity/defsec/pkg/types"
)

func Adapt(modules terraform.Modules) mq.MQ {
	return mq.MQ{
		Brokers: adaptBrokers(modules),
	}
}

func adaptBrokers(modules terraform.Modules) []mq.Broker {
	var brokers []mq.Broker
	for _, module := range modules {
		for _, resource := range module.GetResourcesByType("aws_mq_broker") {
			brokers = append(brokers, adaptBroker(resource))
		}
	}
	return brokers
}

func adaptBroker(resource *terraform.Block) mq.Broker {

	broker := mq.Broker{
		Metadata:                resource.GetMetadata(),
		PublicAccess:            types.BoolDefault(false, resource.GetMetadata()),
		EngineType:              resource.GetAttribute("engine_type").AsStringValueOrDefault("", resource),
		HostInstanceType:        resource.GetAttribute("host_instance_type").AsStringValueOrDefault("", resource),
		AutoMinorVersionUpgrade: resource.GetAttribute("auto_minor_version_upgrade").AsBoolValueOrDefault(true, resource),
		DeploymentMode:          resource.GetAttribute("deployment_mode").AsStringValueOrDefault("SINGLE_INSTANCE", resource),
		KmsKeyId:                types.StringDefault("", resource.GetMetadata()),
		Logging: mq.Logging{
			Metadata: resource.GetMetadata(),
			General:  types.BoolDefault(false, resource.GetMetadata()),
			Audit:    types.BoolDefault(false, resource.GetMetadata()),
		},
	}

	publicAccessAttr := resource.GetAttribute("publicly_accessible")
	broker.PublicAccess = publicAccessAttr.AsBoolValueOrDefault(false, resource)
	if logsBlock := resource.GetBlock("logs"); logsBlock.IsNotNil() {
		broker.Logging.Metadata = logsBlock.GetMetadata()
		auditAttr := logsBlock.GetAttribute("audit")
		broker.Logging.Audit = auditAttr.AsBoolValueOrDefault(false, logsBlock)
		generalAttr := logsBlock.GetAttribute("general")
		broker.Logging.General = generalAttr.AsBoolValueOrDefault(false, logsBlock)
	}
	if encryptBlock := resource.GetBlock("encryption_options"); encryptBlock.IsNotNil() {
		broker.KmsKeyId = encryptBlock.GetAttribute("kms_key_id").AsStringValueOrDefault("", resource)
	}

	return broker
}
