/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EndpointDetailsObservation struct {
}

type EndpointDetailsParameters struct {

	// A list of address allocation IDs that are required to attach an Elastic IP address to your SFTP server's endpoint. This property can only be used when endpoint_type is set to VPC.
	// +kubebuilder:validation:Optional
	AddressAllocationIds []*string `json:"addressAllocationIds,omitempty" tf:"address_allocation_ids,omitempty"`

	// A list of security groups IDs that are available to attach to your server's endpoint. If no security groups are specified, the VPC's default security groups are automatically assigned to your endpoint. This property can only be used when endpoint_type is set to VPC.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A list of subnet IDs that are required to host your SFTP server endpoint in your VPC. This property can only be used when endpoint_type is set to VPC.
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The ID of the VPC endpoint. This property can only be used when endpoint_type is set to VPC_ENDPOINT
	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The VPC ID of the virtual private cloud in which the SFTP server's endpoint will be hosted. This property can only be used when endpoint_type is set to VPC.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type OnUploadObservation struct {
}

type OnUploadParameters struct {

	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	// +kubebuilder:validation:Required
	ExecutionRole *string `json:"executionRole" tf:"execution_role,omitempty"`

	// A unique identifier for the workflow.
	// +kubebuilder:validation:Required
	WorkflowID *string `json:"workflowId" tf:"workflow_id,omitempty"`
}

type ServerObservation struct {

	// Amazon Resource Name  of Transfer Server
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The endpoint of the Transfer Server
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// This value contains the message-digest algorithm  hash of the server's host key. This value is equivalent to the output of the ssh-keygen -l -E md5 -f my-new-server-key command.
	HostKeyFingerprint *string `json:"hostKeyFingerprint,omitempty" tf:"host_key_fingerprint,omitempty"`

	// he Server ID of the Transfer Server
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ServerParameters struct {

	// The Amazon Resource Name  of the AWS Certificate Manager  certificate. This is required when protocols is set to FTPS
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/acm/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateRef *v1.Reference `json:"certificateRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CertificateSelector *v1.Selector `json:"certificateSelector,omitempty" tf:"-"`

	// The directory service ID of the directory service you want to connect to with an identity_provider_type of AWS_DIRECTORY_SERVICE.
	// +kubebuilder:validation:Optional
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// The domain of the storage system that is used for file transfers. Valid values are: S3 and EFS. The default value is S3.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The virtual private cloud  endpoint settings that you want to configure for your SFTP server. Fields documented below.
	// +kubebuilder:validation:Optional
	EndpointDetails []EndpointDetailsParameters `json:"endpointDetails,omitempty" tf:"endpoint_details,omitempty"`

	// The type of endpoint that you want your SFTP server connect to. If you connect to a VPC , your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set PUBLIC.  Defaults to PUBLIC.
	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is false. This option only applies to servers configured with a SERVICE_MANAGED identity_provider_type.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The ARN for a lambda function to use for the Identity provider.
	// +kubebuilder:validation:Optional
	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	// RSA private key .
	// +kubebuilder:validation:Optional
	HostKeySecretRef *v1.SecretKeySelector `json:"hostKeySecretRef,omitempty" tf:"-"`

	// The mode of authentication enabled for this service. The default value is SERVICE_MANAGED, which allows you to store and access SFTP user credentials within the service. API_GATEWAY indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using AWS_DIRECTORY_SERVICE will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors. Use the AWS_LAMBDA value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the function argument.
	// +kubebuilder:validation:Optional
	IdentityProviderType *string `json:"identityProviderType,omitempty" tf:"identity_provider_type,omitempty"`

	// Amazon Resource Name  of the IAM role used to authenticate the user account with an identity_provider_type of API_GATEWAY.
	// +kubebuilder:validation:Optional
	InvocationRole *string `json:"invocationRole,omitempty" tf:"invocation_role,omitempty"`

	// Amazon Resource Name  of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
	// +kubebuilder:validation:Optional
	LoggingRole *string `json:"loggingRole,omitempty" tf:"logging_role,omitempty"`

	// Specify a string to display when users connect to a server. This string is displayed after the user authenticates. The SFTP protocol does not support post-authentication display banners.
	// +kubebuilder:validation:Optional
	PostAuthenticationLoginBannerSecretRef *v1.SecretKeySelector `json:"postAuthenticationLoginBannerSecretRef,omitempty" tf:"-"`

	// Specify a string to display when users connect to a server. This string is displayed before the user authenticates.
	// +kubebuilder:validation:Optional
	PreAuthenticationLoginBannerSecretRef *v1.SecretKeySelector `json:"preAuthenticationLoginBannerSecretRef,omitempty" tf:"-"`

	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint. This defaults to SFTP . The available protocols are:
	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the name of the security policy that is attached to the server. Possible values are TransferSecurityPolicy-2018-11, TransferSecurityPolicy-2020-06, and  TransferSecurityPolicy-FIPS-2020-06. Default value is: TransferSecurityPolicy-2018-11.
	// +kubebuilder:validation:Optional
	SecurityPolicyName *string `json:"securityPolicyName,omitempty" tf:"security_policy_name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// - URL of the service endpoint used to authenticate users with an identity_provider_type of API_GATEWAY.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Specifies the workflow details. See Workflow Details below.
	// +kubebuilder:validation:Optional
	WorkflowDetails []WorkflowDetailsParameters `json:"workflowDetails,omitempty" tf:"workflow_details,omitempty"`
}

type WorkflowDetailsObservation struct {
}

type WorkflowDetailsParameters struct {

	// A trigger that starts a workflow: the workflow begins to execute after a file is uploaded. See Workflow Detail below.
	// +kubebuilder:validation:Optional
	OnUpload []OnUploadParameters `json:"onUpload,omitempty" tf:"on_upload,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API. Provides a AWS Transfer Server resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
