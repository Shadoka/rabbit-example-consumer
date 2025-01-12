# rabbit-example-consumer

Quick and dirty application to consume messages out of rabbit queues.

## Rabbit prerequisites

A local installation under localhost:5672 with credentials guest/guest is required.

Queues(exchange/routingKey/name):

* addition.direct/with-reply/addition
* addition-reply/#/reply
* addition.direct/without-reply/addition-no-reply

For details of the format of input and output message [take a look at the relevant structures](data/data.go).