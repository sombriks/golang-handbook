package sample_kafka

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/kafka"
	"testing"
)

type SampleKafkaTestSuite struct {
	suite.Suite
	kafkaContainer *kafka.KafkaContainer
}

func (suite *SampleKafkaTestSuite) SetupTest() {
	ctx := context.Background()
	var err error
	suite.kafkaContainer, err = kafka.Run(ctx, "confluentinc/confluent-local:7.5.0",
		kafka.WithClusterID("sample-cluster"))

	require.Nil(suite.T(), err)
}

func TestBootstrap(t *testing.T) {
	suite.Run(t, new(SampleKafkaTestSuite))
}

func (suite *SampleKafkaTestSuite) TestProducer() {

}

func (suite *SampleKafkaTestSuite) TestConsumer() {

}
