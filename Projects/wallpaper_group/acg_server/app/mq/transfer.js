'use strict';
const amqp = require('amqplib/callback_api');
let queue;

class Transfer {

    static initMqConsumer() {
        amqp.connect('amqp://localhost', function(error0, connection) {
            if (error0) {
                throw error0;
            }
            connection.createChannel(function(error1, channel) {
                if (error1) {
                    throw error1;
                }

                queue = 'acg.transfer';

                channel.assertQueue(queue, {
                    durable: true
                });

                console.log(' [*] Waiting for messages in %s. To exit press CTRL+C', queue);

                channel.consume(queue, function(msg) {

                    console.log(' [x] Received %s', msg.content.toString());
                }, {
                    noAck: true
                });
            });
        });

        const x = { a: 1, b: 2 };
        x.c = 3;

    }

}

module.exports = Transfer;
