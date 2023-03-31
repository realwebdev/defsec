package network

var terraformNoCommonPrivateElasticLoadBalancerGoodExamples = []string{
	`
 resource "nifcloud_elb" "good_example" {
   elb_name          = "foobar"
   availability_zone = "east-11"
   instance_port     = 80
   protocol          = "HTTP"
   lb_port           = 80

   network_interface {
     network_id = nifcloud_private_lan.main.id
   }
 }
 `,
}

var terraformNoCommonPrivateElasticLoadBalancerBadExamples = []string{
	`
 resource "nifcloud_elb" "bad_example" {
   elb_name          = "foobar"
   availability_zone = "east-11"
   instance_port     = 80
   protocol          = "HTTP"
   lb_port           = 80

   network_interface {
     network_id = "net-COMMON_PRIVATE"
   }
 }
 `,
}

var terraformNoCommonPrivateElasticLoadBalancerLinks = []string{
	`https://registry.terraform.io/providers/nifcloud/nifcloud/latest/docs/resources/elb#network_id`,
}

var terraformNoCommonPrivateElasticLoadBalancerRemediationMarkdown = ``
