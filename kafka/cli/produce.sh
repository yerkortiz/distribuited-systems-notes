#!/bin/bash

source kafka_home.sh

$KAFKA_HOME/bin/kafka-console-producer.sh \
  --topic order-received \
  --bootstrap-server localhost:9092