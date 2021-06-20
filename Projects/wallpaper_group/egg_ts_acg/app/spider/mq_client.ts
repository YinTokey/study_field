import * as amqp from 'amqplib';

const host = 'amqp://localhost';

function generateUuid() {
    return Math.random().toString() +
        Math.random().toString() +
        Math.random().toString();
}

async function connect() {
    try {
        const conn = await amqp.connect(host,'heartbeat=60');
        const ch = await conn.createChannel();

        const q = await ch.assertQueue('',{exclusive:true});
        const correlationId = generateUuid();
        const num = 5;

        console.log(' [x] Requesting fib(%d)', num);

        await ch.consume(q.queue, msg => {

            if (msg.properties.correlationId === correlationId) {
                console.log(' [.] Got %s', msg.content.toString());
                setTimeout(function() {
                    conn.close();
                    process.exit(0);
                }, 500);

            }

        },{noAck: true});

    } catch (e) {
        console.log(e);
    }

}


connect();
