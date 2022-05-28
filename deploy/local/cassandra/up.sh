#!/bin/bash
docker-compose up -d
docker network create --subnet=172.30.0.0/16 --gateway=172.30.254.254  cassandra
docker run --name cassandra-1 \
--network cassandra \
--ip 172.30.1.1 \
-e CASSANDRA_BROADCAST_ADDRESS=172.30.1.1 \
-d \
cassandra:4.0.4

docker run --name cassandra-2 \
--network cassandra \
--ip 172.30.1.2 \
-e CASSANDRA_BROADCAST_ADDRESS=172.30.1.2 \
-e CASSANDRA_SEEDS=172.30.1.1 \
-d \
cassandra:4.0.4

docker run --name cassandra-3 \
--network cassandra \
--ip 172.30.1.3 \
-e CASSANDRA_BROADCAST_ADDRESS=172.30.1.3 \
-e CASSANDRA_SEEDS=172.30.1.1 \
-d \
cassandra:4.0.4