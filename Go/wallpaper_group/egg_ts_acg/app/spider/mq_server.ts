import * as amqp from 'amqplib';

const host = 'amqp://localhost';

function fibonacci(n:number) {
    if (n == 0 || n == 1)
        return n;
    else
        return fibonacci(n - 1) + fibonacci(n - 2);
}

async function connect() {
    try {
        const conn = await amqp.connect(host,'heartbeat=60');
        const ch = await conn.createChannel();
        const queue = 'rpc_queue';
        await ch.assertQueue(queue,{durable:false});
        await ch.prefetch(1);
        console.log(' [x] Awaiting RPC requests');
        await ch.consume(queue,msg => {

            const n = parseInt(msg.content.toString());
            console.log('[.] fib(%d)', n);
            const r = fibonacci(n);

            ch.sendToQueue(msg.properties.replyTo,Buffer.from(r.toString()),{
                correlationId: msg.properties.correlationId
            });

            ch.ack(msg);

        });



    } catch (e) {
        console.log(e);
    }

}


connect();
