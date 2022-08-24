This lambda is a [custom resource provider](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html)
implemented in [golang](https://github.com/aws/aws-lambda-go/blob/master/cfn/README.md).

AWS Cloudformation invokes the lambda during the creation of the stack and the lambda will return the ARN of a load balancer created in a Elastic Beanstalk environment.

The example-template.yaml file shows how to setup this functionality.

In that file there are 2 resources, one is the "AWS::Serverless::Function" to define the lambda, and the other is the "Custom::DescribeElbEnvironment", this is what invokes the lambda during cloudformation execution, returning the load balancer's ARN. The lambda is invoked by the '!GetAtt GetLbArn.LoadBalancerArn' construct in the template file.
