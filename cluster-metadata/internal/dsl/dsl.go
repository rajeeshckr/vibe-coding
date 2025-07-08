package dsl

import (
	"gopkg.in/yaml.v3"
)

// --- Data Structures ---

// AMITag represents a single AMI tag.
type AMITag struct {
	Revision    string `yaml:"Revision,omitempty"`
	Environment string `yaml:"Environment,omitempty"`
	BaseName    string `yaml:"BaseName,omitempty"`
}

// Nodegroup represents a nodegroup configuration.
type Nodegroup struct {
	Autoscale      *bool    `yaml:"autoscale,omitempty"`
	Type           string   `yaml:"type,omitempty"`
	RootEbsSize    string   `yaml:"root_ebs_size,omitempty"`
	MaxSize        *int     `yaml:"max_size,omitempty"`
	ScaleDownGrace string   `yaml:"scale_down_grace_period,omitempty"`
	DrainTimeout   string   `yaml:"drain_timeout,omitempty"`
	InstanceTypes  []string `yaml:"amd64_instance_types,omitempty"`
}

// Nodegroups represents the set of nodegroups in a cluster.
type Nodegroups struct {
	EtcdMember       *Nodegroup `yaml:"etcd-member,omitempty"`
	EtcdEventsMember *Nodegroup `yaml:"etcd-events-member,omitempty"`
	Api              *Nodegroup `yaml:"api,omitempty"`
	Node             *Nodegroup `yaml:"node,omitempty"`
}

// DefaultAttributes represents the zendesk_kubernetes default attributes.
type DefaultAttributes struct {
	Cluster string `yaml:"cluster,omitempty"`
}

// Cluster represents a single cluster definition.
type Cluster struct {
	Name                     string                `yaml:"name"`
	Partition                string                `yaml:"partition,omitempty"`
	Env                      string                `yaml:"env,omitempty"`
	Region                   string                `yaml:"region,omitempty"`
	Amd64AmiTags             *AMITag               `yaml:"amd64_ami_tags,omitempty"`
	Arm64AmiTags             *AMITag               `yaml:"arm64_ami_tags,omitempty"`
	CookbookURL              string                `yaml:"cookbook_url,omitempty"`
	NodeRolloutBatchSize     string                `yaml:"node_rollout_batch_size,omitempty"`
	SlackChannels            []string              `yaml:"slack_channels,omitempty"`
	AdditionalSshTeams       []string              `yaml:"additional_ssh_teams,omitempty"`
	Nodegroups               *Nodegroups           `yaml:"nodegroups,omitempty"`
	DefaultAttributes        *DefaultAttributes    `yaml:"default_attributes,omitempty"`
	AwsProfile               string                `yaml:"aws_profile,omitempty"`
	AccountID                string                `yaml:"account_id,omitempty"`
	ImmutableEni             bool                  `yaml:"immutable_eni,omitempty"`
	DnsZone                  string                `yaml:"dns_zone,omitempty"`
	UseNlb                   string                `yaml:"use_nlb,omitempty"`
	CreateLegacyLb           string                `yaml:"create_legacy_lb,omitempty"`
	ApiInternalLbCertArn     string                `yaml:"api_internal_lb_certificate_arn,omitempty"`
	VpnCidrs                 []string              `yaml:"vpn_cidrs,omitempty"`
	EtcdAzNames              []string              `yaml:"etcd_az_names,omitempty"`
	SshSourceNetwork         string                `yaml:"ssh_source_network,omitempty"`
	EtcdClientNetwork        string                `yaml:"etcd_client_network,omitempty"`
	ApiClientNetwork         string                `yaml:"api_client_network,omitempty"`
	NodeSourceNetwork        string                `yaml:"node_source_network,omitempty"`
	Kube2iamDisabled         []string              `yaml:"kube2iam_disabled,omitempty"`
	UseRecordInjector        string                `yaml:"use_record_injector,omitempty"`
	DrainTimeout             int                   `yaml:"drain_timeout,omitempty"`
	SubnetTags               map[string]string     `yaml:"subnet_tags,omitempty"`
	AdditionalSecurityGroups map[string][][]string `yaml:"additional_security_groups,omitempty"`
	VpcTags                  map[string]string     `yaml:"vpc_tags,omitempty"`
}

// ClusterMetadata is the top-level struct.
type ClusterMetadata struct {
	Clusters []*Cluster `yaml:"clusters"`
}

// --- DSL Engine ---

var (
	// Definition holds the top-level definition as it's being built.
	Definition = &ClusterMetadata{}
	// currentCluster holds the context for the current Cluster being defined.
	currentCluster *Cluster
	// currentAmiTag holds the context for the current AMITag being defined.
	currentAmiTag *AMITag
	// currentNodegroups holds the context for the current Nodegroups being defined.
	currentNodegroups *Nodegroups
	// currentNodegroup holds the context for the current Nodegroup being defined.
	currentNodegroup *Nodegroup
	// currentDefaultAttributes holds the context for the current DefaultAttributes being defined.
	currentDefaultAttributes *DefaultAttributes
)

// Run executes the DSL definition and returns the resulting object.
func Run(dsl func()) *ClusterMetadata {
	dsl()
	return Definition
}

// --- DSL Functions ---

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

// Env sets the environment for the current cluster.
func Env(val string) {
	if currentCluster != nil {
		currentCluster.Env = val
	}
}

// Region sets the region for the current cluster.
func Region(val string) {
	if currentCluster != nil {
		currentCluster.Region = val
	}
}

// Amd64AmiTags defines the amd64 AMI tags for the current cluster.
func Amd64AmiTags(dsl func()) {
	if currentCluster != nil {
		currentCluster.Amd64AmiTags = &AMITag{}
		inAmiTag(currentCluster.Amd64AmiTags, dsl)
	}
}

