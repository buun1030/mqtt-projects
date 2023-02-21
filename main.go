package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	var broker = "test.mosquitto.org"
	var port = 1883
	// create a new MQTT client and connect to an MQTT broker running on "test.mosquitto.org"
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client_kinji94")
	// opts.SetUsername("emqx")
	// opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	// publish(client)

	client.Disconnect(250)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("kinji94_gendata", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "kinji94/gendata"

	// Set up a channel to receive OS signals (e.g. SIGINT or SIGTERM)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Set up a channel to receive MQTT messages
	msgchan := make(chan mqtt.Message, 10)

	// Subscribe to the MQTT topic
	if token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		// Pass a callback function to handle incoming messages
		msgchan <- msg // The callback function simply sends the message to the msgchan channel
	}); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Wait for OS signals or MQTT messages
	for {
		select {
		case sig := <-sigchan:
			fmt.Printf("Received signal %v\n", sig)
			// Disconnect from the MQTT broker and exit the program
			client.Disconnect(250)
			os.Exit(0)
		case msg := <-msgchan:
			fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())
		}
	}
}
