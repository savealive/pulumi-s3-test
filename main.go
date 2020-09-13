package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	vault "github.com/pulumi/pulumi-vault/sdk/v2/go/vault/generic"
	"io/ioutil"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "iw-pulumi-dev-test-bucket", &s3.BucketArgs{
			Website: s3.BucketWebsiteArgs{
				IndexDocument: pulumi.String("index.html"),
			},
			ForceDestroy: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		htmlContent, err := ioutil.ReadFile("site/index.html")
		if err != nil {
			return err
		}

		_, err = s3.NewBucketObject(ctx, "index.html", &s3.BucketObjectArgs{
			Bucket:  bucket.ID(),
			Acl: pulumi.String("public-read"),
			ContentType: pulumi.String("text/html"),
			Content: pulumi.String(string(htmlContent)),
		})

		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName1", bucket.ID())
		ctx.Export("endpoint", pulumi.Sprintf("http://%s", bucket.WebsiteEndpoint))
		return nil
	})
}
