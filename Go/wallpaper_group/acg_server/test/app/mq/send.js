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
            const queue = 'hello1';
            const msg = 'Hello world';

            channel.assertQueue(queue, {
                durable: false
            });

            channel.sendToQueue(queue, Buffer.from(msg));
            console.log(' [x] Sent %s', msg);
        });

        setTimeout(function() {
            connection.close();
            process.exit(0);
        }, 500);
    });
}

main();
