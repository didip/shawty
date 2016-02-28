package migrations

import (
	"log"
	"regexp"
	"sync"

	"github.com/mitchellh/goamz/s3"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

// This function is designed to assist with migrating an S3 storage with
// migrating from the original implementation to the V2 design. I.e. get to this:
//	 sha256(short)/
//	               long -> url
//	               change_history/
//	                              TimeInRFC3339Nano() -> {URL: oldURL, Owner: "TODO"}
// from this:
//	short -> url
func MigrateS3FromV1ToV2(storage *storage.S3, writenew bool, deleteold bool) {
	log.Println("Begining S3 migration from v1 design to v2")

	s3b := storage.Bucket

	var (
		wg     sync.WaitGroup
		marker string
	)

	for retryCount := 0; retryCount < 10; {
		resp, err := s3b.List("", "", marker, 1000)
		if err != nil {
			log.Println("Failed to list S3 bucket because: %v", err)
			retryCount++
			continue
		}

		log.Printf("List found %v available items", len(resp.Contents))
		for _, k := range resp.Contents {
			if isVersionedKey(k.Key) {
				log.Printf("Skipping '%v' because it looks like a versioned key", k.Key)
				continue
			}

			wg.Add(1)
			go func(k s3.Key) {
				defer wg.Done()

				b, err := s3b.Get(k.Key)
				if err != nil {
					log.Printf("Failed to get url from key '%s' because: %v", k.Key, err)
					return
				}

				short := k.Key
				url := string(b)

				log.Printf("Migrating the pair '%v'->'%v' to its new home", short, url)
				if writenew {
					err = storage.SaveName(short, url)
					if err != nil {
						log.Printf("Failed to migrate short '%s' to its new home because: %s", short, err)
						return // This is pretty important :D
					}

					if deleteold {
						if err := s3b.Del(short); err != nil {
							log.Printf("Failed to clean out old short code '%s' because: %s", short, err)
						}
					}
				}
			}(k)
		}

		wg.Wait()

		if resp.IsTruncated {
			marker = resp.NextMarker
		} else {
			break
		}
	}
}

var versionedKey = regexp.MustCompile("^v\\d+\\/.+")

func isVersionedKey(key string) bool {
	if versionedKey.MatchString(key) {
		return true
	}
	return false
}