// Arm64AmiTags defines the arm64 AMI tags for the current cluster.
func Arm64AmiTags(dsl func()) {
	if currentCluster != nil {
		currentCluster.Arm64AmiTags = &AMITag{}
		inAmiTag(currentCluster.Arm64AmiTags, dsl)
	}
}

// Revision sets the revision for the current AMI tag.
func Revision(val string) {
	if currentAmiTag != nil {
		currentAmiTag.Revision = val
	}
}

// Environment sets the environment for the current AMI tag.
func Environment(val string) {
	if currentAmiTag != nil {
		currentAmiTag.Environment = val
	}
}

// BaseName sets the base name for the current AMI tag.
func BaseName(val string) {
	if currentAmiTag != nil {
		currentAmiTag.BaseName = val
	}
}

// CookbookURL sets the cookbook URL for the current cluster.
func CookbookURL(val string) {
	if currentCluster != nil {
		currentCluster.CookbookURL = val
	}
}

// NodeRolloutBatchSize sets the node rollout batch size for the current cluster.
func NodeRolloutBatchSize(val string) {
	if currentCluster != nil {
		currentCluster.NodeRolloutBatchSize = val
	}
}

// SlackChannels sets the slack channels for the current cluster.
func SlackChannels(vals ...string) {
	if currentCluster != nil {
		currentCluster.SlackChannels = vals
	}
}

// AdditionalSshTeams sets the additional SSH teams for the current cluster.
func AdditionalSshTeams(vals ...string) {
	if currentCluster != nil {
		currentCluster.AdditionalSshTeams = vals
	}
}

// NodegroupsDef defines the nodegroups for the current cluster.
func NodegroupsDef(dsl func()) {
	if currentCluster != nil {
		currentCluster.Nodegroups = &Nodegroups{}
		inNodegroups(currentCluster.Nodegroups, dsl)
	}
}

// EtcdMember defines the etcd-member nodegroup.
func EtcdMember(dsl func()) {
	if currentNodegroups != nil {
		currentNodegroups.EtcdMember = &Nodegroup{}
		inNodegroup(currentNodegroups.EtcdMember, dsl)
	}
}

// EtcdEventsMember defines the etcd-events-member nodegroup.
func EtcdEventsMember(dsl func()) {
	if currentNodegroups != nil {
		currentNodegroups.EtcdEventsMember = &Nodegroup{}
		inNodegroup(currentNodegroups.EtcdEventsMember, dsl)
	}
}

// Api defines the api nodegroup.
func Api(dsl func()) {
	if currentNodegroups != nil {
		currentNodegroups.Api = &Nodegroup{}
		inNodegroup(currentNodegroups.Api, dsl)
	}
}

// Node defines the node nodegroup.
func Node(dsl func()) {
	if currentNodegroups != nil {
		currentNodegroups.Node = &Nodegroup{}
		inNodegroup(currentNodegroups.Node, dsl)
	}
}

// Autoscale sets the autoscale property for the current nodegroup.
func Autoscale(val bool) {
	if currentNodegroup != nil {
		currentNodegroup.Autoscale = &val
	}
}

// Type sets the type for the current nodegroup.
func Type(val string) {
	if currentNodegroup != nil {
		currentNodegroup.Type = val
	}
}

// RootEbsSize sets the root EBS size for the current nodegroup.
func RootEbsSize(val string) {
	if currentNodegroup != nil {
		currentNodegroup.RootEbsSize = val
	}
}

// MaxSize sets the max size for the current nodegroup.
func MaxSize(val int) {
	if currentNodegroup != nil {
		currentNodegroup.MaxSize = &val
	}
}

// ScaleDownGracePeriod sets the scale down grace period for the current nodegroup.
func ScaleDownGracePeriod(val string) {
	if currentNodegroup != nil {
		currentNodegroup.ScaleDownGrace = val
	}
}

// DrainTimeoutNodegroup sets the drain timeout for the current nodegroup.
func DrainTimeoutNodegroup(val string) {
	if currentNodegroup != nil {
		currentNodegroup.DrainTimeout = val
	}
}

// InstanceTypes sets the instance types for the current nodegroup.
func InstanceTypes(vals ...string) {
	if currentNodegroup != nil {
		currentNodegroup.InstanceTypes = vals
	}
}

// DefaultAttributesDef defines the default attributes for the current cluster.
func DefaultAttributesDef(dsl func()) {
	if currentCluster != nil {
		currentCluster.DefaultAttributes = &DefaultAttributes{}
		inDefaultAttributes(currentCluster.DefaultAttributes, dsl)
	}
}

// ClusterAttribute sets the cluster attribute within default_attributes.
func ClusterAttribute(val string) {
	if currentDefaultAttributes != nil {
		currentDefaultAttributes.Cluster = val
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

// --- Context Helpers ---

func inCluster(cluster *Cluster, dsl func()) {
	original := currentCluster
	currentCluster = cluster
	dsl()
	currentCluster = original
}

func inAmiTag(tag *AMITag, dsl func()) {
	original := currentAmiTag
	currentAmiTag = tag
	dsl()
	currentAmiTag = original
}

func inNodegroups(ngs *Nodegroups, dsl func()) {
	original := currentNodegroups
	currentNodegroups = ngs
	dsl()
	currentNodegroups = original
}

func inNodegroup(ng *Nodegroup, dsl func()) {
	original := currentNodegroup
	currentNodegroup = ng
	dsl()
	currentNodegroup = original
}

func inDefaultAttributes(da *DefaultAttributes, dsl func()) {
	original := currentDefaultAttributes
	currentDefaultAttributes = da
	dsl()
	currentDefaultAttributes = original
}
