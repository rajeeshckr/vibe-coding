package dsl

import (
	"gopkg.in/yaml.v3"
)

// --- Data Structures (similar to original main.go) ---

// AMITag represents a single AMI tag.
type AMITag struct {
	Revision    string `yaml:"Revision,omitempty"`
	Environment string `yaml:"Environment,omitempty"`
	Hostgroup   string `yaml:"Hostgroup,omitempty"`
	Platform    string `yaml:"Platform,omitempty"`
	BaseName    string `yaml:"BaseName,omitempty"`
}

// Tags mirrors the 'tags' section of the YAML.
type Tags struct {
	AmiRevisionMain   *AMITag `yaml:"ami_revision_main"`
	AmiRevisionBranch *AMITag `yaml:"ami_revision_branch"`
	Envs              struct {
		ProductionAmiEnv *AMITag `yaml:"production_ami_env"`
		StagingAmiEnv    *AMITag `yaml:"staging_ami_env"`
	} `yaml:"envs"`
	Providers struct {
		Chef *AMITag `yaml:"chef"`
		Eks  *AMITag `yaml:"eks"`
	} `yaml:"providers"`
	Platforms struct {
		Ubuntu2204 struct {
			Arm64 *AMITag `yaml:"arm64"`
			Amd64 *AMITag `yaml:"amd64"`
		} `yaml:"ubuntu_22_04"`
	} `yaml:"platforms"`
	BaseNames struct {
		EksArm64 *AMITag `yaml:"eks_arm64"`
		EksAmd64 *AMITag `yaml:"eks_amd64"`
		SfnAmd64 *AMITag `yaml:"sfn_amd64"`
		SfnArm64 *AMITag `yaml:"sfn_arm64"`
	} `yaml:"base_names"`
}

// Vars mirrors the 'vars' section of the YAML.
type Vars struct {
	SandboxAmd64Ami    *AMITag `yaml:"sandbox_amd64_ami"`
	SandboxArm64Ami    *AMITag `yaml:"sandbox_arm64_ami"`
	SandboxEksAmd64Ami *AMITag `yaml:"sandbox_eks_amd64_ami"`
}

// Cluster represents a single cluster definition.
type Cluster struct {
	Name      string `yaml:"name"`
	Partition string `yaml:"partition,omitempty"`
	// Other cluster properties will go here
}

// ClusterMetadata is the top-level struct.
type ClusterMetadata struct {
	Tags     *Tags      `yaml:"tags"`
	Vars     *Vars      `yaml:"vars"`
	Clusters []*Cluster `yaml:"clusters"`
}

// --- DSL Engine ---

var (
	// Definition holds the top-level definition as it's being built.
	Definition = &ClusterMetadata{
		Tags: &Tags{},
		Vars: &Vars{},
	}
	// currentTag holds the context for the current AMITag being defined.
	currentTag *AMITag
	// currentCluster holds the context for the current Cluster being defined.
	currentCluster *Cluster
)

// Run executes the DSL definition and returns the resulting object.
func Run(dsl func()) *ClusterMetadata {
	dsl()
	return Definition
}

// --- DSL Functions ---

// TagsDef defines the tags section.
func TagsDef(dsl func()) {
	dsl()
}

// AmiRevisionMain defines the main AMI revision.
func AmiRevisionMain(revision string) {
	Definition.Tags.AmiRevisionMain = &AMITag{Revision: revision}
}

// AmiRevisionBranch defines the branch AMI revision.
func AmiRevisionBranch(revision string) {
	Definition.Tags.AmiRevisionBranch = &AMITag{Revision: revision}
}

// Envs defines the environments.
func Envs(dsl func()) {
	dsl()
}

// ProductionAmiEnv defines the production AMI environment.
func ProductionAmiEnv(dsl func()) {
	Definition.Tags.Envs.ProductionAmiEnv = &AMITag{}
	inTag(Definition.Tags.Envs.ProductionAmiEnv, dsl)
}

// StagingAmiEnv defines the staging AMI environment.
func StagingAmiEnv(dsl func()) {
	Definition.Tags.Envs.StagingAmiEnv = &AMITag{}
	inTag(Definition.Tags.Envs.StagingAmiEnv, dsl)
}

// Providers defines the providers.
func Providers(dsl func()) {
	dsl()
}

// Chef defines the Chef provider.
func Chef(dsl func()) {
	Definition.Tags.Providers.Chef = &AMITag{}
	inTag(Definition.Tags.Providers.Chef, dsl)
}

// Eks defines the EKS provider.
func Eks(dsl func()) {
	Definition.Tags.Providers.Eks = &AMITag{}
	inTag(Definition.Tags.Providers.Eks, dsl)
}

// Platforms defines the platforms.
func Platforms(dsl func()) {
	dsl()
}

