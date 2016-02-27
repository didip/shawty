package storage

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path"
	"time"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

type S3 struct {
	Bucket         *s3.Bucket
	hashFunc       func(string) string
	storageVersion string
}

func NewS3(auth aws.Auth, region aws.Region, bucketName string) (*S3, error) {
	s := &S3{
		Bucket: s3.New(auth, region).Bucket(bucketName),
		hashFunc: func(s string) string {
			h := sha512.Sum384([]byte(s))
			return base64.StdEncoding.EncodeToString(h[:])
		},
		storageVersion: "v2",
	}

	_, err := s.Bucket.GetBucketContents()
	if s3err, ok := err.(*s3.Error); ok && s3err.Code == "NoSuchBucket" {
		err = s.Bucket.PutBucket(s3.BucketOwnerFull)
	}

	return s, err
}

func (s *S3) saveKey(short, url string) (err error) {
	hashedShort := s.hashFunc(short)

	err = s.Bucket.Put(
		path.Join(s.storageVersion, hashedShort, "long"),
		[]byte(url),
		"text/plain",
		s3.BucketOwnerFull,
	)
	if err != nil {
		return err
	}

	changeLog, err := json.Marshal(
		struct {
			URL  string
			User string
		}{
			url,
			"TODO",
		},
	)
	if err != nil {
		return fmt.Errorf("unable to format change history: %v", err)
	}

	err = s.Bucket.Put(
		path.Join(s.storageVersion, hashedShort, "change_history", time.Now().Format(time.RFC3339Nano)),
		changeLog,
		"application/json",
		s3.BucketOwnerFull,
	)
	if err != nil {
		return fmt.Errorf("unable to save change history: %v", err)
	}

	return nil
}

func (s *S3) Save(url string) (string, error) {
	if _, err := validateURL(url); err != nil {
		return "", err
	}

	var (
		short string
		err   error
	)

	for i := 0; i < 10; i++ {
		short = getRandomString(8)
		pathToShort := path.Join(s.storageVersion, s.hashFunc(short))
		// pathToShort := path.Join(s.storageVersion, s.hashFunc(short), "long")

		if _, err = s.Bucket.GetKey(pathToShort); err != nil {
			return short, s.saveKey(short, url)
		}
	}

	return "", ErrShortNotSet
}

func (s *S3) SaveName(short string, url string) error {
	if err := validateShort(short); err != nil {
		return err
	}
	if _, err := validateURL(url); err != nil {
		return err
	}

	return s.saveKey(short, url)
}

func (s *S3) Load(short string) (string, error) {
	if err := validateShort(short); err != nil {
		return "", err
	}

	url, err := s.Bucket.Get(path.Join(s.storageVersion, s.hashFunc(short), "long"))
	if s3err, ok := err.(*s3.Error); ok && s3err.Code == "NoSuchKey" {
		return "", ErrShortNotSet
	}
	return string(url), err
}
