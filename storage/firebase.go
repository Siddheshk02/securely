package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/Siddheshk02/securely/database"
	"google.golang.org/api/option"
)

// func Upload(path string, data1 []byte, shares int, threshold int) error {
// 	ctx := context.Background()
// 	conf := &firebase.Config{
// 		StorageBucket: "ly-f41b7.appspot.com",
// 	}
// 	opt := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
// 	app, err := firebase.NewApp(ctx, conf, opt)
// 	if err != nil {
// 		log.Fatalf("Failed to initialize Firebase app: %v", err)
// 	}

// 	client, err := app.Storage(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to initialize Firebase Cloud Storage client: %v", err)
// 	}

// 	bucket, err := client.DefaultBucket()
// 	if err != nil {
// 		log.Fatalf("Failed to get default Firebase Cloud Storage bucket: %v", err)
// 	}

// 	var result map[string]interface{}
// 	json.Unmarshal(data1, &result)

// 	name := result["name"].(string)

// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		log.Fatalf("Failed to read file: %v", err)
// 	}

// 	filename := name + " : " + filepath.Base(path) + "/" + filepath.Base(path)
// 	//fmt.Println(path.Base(filepath))

// 	obj := bucket.Object(filename)
// 	wc := obj.NewWriter(ctx)
// 	if _, err = wc.Write(data); err != nil {
// 		log.Fatalf("Failed to write file to Firebase Cloud Storage: %v", err)
// 	}
// 	if err := wc.Close(); err != nil {
// 		log.Fatalf("Failed to close Firebase Cloud Storage writer: %v", err)
// 	}
// 	fmt.Println("File uploaded!")
// 	fmt.Println(" ")

// 	database.FilesDB(name, filepath.Base(path), shares, threshold)

// 	return nil

// }

func Files(info []byte, data []byte, files string, shares int, threshold int, ch int) error {

	ctx := context.Background()
	conf := &firebase.Config{
		StorageBucket: "ly-f41b7.appspot.com",
	}
	opt := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Cloud Storage client: %v", err)
	}

	fmt.Println("test2")
	// Initialize Cloud Storage client
	// client, err := app.Storage(context.Background())
	// if err != nil {
	// 	// Handle error
	// }

	// Create a file in memory
	//data := []byte("Hello, world!")
	var result map[string]interface{}
	json.Unmarshal(data, &result)

	name := result["name"].(string)
	email := result["email"].(string)

	// var file io.ReadCloser
	// if ch == 1 {

	// 	file = ioutil.NopCloser(bytes.NewReader(info))

	// } else if ch == 3 {
	// 	file = ioutil.NopCloser(bytes.NewReader(info))
	// }
	// Upload the file to Cloud Storage
	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalf("Failed to get default Firebase Cloud Storage bucket: %v", err)
	}

	fmt.Println("test3")
	var object *storage.ObjectHandle
	fmt.Println(name, files)

	if ch == 1 {
		filename := name + " : " + files + "/" + "key"
		object = bucket.Object(filename)
	} else if ch == 3 {
		object = bucket.Object(name + " : " + files + "/" + files)
	}
	fmt.Println("test4")
	wc := object.NewWriter(context.Background())
	if _, err = wc.Write(info); err != nil {
		log.Fatalf("Failed to write file to Firebase Cloud Storage: %v", err)
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("Failed to close Firebase Cloud Storage writer: %v", err)
	}
	fmt.Println("test5")

	if ch == 3 {
		fmt.Println("File uploaded!")
		fmt.Println(" ")

		database.FilesDB(name, email, files, shares, threshold)
	}
	return nil
}

