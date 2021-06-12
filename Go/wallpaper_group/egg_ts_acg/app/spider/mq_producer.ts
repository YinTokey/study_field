import * as amqp from 'amqplib';

const host = 'amqp://localhost';

async function directProducer(): Promise<void>{

    // 创建链接对象
    const connection = await amqp.connect(host,'heartbeat=60');

    // 获取通道
    const channel = await connection.createChannel();

    // 声明参数
    const exchangeName = 'direct_exchange_name';
    const routingKey = 'direct_routingKey';
    const msg = 'hello world';

    // 交换机
    await channel.assertExchange(exchangeName, 'direct', {
        durable: true,
    });

    // 发送消息
    await channel.publish(exchangeName, routingKey, Buffer.from(msg));

    // 关闭链接
    await channel.close();
    await connection.close();

    console.log('消息发送结束');

}

async function topicProducer(): Promise<void> {
    // 创建链接对象
    const connection = await amqp.connect(host);

    // 获取通道
    const channel = await connection.createChannel();

    /** 改造部分 */
    // 声明参数
    const exchangeName = 'topic_exchange_name';
    const routingKey1 = 'topic_routingKey.test1';
    const routingKey2 = 'topic_routingKey.test2';
    const routingKey3 = 'topic_routingKey.test3.1';
    const routingKey4 = 'topic_routingKey2.test4';
    const routingKey5 = 'topic_routingKey2.test5.1';
    const msg = 'hello world form topic';

    // 交换机
    await channel.assertExchange(exchangeName, 'topic', {
        durable: true,
    });

    // 发送消息
    await channel.publish(exchangeName, routingKey1, Buffer.from(msg + routingKey1));
    await channel.publish(exchangeName, routingKey2, Buffer.from(msg + routingKey2));
    await channel.publish(exchangeName, routingKey3, Buffer.from(msg + routingKey3));
    await channel.publish(exchangeName, routingKey4, Buffer.from(msg + routingKey4));
    await channel.publish(exchangeName, routingKey5, Buffer.from(msg + routingKey5));
    /** 改造部分 */

    // 关闭链接
    await channel.close();
    await connection.close();
}

// directProducer();
topicProducer();
