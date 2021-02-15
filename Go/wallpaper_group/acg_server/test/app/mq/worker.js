'use strict';
const amqp = require('amqplib/callback_api');

function main() {

    amqp.connect('amqp://localhost', function(error0, connection) {
        if (error0) {
            throw error0;
        }
        connection.createChannel(function(error1, channel) {
            if (error1) {
                throw error1;
            }

            const queue = 'hello';

            channel.assertQueue(queue, {
                durable: true
            });

            console.log(' [*] Waiting for messages in %s. To exit press CTRL+C', queue);

            channel.consume(queue, function(msg) {
                const secs = msg.content.toString().split('.').length - 1;
                setTimeout(function() {
                    console.log(' [x] Done');
                }, secs * 1000);
                console.log(' [x] Received %s', msg.content.toString());
            }, {
                noAck: true
            });
        });
    });
}

main();
