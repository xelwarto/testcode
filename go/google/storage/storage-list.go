package main

import (
	"fmt"
  //"io/ioutil"
  "io"
	"os"

  "google.golang.org/cloud"
	"google.golang.org/cloud/storage"
  //"google.golang.org/api/storage/v1"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

var bucket = "bucketname"
var jsonKey = []byte(`
`)

func main() {
  fmt.Printf("Storage List: %v\n\n", bucket)

  //jsonKey, err := ioutil.ReadFile("/vagrant/testcode/go/google/storage/key.json")
  //if err != nil {
  //    fmt.Printf("Error: %v\n\n", err)
  //}

  conf, err := google.JWTConfigFromJSON(
      jsonKey,
      storage.ScopeFullControl,
  )
  if err != nil {
      fmt.Printf("Error: %v\n\n", err)
  }

  ctx := context.Background()
  client, err := storage.NewClient(ctx, cloud.WithTokenSource(conf.TokenSource(ctx)))
  if err != nil {
      fmt.Printf("Error: %v\n\n", err)
  }

  for {
    var query *storage.Query
    objects, err := client.Bucket(bucket).List(ctx, query)
    if err != nil {
        fmt.Printf("HERE Error: %v\n\n", err)
        os.Exit(1)
    }

    for _, obj := range objects.Results {
        fmt.Printf("object name: %s, size: %v\n", obj.Name, obj.Size)
    }
    // If there are more results, objects.Next will be non-nil.
    if objects.Next == nil {
        break
    }
    query = objects.Next
  }
/*
  wc := client.Bucket(bucket).Object("test/filename2").NewWriter(ctx)
  wc.ContentType = "text/plain"
  wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
  if _, err := wc.Write([]byte("hello world")); err != nil {
      // TODO: handle error.
  }
  if err := wc.Close(); err != nil {
      // TODO: handle error.
  }
  fmt.Println("updated object:", wc.Attrs())

*/

file, err := os.Open("/vagrant/Vagrantfile") // For read access.
if err != nil {
	fmt.Printf("Error: %v\n\n", err)
}

file_stat, err := file.Stat()
file_size := file_stat.Size()
fmt.Printf("Size: %v\n\n", file_size)

wc := client.Bucket(bucket).Object("test/Vagrantfile").NewWriter(ctx)
wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}

bytes_rem := file_size
for {
  take_bytes := int64(100)
  if (bytes_rem < 100 && bytes_rem > 0) {
    take_bytes = bytes_rem
  }
  bytes_rem = bytes_rem - 100
  fmt.Printf("Bytes: %v\n", take_bytes)

  data := make([]byte, take_bytes)
  count, err := file.Read(data)
  if (err == io.EOF) {
    break
  }

  if err != nil {
  	fmt.Printf("Error: %v\n\n", err)
  }
  if _, err := wc.Write(data); err != nil {
      fmt.Printf("Error: %v\n\n", err)
  }
  fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

if err := wc.Close(); err != nil {
    fmt.Printf("Error: %v\n\n", err)
}

if err := file.Close(); err != nil {
    fmt.Printf("Error: %v\n\n", err)
}


/*
inFile, err := ioutil.ReadFile("/vagrant/Vagrantfile")
if err != nil {
    fmt.Printf("Error: %v\n\n", err)
} else {
  wc := client.Bucket(bucket).Object("test/Vagrantfile").NewWriter(ctx)
  wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
  if _, err := wc.Write(inFile); err != nil {
      fmt.Printf("Error: %v\n\n", err)
  }
  if err := wc.Close(); err != nil {
      fmt.Printf("Error: %v\n\n", err)
  }
}
*/

  // Close the client when finished.
  if err := client.Close(); err != nil {
      fmt.Printf("Error: %v\n\n", err)
  }
}
