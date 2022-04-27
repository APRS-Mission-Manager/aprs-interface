package amazon

import (
	"context"

	c "github.com/APRS-Mission-Manager/aprs-interface/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/rs/zerolog/log"
)

type AmazonAPI struct {
	appConfig      c.Config
	awsConfig      aws.Config
	dynamodbClient dynamodb.Client
}

func CreateAPI(appConfig c.Config) AmazonAPI {
	log.Info().Msg("[AmazonAPI] Creating AmazonAPI.")
	amazonApi := AmazonAPI{appConfig: appConfig}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("[AmazonAPI] Failed while trying to load config.")
	}
	amazonApi.awsConfig = cfg
	amazonApi.dynamodbClient = *dynamodb.NewFromConfig(cfg)

	return amazonApi
}
