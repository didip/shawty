package storage

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

type S3 struct {
	Bucket *s3.Bucket
}

func (s *S3) Init(auth aws.Auth, region aws.Region, bucketName string) error {
	client := s3.New(auth, region)
	buckets, err := client.ListBuckets()
	if err != nil {
		return err
	}

	for _, b := range buckets.Buckets {
		if b.Name == bucketName {
			s.Bucket = &b
			break
		}
	}
	if s.Bucket == nil {
		s.Bucket = client.Bucket(bucketName)
		err = s.Bucket.PutBucket(s3.BucketOwnerFull)
	}

	return err
}

func NewS3(auth aws.Auth, region aws.Region, bucketName string) (*S3, error) {
	s := new(S3)
	return s, s.Init(auth, region, bucketName)
}

func (s *S3) Save(url string) (string, error) {
	var (
		code string
		err  error
	)

	for i := 0; i < 10; i++ {
		code = getRandomString(8)

		if _, err = s.Bucket.GetKey(code); err != nil {
			err = s.Bucket.Put(code, []byte(url), "text/plain", s3.BucketOwnerFull)
			if err == nil {
				return code, err
			}
		}
	}

	return "", ErrCodeNotSet
}

func (s *S3) SaveName(code string, url string) error {
	return s.Bucket.Put(code, []byte(url), "text/plain", s3.BucketOwnerFull)
}

func (s *S3) Load(code string) (string, error) {
	url, err := s.Bucket.Get(code)
	if s3err, ok := err.(*s3.Error); ok && s3err.Code == "NoSuchKey" {
		return "", ErrCodeNotSet
	}
	return string(url), err
}