// Ubuntu2204 defines the Ubuntu 22.04 platform.
func Ubuntu2204(dsl func()) {
	dsl()
}

// Arm64 defines the arm64 architecture for the current platform.
func Arm64(dsl func()) {
	Definition.Tags.Platforms.Ubuntu2204.Arm64 = &AMITag{}
	inTag(Definition.Tags.Platforms.Ubuntu2204.Arm64, dsl)
}

// Amd64 defines the amd64 architecture for the current platform.
func Amd64(dsl func()) {
	Definition.Tags.Platforms.Ubuntu2204.Amd64 = &AMITag{}
	inTag(Definition.Tags.Platforms.Ubuntu2204.Amd64, dsl)
}

// BaseNames defines the base names.
func BaseNames(dsl func()) {
	dsl()
}

// EksArm64 defines the EKS arm64 base name.
func EksArm64(dsl func()) {
	Definition.Tags.BaseNames.EksArm64 = &AMITag{}
	inTag(Definition.Tags.BaseNames.EksArm64, dsl)
}

// EksAmd64 defines the EKS amd64 base name.
func EksAmd64(dsl func()) {
	Definition.Tags.BaseNames.EksAmd64 = &AMITag{}
	inTag(Definition.Tags.BaseNames.EksAmd64, dsl)
}

// SfnAmd64 defines the SFN amd64 base name.
func SfnAmd64(dsl func()) {
	Definition.Tags.BaseNames.SfnAmd64 = &AMITag{}
	inTag(Definition.Tags.BaseNames.SfnAmd64, dsl)
}

// SfnArm64 defines the SFN arm64 base name.
func SfnArm64(dsl func()) {
	Definition.Tags.BaseNames.SfnArm64 = &AMITag{}
	inTag(Definition.Tags.BaseNames.SfnArm64, dsl)
}

// VarsDef defines the vars section.
func VarsDef(dsl func()) {
	dsl()
}

// SandboxAmd64Ami defines the sandbox amd64 AMI.
func SandboxAmd64Ami(dsl func()) {
	Definition.Vars.SandboxAmd64Ami = &AMITag{}
	inTag(Definition.Vars.SandboxAmd64Ami, dsl)
}

// SandboxArm64Ami defines the sandbox arm64 AMI.
func SandboxArm64Ami(dsl func()) {
	Definition.Vars.SandboxArm64Ami = &AMITag{}
	inTag(Definition.Vars.SandboxArm64Ami, dsl)
}

// SandboxEksAmd64Ami defines the sandbox EKS amd64 AMI.
func SandboxEksAmd64Ami(dsl func()) {
	Definition.Vars.SandboxEksAmd64Ami = &AMITag{}
	inTag(Definition.Vars.SandboxEksAmd64Ami, dsl)
}

// ClusterDef defines a new cluster.
func ClusterDef(dsl func()) {
	cluster := &Cluster{}
	Definition.Clusters = append(Definition.Clusters, cluster)
	inCluster(cluster, dsl)
}

// --- Attribute setters ---

// Name sets the name for the current cluster.
func Name(val string) {
	if currentCluster != nil {
		currentCluster.Name = val
	}
}

// Partition sets the partition for the current cluster.
func Partition(val string) {
	if currentCluster != nil {
		currentCluster.Partition = val
	}
}

// Environment sets the Environment field on the current AMI tag.
func Environment(val string) {
	if currentTag != nil {
		currentTag.Environment = val
	}
}

// Hostgroup sets the Hostgroup field on the current AMI tag.
func Hostgroup(val string) {
	if currentTag != nil {
		currentTag.Hostgroup = val
	}
}

// Platform sets the Platform field on the current AMI tag.
func Platform(val string) {
	if currentTag != nil {
		currentTag.Platform = val
	}
}

// BaseName sets the BaseName field on the current AMI tag.
func BaseName(val string) {
	if currentTag != nil {
		currentTag.BaseName = val
	}
}

// Merge copies fields from a source tag to the current tag.
func Merge(source *AMITag) {
	if currentTag != nil && source != nil {
		if source.Revision != "" {
			currentTag.Revision = source.Revision
		}
		if source.Environment != "" {
			currentTag.Environment = source.Environment
		}
		// Add other fields as needed
	}
}

// ToYAML marshals the metadata to a YAML string.
func ToYAML(data *ClusterMetadata) (string, error) {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(yamlData), nil
}

// inTag is a helper to set the context for attribute setters.
func inTag(tag *AMITag, dsl func()) {
	original := currentTag
	currentTag = tag
	dsl()
	currentTag = original
}

// inCluster is a helper to set the context for cluster attribute setters.
func inCluster(cluster *Cluster, dsl func()) {
	original := currentCluster
	currentCluster = cluster
	dsl()
	currentCluster = original
}
