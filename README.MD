Este lambda é um "custom resource provider"
https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html
Implementado em go
https://github.com/aws/aws-lambda-go/blob/master/cfn/README.md

O lambda é invocado pelo AWS Cloudformation durante o momento da criação de uma stack e retorna o ARN de um load balancer criado em um ambiente Elastic Beanstalk.

O arquivo example-template.yaml mostra como utilizar esta funcionalidade.

Neste arquivo estão declarados 2 recursos, um é o "AWS::Serverless::Function", que deploya o lambda na AWS, o outro é o "Custom::DescribeElbEnvironment", este é o que invoca a lambda durante o cloudformation e retorna o Arn do Loadbalancer
quado é utilizada a funcão '!GetAtt GetLbArn.LoadBalancerArn'