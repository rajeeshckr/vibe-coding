package main

import (
	"fmt"
	"log"

	. "github.com/rajeeshckr/vibe-coding/cluster-metadata/internal/dsl"
)

var Design = Run(func() {
	ClusterDef(func() {
		Name("sandbox")
		Partition("sandbox")
		Env("staging")
		Region("us-west-2")

		Amd64AmiTags(func() {
			Environment("production")
			Revision("1f5915df9330551ceb673d198d7ec2c458d6e3a8")
			BaseName("ubuntu22.04_k8s_base_singlemount_cgroupv2")
		})

		Arm64AmiTags(func() {
			Environment("production")
			Revision("1f5915df9330551ceb673d198d7ec2c458d6e3a8")
			BaseName("ubuntu22.04_k8s_base_singlemount_arm_cgroupv2")
		})

		CookbookURL("<%= ENV['COOKBOOK_URL'].to_json %>")
		NodeRolloutBatchSize(`<%= ENV["NODE_ROLLOUT_BATCH_SIZE"] || "50%" %>`)
		SlackChannels("compute-cluster-deploys-staging")
		AdditionalSshTeams("engineering")

		NodegroupsDef(func() {
			EtcdMember(func() {
				Autoscale(false)
				Type("etcd")
				RootEbsSize("100")
			})
			EtcdEventsMember(func() {
				Autoscale(false)
				Type("etcd-events")
				RootEbsSize("100")
			})
			Api(func() {
				Autoscale(false)
				DrainTimeoutNodegroup("15m")
				RootEbsSize("100")
				InstanceTypes("m7g.2xlarge", "m6g.2xlarge")
			})
			Node(func() {
				MaxSize(10)
				ScaleDownGracePeriod("30s")
				DrainTimeoutNodegroup("1h")
				InstanceTypes("m7i-flex.2xlarge", "t3.2xlarge", "m7i.2xlarge", "m6i.2xlarge")
			})
		})

		DefaultAttributesDef(func() {
			ClusterAttribute("pod998") // Mismatch in source yaml, using pod998 from default_attributes
		})
	})
})

func main() {
	// Marshal the Go struct to YAML
	yamlData, err := ToYAML(Design)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(string(yamlData))
}
