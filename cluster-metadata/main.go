package main

import (
	"fmt"
	"log"

	. "github.com/rajeeshckr/vibe-coding/cluster-metadata/dsl"
)

var Design = Run(func() {
	Tags(func() {
		AmiRevisionMain("1f5915df9330551ceb673d198d7ec2c458d6e3a8")
		AmiRevisionBranch("1990f8966675bd0bdf9172e916eca2739303a4f1")

		Envs(func() {
			ProductionAmiEnv(func() {
				Environment("production")
			})
			StagingAmiEnv(func() {
				Environment("staging")
			})
		})

		Providers(func() {
			Chef(func() {
				Hostgroup("k8s-singlemount")
			})
			Eks(func() {
				Hostgroup("k8s-eks")
			})
		})

		Platforms(func() {
			Ubuntu2204(func() {
				Arm64(func() {
					Platform("ubuntu-graviton-22.04")
				})
				Amd64(func() {
					Platform("ubuntu-22.04")
				})
			})
		})

		BaseNames(func() {
			EksArm64(func() {
				BaseName("ubuntu22.04_eks_base_arm")
			})
			EksAmd64(func() {
				BaseName("ubuntu22.04_eks_base")
			})
			SfnAmd64(func() {
				BaseName("ubuntu22.04_k8s_base_singlemount")
			})
			SfnArm64(func() {
				BaseName("ubuntu22.04_k8s_base_singlemount_arm")
			})
		})
	})

	Vars(func() {
		// This shows how to reuse definitions, similar to YAML anchors.
		// We get the definitions from the already-processed `Tags` section
		// by accessing the exported `Definition` object from the dsl package.
		productionAmiEnv := Definition.Tags.Envs.ProductionAmiEnv
		amiRevisionMain := Definition.Tags.AmiRevisionMain

		SandboxAmd64Ami(func() {
			Merge(productionAmiEnv)
			Merge(amiRevisionMain)
			BaseName("ubuntu22.04_k8s_base_singlemount_cgroupv2")
		})

		SandboxArm64Ami(func() {
			Merge(productionAmiEnv)
			Merge(amiRevisionMain)
			BaseName("ubuntu22.04_k8s_base_singlemount_arm_cgroupv2")
		})

		SandboxEksAmd64Ami(func() {
			Merge(productionAmiEnv)
			Merge(amiRevisionMain)
			BaseName("ubuntu22.04_eks_base_cgroupv2")
		})
	})

	Cluster("pod998", func() {
		Partition("pod998")
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
