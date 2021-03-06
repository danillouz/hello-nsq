# Hello NSQ

Learning about [NSQ](https://nsq.io/).

## What?

NSQ is a realtime distributed messaging platform where you consume messages directly from all producers (messages are pushed to consumers).

It is primarily designed as an _in memory_ message queue, but messages can be written to disk.

## Design Goals

- Simplicity:
  - Configuration
  - Administration
  - High Availability (HA) topologies
- Eliminate Single Point of Failure (SPOF)
- Efficiency

## Components

- `nsqd`: receives, queues, and delivers messages to clients.
- `nsqlookupd`: topology information directory (used for topic lookup).
- `nsqadmin`: Web UI to view cluster stats and perform admin tasks.

## Terminology

- Topic: distinct stream of message (single `nsd` instance can have multiple Topics).
- Channel: independant queue for a topic (single Topic can have multiple Channels).
- Discovery: consumers discover producers by querying `nsqlookupd` (eventual consistent).

## Caveats

- Topics and Channels are created on runtime (you have to start publishing/subscribing).
- Channels receive a copy of all messages for a Topic (multicast).
- Channel consumer receive a portion of the messages from the Channel (evenly distributed).

## Message Quarantees

- At least once delivery.
- Messages received are unordered.
- Messages are not durable by default (see [--mem-queue-size](https://nsq.io/overview/features_and_guarantees.html#messages-are-not-durable-by-default)) and there's no built in replication.

## Best Practices

- A Topic name should describe the data in the stream, e.g. `clicks`.
- A Channel name should describe the work performed by its consumers, e.g. `anaylitics_increment`.

## Resources

- [NSQ Design](https://nsq.io/overview/design.html)
- [NSQ Topology Patterns](https://nsq.io/deployment/topology_patterns.html)
- [NSQ FAQ](https://nsq.io/overview/faq.html)
- [NSQ Golang Meetup Slides](https://speakerdeck.com/snakes/nsq-nyc-golang-meetup)

## Running Locally

### Starting/Stopping all NSQ Services

1. Start `nsqlookupd`, `nsqd` and `nsqadmin` by running:

```
> docker-compose up -d
```

2. Check if it's up:

```
> docker-compose ps
```

3. Ping `nsqd`:

```
> curl http://127.0.0.1:4151/ping

OK
```

4. To stop all containers and clean up run:

```
docker-compose down
```

### Admin UI

The Admin interface can then be used by opening a browser at [http://127.0.0.1:4171](http://127.0.0.1:4171).

### Logs

Logs from the running containers can be viewed with:

```
> docker-compose logs
```

### Pub/Sub Messages Locally

1. Make sure all NSQ services are running.

2. Run the consumer in a separate shell:

```
> go run consumer.go

2019/01/19 17:52:38 INF    1 [test_topic/test_channel] (127.0.0.1:4150) connecting to nsqd
2019/01/19 17:52:38 waiting for msgs..
```

3. Publish a message to `test_topic`:

```
> curl -d 'hello world 1' 'http://127.0.0.1:4151/pub?topic=test_topic'

OK
```

The first published message will also create the topic. Note that the topic name must match the `topicName` constant in `consumer.go`. Additionaly the consumer in `consumer.go` also creates a channel with the name specified by the `chName` constant.

All published messages will be printed to shell that runs `consumer.go`:

```
2019/01/19 17:52:38 INF    1 [TEST_TOPIC/TEST_CHANNEL] (127.0.0.1:4150) connecting to nsqd
2019/01/19 17:52:38 waiting for msgs..
2019/01/19 17:52:38 msg hello world 1
```

Addition message stats can be viewed in the [Admin UI](http://127.0.0.1:4171/topics/test_topic/test_channel).
