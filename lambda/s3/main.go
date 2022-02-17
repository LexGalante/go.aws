package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ObjectS3 struct {
	BucketName string
	ObjectKey  string
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		objectS3 := ObjectS3{
			BucketName: record.S3.Bucket.Name,
			ObjectKey:  record.S3.Object.Key,
		}

		go storeMetadataInDynamodb(objectS3)
	}
}

func storeMetadataInDynamodb(objectS3 ObjectS3) {
	//TODO
}
