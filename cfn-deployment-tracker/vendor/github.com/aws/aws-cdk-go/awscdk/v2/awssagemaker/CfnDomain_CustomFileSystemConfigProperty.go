package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customFileSystemConfigProperty := &CustomFileSystemConfigProperty{
//   	EfsFileSystemConfig: &EFSFileSystemConfigProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//
//   		// the properties below are optional
//   		FileSystemPath: jsii.String("fileSystemPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-customfilesystemconfig.html
//
type CfnDomain_CustomFileSystemConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-customfilesystemconfig.html#cfn-sagemaker-domain-customfilesystemconfig-efsfilesystemconfig
	//
	EfsFileSystemConfig interface{} `field:"optional" json:"efsFileSystemConfig" yaml:"efsFileSystemConfig"`
}

