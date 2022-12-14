package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config is application configuration
type Config struct {
	AwsAccessKeyID    string `env:"AWS_ACCESS_KEY_ID"`
	AwsRegion         string `env:"AWS_REGION"`
	CdkDefaultAccount string `env:"CDK_DEFAULT_ACCOUNT"`
	CdkDefaultRegion  string `env:"CDK_DEFAULT_REGION"`

	Vpc struct {
		Name      string `env:"VPC_NAME"`
		CidrBlock string `env:"CIDR_BLOCK"`
	}
	Cluster struct {
		Name  string `env:"CLUSTER_NAME"`
		Image struct {
			RepositoryName string `env:"CLUSTER_IMAGE_REPOSITORY_NAME"`
			Tag            string `env:"CLUSTER_IMAGE_TAG"`
		}
	}
	LoadBalancer struct {
		DomainName     string `env:"LOAD_BALANCER_DOMAIN_NAME"`
		CertificateArn string `env:"LOAD_BALANCER_CERTIFICATE_ARN"`
	}
	Route53 struct {
		DomainName string `env:"ROUTE53_DOMAIN_NAME"`
	}

	App struct {
		GinMode         string `env:"APP_GIN_MODE"`
		MailFromAddress string `env:"APP_MAIL_FROM_ADDRESS"`

		Aws struct {
			AwsAccessKeyID     string `env:"APP_AWS_ACCESS_KEY_ID"`
			AwsSecretAccessKey string `env:"APP_AWS_SECRET_ACCESS_KEY"`
			AwsRegion          string `env:"APP_AWS_REGION"`
		}
		SendGrid struct {
			SendGridAPIKEY string `env:"APP_SENDGRID_API_KEY"`
		}
	}
}

// Load returns configuration made from environment variables
func Load() *Config {
	_ = godotenv.Load()
	c := Config{}
	_ = env.Parse(&c)

	return &c
}
