# Event-Based Architectures in Go

Code examples from the talk *Event-Based Architectures in Go*, given at Detroit GoLang Meetup on April 27, 2018.

## Repo Contents

All examples assume that you are running Kafka and Zookeeper using the single node configuration from Confluence's [cp-docker-images repository](https://github.com/confluentinc/cp-docker-images/tree/master/examples/kafka-single-node). If you are running a different configuration, you'll just need to go into the source code and adjust the host and port configurations.

### `stdin/`

A small Apache Kafka producer and consumer that pipe text from standard in to each other.

**Setup**

1. Create a topic with the name `stdin` and one partition.
2. Install dependencies

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/stdin
    $ go get -u ./...
    ```
    
3. Start the producer

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/stdin
    $ go run producer.go
    ```
    
4. In a new terminal, start the consumer

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/stdin
    $ go run consumer.go
    ```
    
**Usage**

The producer has text prompts, just follow along.
    
### `partitions/`

**Setup**

1. Create a topic with the name `stdin` and two partitions.
2. Install dependencies

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/partitions
    $ go get -u ./...
    ```
    
3. Start the producer

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/partitions
    $ go run producer.go
    ```
    
4. In a new terminal, start the first consumer

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/partitions
    $ go run consumer.go -partition=0
    ```
    
5. In a new terminal, start the second consumer

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/partitions
    $ go run consumer.go -partition=1
    ```
 
**Usage**

The producer has text prompts, just follow along.

### `commerce/`

This is a really simple REST API that mimics the creation of a cart and produces events based on the actions taken by the sure.

**Setup**

1. Create a topic with the name `carts` and one partition.
2. Install dependencies

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/commerce
    $ go get -u ./...
    ```
    
3. Start the producer/API

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/cmd
    $ go run producer.go
    ```
    
4. Start the consumer that mimics storage behavior

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/cmd
    $ go run storage_consumer.go
    ```
    
5. Start the consumer that mimics audit log behavior

    ```
    $ cd $GOPATH/src/github.com/jmataya/goeventtalk/cmd
    $ go run audit_consumer.go
    ```
    
**Usage**

The folder `scripts/` contains a set of CURL commands that you can use to run HTTP queries against the producer. Feel free to modify the scripts to test out some interesting behavior!

# License

MIT
