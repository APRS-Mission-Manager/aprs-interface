package amazon

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type APRSPacket struct {
	Callsign  string
	Timestamp int64
	Packet    string
}

func (amazonApi AmazonAPI) LogAPRSPacket(callsign string, timestamp int64, rawPacket string) {
	packet := APRSPacket{Callsign: callsign, Timestamp: timestamp, Packet: rawPacket}
	tableName := amazonApi.appConfig.Amazon.DBAPRSLog.Name

	av, err := attributevalue.MarshalMap(packet)
	if err != nil {
		log.Error().AnErr("error", err).Msg("[AmazonAPI] Failed while marshalling aprs packet.")
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &tableName,
	}

	_, err = amazonApi.dynamodbClient.PutItem(context.TODO(), input)
	if err != nil {
		log.Error().AnErr("error", err).Msg("[AmazonAPI] Failed while putting item into dynamodb.")
	}
}
