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
- [NSQ Golang Meetup](https://speakerdeck.com/snakes/nsq-nyc-golang-meetup)