func SharesFile(n [][]byte, data []byte, loc string, shares int) {

	ctx := context.Background()
	conf := &firebase.Config{
		StorageBucket: "ly-f41b7.appspot.com",
	}
	opt := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Cloud Storage client: %v", err)
	}

	var result map[string]interface{}
	json.Unmarshal(data, &result)

	name := result["name"].(string)

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalf("Failed to get default Firebase Cloud Storage bucket: %v", err)
	}

	for i := 0; i < shares; i++ {
		filename := name + " : " + loc + "/" + "shares" + "/" + strconv.Itoa(i)
		object := bucket.Object(filename)

		wc := object.NewWriter(context.Background())
		if _, err = wc.Write(n[i]); err != nil {
			log.Fatalf("Failed to write file to Firebase Cloud Storage: %v", err)
		}
		if err := wc.Close(); err != nil {
			log.Fatalf("Failed to close Firebase Cloud Storage writer: %v", err)
		}
	}

	// filename := name + " : " + loc + "/" + "shares"
	// object := bucket.Object(filename)

	// wc := object.NewWriter(context.Background())
	// wc.ContentType = "text/plain; charset=utf-8"
	// wc.ObjectAttrs.ContentType = "text/plain; charset=utf-8"
	// if _, err = wc.Write([]byte("\n")); err != nil {
	// 	log.Fatalf("Failed to write file to Firebase Cloud Storage: %v", err)
	// }
	// if err := wc.Close(); err != nil {
	// 	log.Fatalf("Failed to close Firebase Cloud Storage writer: %v", err)
	// }

	// for i := 0; i < shares; i++ {
	// 	wc = object.NewWriter(context.Background())
	// 	wc.ContentType = "text/plain; charset=utf-8"
	// 	wc.ObjectAttrs.ContentType = "text/plain; charset=utf-8"
	// 	if _, err = wc.Write(n[i]); err != nil {
	// 		log.Fatalf("Failed to write file to Firebase Cloud Storage: %v", err)
	// 	}
	// 	if _, err = wc.Write([]byte("\n")); err != nil {
	// 		log.Fatalf("Failed to write file to Firebase Cloud Storage: %v", err)
	// 	}

	// 	if err := wc.Close(); err != nil {
	// 		log.Fatalf("Failed to close Firebase Cloud Storage writer: %v", err)
	// 	}
	//}
	// fmt.Println("File uploaded!")
	// fmt.Println(" ")
}

func ReadShares(file string, admin string, shares int) ([][]byte, []byte, []byte, error) {

	ctx := context.Background()
	conf := &firebase.Config{
		StorageBucket: "ly-f41b7.appspot.com",
	}
	opt := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Cloud Storage client: %v", err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalf("Failed to get default Firebase Cloud Storage bucket: %v", err)
	}

	// filename := admin + " : " + file + "/" + "shares"
	// object := bucket.Object(filename)

	// reader, err := object.NewReader(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	// defer reader.Close()
	// contents := make([]byte, 53)
	//var contents []byte
	//n := make([][]byte, shares, 53)

	filename := admin + " : " + file + "/" + file
	object := bucket.Object(filename)

	reader, err := object.NewReader(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	defer reader.Close()

	data1, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error reading data: %v\n", err)
	}

	filename = admin + " : " + file + "/" + "key"
	object = bucket.Object(filename)

	reader, err = object.NewReader(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error reading data: %v\n", err)
	}

	var n [][]byte

	for i := 0; i < shares; i++ {
		filename = admin + " : " + file + "/" + "shares" + "/" + strconv.Itoa(i)
		object := bucket.Object(filename)

		reader, err := object.NewReader(context.Background())
		if err != nil {
			return nil, nil, nil, err
		}
		defer reader.Close()

		for {
			buf := make([]byte, 1024)
			//fmt.Println(buf)
			d, err := reader.Read(buf)
			//fmt.Println(d)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, nil, nil, err
			}
			n = append(n, buf[:d])
			// fmt.Println(contents)
		}

		//n = append(n, contents)

	}
	// fmt.Println(n)

	// var n [][]byte
	// lines := strings.Split(string(contents), "\n")
	// for _, line := range lines {
	// 	if len(line) > 0 {
	// 		n = append(n, []byte(line))
	// 	}
	// }
	return n, data, data1, nil

}

// func Download(filepath) error {

// 	return nil
// }
