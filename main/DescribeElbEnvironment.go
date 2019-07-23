package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

func describeElasticBeanstalkEnvironment(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {

	if event.RequestType == "Delete" {
		return physicalResourceID, data, err
	}

	elasticBeanstalkEnvironmentName, ok := event.ResourceProperties["ElasticBeanstalkEnvironmentName"]
	if !ok {
		return physicalResourceID, data, errors.New("Property 'ElasticBeanstalkEnvironmentName' is missing.")
	}

	loadBalancerArn, err := getLbArn(elasticBeanstalkEnvironmentName.(string))

	data = map[string]interface{}{
		"LoadBalancerArn": loadBalancerArn,
	}

	return physicalResourceID, data, err
}

func main() {
	lambda.Start(cfn.LambdaWrap(describeElasticBeanstalkEnvironment))
}

func getLbArn(elasticBeanstalkEnvironmentName string) (lbArn string, err error) {
	session := session.Must(session.NewSession(aws.NewConfig().WithCredentials(credentials.NewEnvCredentials())))
	elbClient := elasticbeanstalk.New(session)

	describeEnvironmentResourcesOutput, err := elbClient.DescribeEnvironmentResources(&elasticbeanstalk.DescribeEnvironmentResourcesInput{
		EnvironmentName: &elasticBeanstalkEnvironmentName,
	})
	if err != nil {
		return lbArn, err
	}

	lbArn = *describeEnvironmentResourcesOutput.EnvironmentResources.LoadBalancers[0].Name

	return lbArn, err
}
