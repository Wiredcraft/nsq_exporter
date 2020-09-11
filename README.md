# NSQ Exporter

[![Build Status](https://travis-ci.com/Wiredcraft/nsq_exporter.svg?branch=master)](https://travis-ci.com/Wiredcraft/nsq_exporter) 

NSQ exporter for prometheus.io, written in go.

## Usage

    podman run --rm -d --name nsq_exporter -l nsqd:nsqd -p 9117:9117 Wiredcraft/nsq_exporter:latest -nsqd.addr=http://nsqd:4151 -collect=stats.topics,stats.channels,stats.clients

## Building

    make

    OR

    go get -u github.com/Wiredcraft/nsq_exporter
    go install github.com/Wiredcraft/nsq_exporter

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request
