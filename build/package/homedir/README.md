# Welcome to merakibeat 6.3.2

Sends events to Elasticsearch or Logstash

## Getting Started

To get started with merakibeat, you need to set up Elasticsearch on your localhost first. After that, start merakibeat with:

     ./merakibeat -c merakibeat.yml -e

This will start the beat and send the data to your Elasticsearch instance. To load the dashboards for merakibeat into Kibana, run:

    ./merakibeat setup -e

For further steps visit the [Getting started](https://www.elastic.co/guide/en/beats/merakibeat/6.3/merakibeat-getting-started.html) guide.

## Documentation

Visit [Elastic.co Docs](https://www.elastic.co/guide/en/beats/merakibeat/6.3/index.html) for the full merakibeat documentation.

## Release notes

https://www.elastic.co/guide/en/beats/libbeat/6.3/release-notes-6.3.2.html
