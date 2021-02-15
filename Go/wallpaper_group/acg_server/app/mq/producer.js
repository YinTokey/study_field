'use strict';
const amqp = require('amqplib/callback_api');
let channel;
let queue;

class Producer {

    static initMqProducer() {
        amqp.connect('amqp://localhost', function(error0, connection) {
            if (error0) {
                throw error0;
            }
            // connection.createChannel(function(error1, innerChannel) {
            //     if (error1) {
            //         throw error1;
            //     }
            //     queue = 'acg.transfer';
            //     //  channel = innerChannel;
            //     innerChannel.assertQueue(queue, {
            //         durable: true
            //     });
            //
            // });

            setTimeout(function() {
                connection.close();
                process.exit(0);
            }, 500);
        });
    }

    static publish(msg) {

        channel.sendToQueue(queue, Buffer.from(msg));
        console.log(' [x] Sent %s', msg);

    }

}

module.exports = Producer;
