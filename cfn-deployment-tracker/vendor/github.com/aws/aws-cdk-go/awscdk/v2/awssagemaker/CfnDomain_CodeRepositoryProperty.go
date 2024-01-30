package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeRepositoryProperty := &CodeRepositoryProperty{
//   	RepositoryUrl: jsii.String("repositoryUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-coderepository.html
//
type CfnDomain_CodeRepositoryProperty struct {
	// A CodeRepository (valid URL) to be used within Jupyter's Git extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-coderepository.html#cfn-sagemaker-domain-coderepository-repositoryurl
	//
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
}

