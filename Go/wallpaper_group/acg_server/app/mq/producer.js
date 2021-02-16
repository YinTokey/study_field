'use strict';
const amqp = require('amqplib/callback_api');

class Producer {

    static sendMsg(msg) {

        amqp.connect('amqp://localhost', function(error0, connection) {
            if (error0) {
                throw error0;
            }
            connection.createChannel(function(error1, channel) {
                if (error1) {
                    throw error1;
                }
                const queue = 'acg.transfer';

                channel.assertQueue(queue, {
                    durable: true
                });

                console.log(msg);

                channel.sendToQueue(queue, Buffer.from(JSON.stringify(msg)));
                console.log(' [x] Sent %s', JSON.stringify(msg));

                channel.close();
            });

        });

    }

}

module.exports = Producer;

