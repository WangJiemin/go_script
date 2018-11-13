package kafka

import (
	"fmt"
	"golang_Learn/synckafka/database"
	//"os"
	//"strings"
	//"sync"
	//"bufio"

	"github.com/Shopify/sarama"
	"github.com/WangJiemin/gocomm/json"
)

/*
func Initialize(confing map[string]string) {
	var kahost, kaport, katopic string

	kahost = confing["ka_host"]
	kaport = confing["ka_port"]
	katopic = confing["ka_topic"]

	var wg sync.WaitGroup
	var logger = log.New(os.Stderr, fmt.Sprintf("[%s]", katopic), log.LstdFlags)

	conf := fmt.Sprintf("%s:%s", kahost, kaport)
	sarama.Logger = logger

	// 链接kafka消息服务器
	consumer, err := sarama.NewConsumer(strings.Split(conf, ","), nil)
	if err != nil {
		logger.Println("Failed to start consumer: %s", err)
	}
}
*/

func SyncToKafka(confing map[string]string, tablename string) {
	var kahost, katopic string

	kahost = confing["ka_host"]
	katopic = confing["ka_topic"]

	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Partitioner = sarama.NewRandomPartitioner

	//producer, err := sarama.NewSyncProducer([]string{strings.Split(fmt.Sprintf("%s", kahost), ",")[0], strings.Split(fmt.Sprintf("%s", kahost), ",")[1], strings.Split(fmt.Sprintf("%s", kahost), ",")[2]}, conf)
	producer, err := sarama.NewSyncProducer([]string{fmt.Sprintf("%s", kahost)}, conf)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	data := database.SELECTTABLES(tablename)

	//jsondata := database.SELECTTABLES(tablename)
	//fmt.Println(jsondata)

	msg := &sarama.ProducerMessage{
		Topic:     katopic,
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	// 生产消息
	for _, v := range data {
		jsondata, _ := json.Marshal(&v)
		fmt.Println(string(jsondata))

		msg.Value = sarama.ByteEncoder(string(jsondata))
		paritition, offset, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println("Send Message Fail")
		}

		fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
	}

}
