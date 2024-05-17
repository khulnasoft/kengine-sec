package ingesters

import (
	"context"
	"encoding/json"

	"github.com/khulnasoft/kengine/khulnasoft_utils/directory"
	"github.com/khulnasoft/kengine/khulnasoft_utils/log"
	"github.com/khulnasoft/kengine/khulnasoft_utils/utils"
	ingestersUtil "github.com/khulnasoft/kengine/khulnasoft_utils/utils/ingesters"
	"github.com/twmb/franz-go/pkg/kgo"
)

type SecretIngester struct{}

func NewSecretIngester() KafkaIngester[[]map[string]interface{}] {
	return &SecretIngester{}
}

func (tc *SecretIngester) Ingest(
	ctx context.Context,
	cs []map[string]interface{},
	ingestC chan *kgo.Record,
) error {
	tenantID, err := directory.ExtractNamespace(ctx)
	if err != nil {
		return err
	}

	rh := []kgo.RecordHeader{
		{Key: "namespace", Value: []byte(tenantID)},
	}

	for _, c := range cs {
		cb, err := json.Marshal(c)
		if err != nil {
			log.Error().Msg(err.Error())
		} else {
			ingestC <- &kgo.Record{
				Topic:   utils.SecretScan,
				Value:   cb,
				Headers: rh,
			}
		}
	}

	return nil
}

type SecretScanStatusIngester struct{}

func NewSecretScanStatusIngester() KafkaIngester[[]ingestersUtil.SecretScanStatus] {
	return &SecretScanStatusIngester{}
}

func (tc *SecretScanStatusIngester) Ingest(
	ctx context.Context,
	statuses []ingestersUtil.SecretScanStatus,
	ingestC chan *kgo.Record,
) error {
	tenantID, err := directory.ExtractNamespace(ctx)
	if err != nil {
		return err
	}

	rh := []kgo.RecordHeader{
		{Key: "namespace", Value: []byte(tenantID)},
	}

	for _, c := range statuses {
		cb, err := json.Marshal(c)
		if err != nil {
			log.Error().Msg(err.Error())
		} else {
			ingestC <- &kgo.Record{
				Topic:   utils.SecretScanStatus,
				Value:   cb,
				Headers: rh,
			}
		}
	}

	return nil
}
