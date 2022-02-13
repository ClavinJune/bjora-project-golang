// Copyright 2022 ClavinJune/bjora
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/clavinjune/bjora-project-golang/user"

	"github.com/julienschmidt/httprouter"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := httprouter.New()

	u := user.Wire(nil)
	r.GET(u.Store())

	t := 60 * time.Second
	s := &http.Server{
		Handler:           r,
		ReadTimeout:       t,
		ReadHeaderTimeout: t,
		WriteTimeout:      t,
		IdleTimeout:       t,
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		panic(err)
	}

	go func() {
		log.Printf("run at %s", l.Addr())
		if err := s.Serve(l); err != nil {
			log.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	//resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	//	return aws.Endpoint{
	//		PartitionID:       "aws",
	//		URL:               "http://localhost:9000",
	//		SigningRegion:     region,
	//		HostnameImmutable: true,
	//	}, nil
	//})
	//
	//cfg, err := config.LoadDefaultConfig(context.Background(),
	//	config.WithEndpointResolverWithOptions(resolver))
	//if err != nil {
	//	panic(err)
	//}
	//
	//s3Client := s3.NewFromConfig(cfg)
	//_, err = s3Client.PutObject(context.Background(), &s3.PutObjectInput{
	//	Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
	//	Key:    aws.String("testing.txt"),
	//	Body:   strings.NewReader("ehehehhe"),
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//resp, err := http.Get("http://localhost:9000/pictures/testing.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	_ = resp.Body.Close()
	//}()
	//
	//b, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(b))
}
