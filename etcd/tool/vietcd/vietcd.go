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
	tempFile, err := ioutil.TempFile(os.TempDir(), "vietcd")
	if err != nil {
		log.Fatal(err)
	}
	tempName := tempFile.Name()
	defer func() {
		tempFile.Close()
		os.Remove(tempName)
	}()

	key := os.Args[1]
	resp, err := cli.Get(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Kvs) > 0 {
		if _, err := tempFile.Write(resp.Kvs[0].Value); err != nil {
			log.Fatal(err)
		}
		log.Print("restore into temp file ", tempName)
	}

	procAttr := os.ProcAttr{
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	proc, err := os.StartProcess(editor, []string{editor, tempName}, &procAttr)
	if err != nil {
		log.Fatal(err)
	}
	//log.Print("edit with editor ", editor)

	if state, err := proc.Wait(); err != nil {
		log.Fatalln("editor failed with:", state, err)
	} else {
		log.Print(state)
	}

	if value, err := ioutil.ReadFile(tempName); err != nil {
		log.Fatal(err)
	} else if _, err := cli.Put(context.Background(), key, string(value)); err != nil {
		log.Fatal(err)
	}
	log.Println(key, "saved")

}
