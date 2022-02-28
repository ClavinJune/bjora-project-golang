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
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	// 	return aws.Endpoint{
	// 		PartitionID:       "aws",
	// 		URL:               "http://localhost:9000",
	// 		SigningRegion:     region,
	// 		HostnameImmutable: true,
	// 	}, nil
	// })
	//
	// cfg, err := config.LoadDefaultConfig(context.Background(),
	// 	config.WithEndpointResolverWithOptions(resolver))
	// if err != nil {
	// 	panic(err)
	// }
	//
	// s3Client := s3.NewFromConfig(cfg)
	// _, err = s3Client.PutObject(context.Background(), &s3.PutObjectInput{
	// 	Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
	// 	Key:    aws.String("testing.txt"),
	// 	Body:   strings.NewReader("ehehehhe"),
	// })
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// resp, err := http.Get("http://localhost:9000/pictures/testing.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	_ = resp.Body.Close()
	// }()
	//
	// b, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(string(b))
}
