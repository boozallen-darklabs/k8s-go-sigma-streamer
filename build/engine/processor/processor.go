package processor

import (
	"context"
	"os"
	"strconv"

	"github.com/Adversary-Informed-Defense/singe/pkg/singe"
	"github.com/Jeffail/benthos/v3/public/service"
)

var engine = singe.CreateEngine(getRulesetPath())

func getRulesetPath() string {
	rulesetPath := string(os.Getenv("RULESET_PATH"))
	if rulesetPath != "" {
		return rulesetPath
	} else {
		return "./rules/"
	}
}

func init() {
	// Config spec is empty for now as we don't have any dynamic fields.
	configSpec := service.NewConfigSpec()
	constructor := func(conf *service.ParsedConfig, mgr *service.Resources) (service.Processor, error) {
		return newSigmaProcessor(mgr.Logger(), mgr.Metrics()), nil
	}

	err := service.RegisterProcessor("singe", configSpec, constructor)
	if err != nil {
		panic(err)
	}
}

type sigmaProcessor struct {
	logger       *service.Logger
	countMatches *service.MetricCounter
}

func newSigmaProcessor(logger *service.Logger, metrics *service.Metrics) *sigmaProcessor {
	return &sigmaProcessor{
		logger:       logger,
		countMatches: metrics.NewCounter("matches"),
	}
}

func (s *sigmaProcessor) Process(ctx context.Context, m *service.Message) (service.MessageBatch, error) {
	row, err := m.AsBytes()
	if err != nil {
		return nil, err
	}
	rowStr := string(row[:])

	// TODO: get vendor for second argument of Match() from benthos config.yaml instead of hard coding to json
	matches, didMatch, _ := engine.Match(rowStr, "json")

	// Increment count matches if any sigma rules matched
	// TODO: This only increments once per log, it should increment once per log * number of sigma rules matched
	matches = []byte(strconv.Quote(string(matches[:])))

	if !didMatch {
		s.countMatches.Incr(1)
		return nil, nil
	}
	newMsg := service.NewMessage(matches)
	return []*service.Message{newMsg}, nil
}

func (r *sigmaProcessor) Close(ctx context.Context) error {
	return nil
}
