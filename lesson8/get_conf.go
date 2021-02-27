package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

var (
	port        = flag.Int("port", 8080, "Порт прослушивания")
	dbUrl       = flag.String("db_url", "", "URL баы данных")
	jaegerUrl   = flag.String("jaeger_url", "", "URL Jaeger")
	sentryUrlUrl   = flag.String("sentry_url", "", "URL Sentry приложения")
	kafkaServer = flag.String("kafka_server", "", "Имя сервера kafka")
	kafkaPort   = flag.Int("kafka_port", 8080, "Порт сервера kafka")
	someAppID   = flag.String("some_app_id", "", "ID приложения")
	someAppKey  = flag.String("some_app_key", "", "Ключ приложения")
)

func GetConf() {
	flag.Parse()

	filename:= flag.Arg(0)

	if flag.NArg() != 1 {
		log.Fatal("Неверно задано количество аргументов командной строки. Проверьте, что вы задали имя файла.")
	}

	if *port < 1 || *port > 65536 {
		log.Fatal("Некорректное значение порта (port)")
	}
	checkPort(*port, "port")

	checkUrl(*dbUrl, "db_url")
	checkUrl(*jaegerUrl, "jaeger_url")
	checkUrl(*sentryUrlUrl, "sentry_url")
	checkUrl(*kafkaServer, "*kafkaServer")

	checkPort(*kafkaPort, "kafka_port")

	filename = flag.Arg(0)
	kafkaBroker := *kafkaServer + ":" + strconv.Itoa(*kafkaPort)

	log.Printf("port: %d\n", port)
	log.Printf("db_url: %d\n", *dbUrl)
	log.Printf("jaeger_url: %s\n", *jaegerUrl)
	log.Printf("sentry_url: %s\n", *sentryUrlUrl)
	log.Printf("kafka_server: %s\n", kafkaBroker)
	log.Printf("some_app_key: %s\n", *someAppID)
	log.Printf("some_app_key: %s\n", *someAppKey)
	log.Printf("filename: %s\n", filename)
}

func checkPort(port int, nameFlag string) {
	if port < 1 || port > 65536 {
		log.Fatalf("Некорректное значение порта (%s)\n", nameFlag)
	}
}

func checkUrl(url, nameFlag string) {
	if strings.Contains(url, " ") {
		log.Fatalf("Адрес не может сожержать пробелы (%s)\n", nameFlag)
	}
}