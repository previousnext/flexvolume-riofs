package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kingpin"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AttachCommand struct {
	Options string
}

func Attach(app *kingpin.Application) {
	c := &AttachCommand{}
	cmd := app.Command("attach", "Attach the volume to the host").Action(c.Run)
	cmd.Arg("options", "Raw JSON set of options").Required().StringVar(&c.Options)
}

func (c *AttachCommand) Run(k *kingpin.ParseContext) error {
	// These are the options provided to us by the Pod spec.
	opts, err := NewOptions(c.Options)
	if err != nil {
		fatal(err)
	}

	// Get the region of the EC2 instance.
	meta := ec2metadata.New(session.New(), &aws.Config{})
	region, err := meta.Region()
	if err != nil {
		fatal(err)
	}

	// Check if the bucket exists, if it does not, create the bucket if it does not exist.
	svc := s3.New(session.New(&aws.Config{Region: aws.String(region)}))
	params := &s3.CreateBucketInput{
		Bucket: aws.String(opts.Name),
	}
	_, err = svc.CreateBucket(params)
	if err != nil {
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			if reqErr.StatusCode() != 409 {
				fatal(err)
			}
		}
	}

	// Turn on bucket versioning.
	_, err = svc.PutBucketVersioning(&s3.PutBucketVersioningInput{
		Bucket: aws.String(opts.Name),
		VersioningConfiguration: &s3.VersioningConfiguration{
			Status: aws.String(s3.BucketVersioningStatusEnabled),
		},
	})
	if err != nil {
		fatal(err)
	}

	b, err := json.Marshal(map[string]string{
		"status": "Success",
		"device": opts.Name,
	})
	if err != nil {
		fatal(err)
	}
	fmt.Println(string(b))

	return nil
}
