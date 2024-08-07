package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/prometheus/prompb"
	"github.com/prometheus/prometheus/promql/parser"
)

type PrometheusRemoteStorageAdapter interface {
	Write(records *prompb.WriteRequest) error
	Read(request *prompb.ReadRequest) (*prompb.ReadResponse, error)
}

func serve(addr string, storageAdapter PrometheusRemoteStorageAdapter) error {
	log.Print("Start remote read server")
	http.Handle("/write", writeHandler(storageAdapter))
	http.Handle("/read", readHandler(storageAdapter))
	return http.ListenAndServe(addr, nil)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func ParseSeriesDescToTimeSeries(input string) (ts prompb.TimeSeries, err error) {
	lab, val, err := parser.ParseSeriesDesc(input)
	if err != nil {
		return
	}

	labels := []prompb.Label{}
	for _, l := range lab {
		labels = append(labels, prompb.Label{
			Name:                 l.Name,
			Value:                l.Value,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     []byte{},
			XXX_sizecache:        0,
		})
	}

	samples := []prompb.Sample{}

	for i, v := range val {

		t := makeTimestamp() - int64(time.Minute.Milliseconds())*int64(len(val)-i)

		samples = append(samples, prompb.Sample{
			Value:                v.Value,
			Timestamp:            t,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     []byte{},
			XXX_sizecache:        0,
		})
	}

	ts = prompb.TimeSeries{
		Labels:               labels,
		Samples:              samples,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     []byte{},
		XXX_sizecache:        0,
	}
	return
}

func main() {
	configPath := flag.String("config", "", "Path to the configuration file")
	port := flag.String("port", "9999", "Port to serve on")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	config := Configuration{}
	err := config.LoadConfig(*configPath)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	stor := NewFakeStorage(config.InputSeries)

	address := fmt.Sprintf(":%s", *port)
	if err := serve(address, stor); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
