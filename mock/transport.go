package mock

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/execute"
	"github.com/syoliver-se/flux/internal/errors"
)

type Transport struct {
	ProcessMessageFn func(m execute.Message) error
}

func (t *Transport) ProcessMessage(m execute.Message) error {
	return t.ProcessMessageFn(m)
}

func (t *Transport) RetractTable(id execute.DatasetID, key flux.GroupKey) error {
	return errors.New(codes.Unimplemented)
}
func (t *Transport) Process(id execute.DatasetID, tbl flux.Table) error {
	return errors.New(codes.Unimplemented)
}
func (t *Transport) UpdateWatermark(id execute.DatasetID, ts execute.Time) error {
	return errors.New(codes.Unimplemented)
}
func (t *Transport) UpdateProcessingTime(id execute.DatasetID, ts execute.Time) error {
	return errors.New(codes.Unimplemented)
}
func (t *Transport) Finish(id execute.DatasetID, err error) {
}
