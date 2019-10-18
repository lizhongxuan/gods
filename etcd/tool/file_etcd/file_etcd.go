package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	editor := "/usr/bin/vim"
	if s := os.Getenv("EDITOR"); s != "" {
		editor = s
	}
	log.Print("Getenv EDITOR :", editor)
	endpoints := []string{"http://127.0.0.1:2379"}
	if s := os.Getenv("ETCD_ENDPOINTS"); s != "" {
		endpoints = strings.Split(s, ",")
	}
	log.Print("Getenv ETCD_ENDPOINTS :", endpoints)

	config := clientv3.Config{
		Endpoints:            endpoints,
		DialTimeout:          time.Second * 10,
		DialKeepAliveTime:    time.Second * 10,
		DialKeepAliveTimeout: time.Second * 30,
	}
	cli, err := clientv3.New(config)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	log.Print("etcd connected")

	key := os.Args[1]
	if key == "" {
		log.Fatal("key arg is nil")
	}

	if value, err := ioutil.ReadFile(key); err != nil {
		log.Fatal(err)
	} else if _, err := cli.Put(context.Background(), key, string(value)); err != nil {
		log.Fatal(err)
	}
	log.Println(key, "saved")
}
