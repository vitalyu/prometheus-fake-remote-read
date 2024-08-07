package main

import (
	"github.com/prometheus/prometheus/prompb"
)

type FakeStorage struct {
	inputSeries []Series
	timeSeries  []*prompb.TimeSeries
}

func (fs *FakeStorage) processSeries() (err error) {
	for _, is := range fs.inputSeries {
		ts, err := ParseSeriesDescToTimeSeries(is.Series + " " + is.Values)
		if err != nil {
			return err
		}
		fs.timeSeries = append(fs.timeSeries, &ts)
	}
	return nil
}

func NewFakeStorage(inputSeries []Series) PrometheusRemoteStorageAdapter {
	fs := FakeStorage{
		inputSeries: inputSeries,
		timeSeries:  []*prompb.TimeSeries{},
	}
	_ = fs.processSeries()
	return &fs
}

func (fs *FakeStorage) Read(request *prompb.ReadRequest) (*prompb.ReadResponse, error) {

	testRequest := request
	_ = testRequest

	var tsFiltered []*prompb.TimeSeries
	_ = tsFiltered

	for _, ts := range fs.timeSeries {
		cnt := 0
		for _, tsLabel := range ts.Labels {
			for _, requestLabel := range request.Queries[0].Matchers {
				if requestLabel.Name == tsLabel.Name && requestLabel.Value == tsLabel.Value {
					cnt++
				}
			}
		}
		if cnt == len(request.Queries[0].Matchers) {
			tsFiltered = append(tsFiltered, ts)
		}
	}

	qr := prompb.QueryResult{
		Timeseries:           tsFiltered, // Timeseries:           fs.timeSeries,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     []byte{},
		XXX_sizecache:        0,
	}

	resp := prompb.ReadResponse{
		Results: []*prompb.QueryResult{
			&qr,
		},
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     []byte{},
		XXX_sizecache:        0,
	}
	return &resp, nil
}

func (rs *FakeStorage) Write(records *prompb.WriteRequest) error {
	println("Write")
	return nil
}
