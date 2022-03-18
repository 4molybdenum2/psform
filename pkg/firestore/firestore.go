package firestore

import (
	"context"
	"io/ioutil"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	projectID = "someapiproject"
)

// Define Response type
type Response struct {
	Author   string `json:"author"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Solution string `json:"solution"`
}

func GetFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	data, err := ioutil.ReadFile("serviceAccountFirebase.json")
	if err != nil {
		log.Fatalf("failed to parse config file : %v", err)
		return nil, err
	}
	opt := option.WithCredentialsJSON(data)
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("failed to create firestore client : %v", err)
		return nil, err
	}
	return client, nil
}

func GetResponse() (*[]Response, error) {
	ctx := context.Background()
	client, err := GetFirestoreClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	var responses []Response
	docIterator := client.Collection("responses_collection").Documents(ctx)
	for {
		docSnapshot, err := docIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to retrieve list of responses : %v", err)
			return nil, err
		}
		response := Response{
			Author: docSnapshot.Data()["author"].(string),
		}
		responses = append(responses, response)
	}

	return &responses, nil
}

func CreateResponse(response *Response) (*Response, error) {
	ctx := context.Background()
	client, err := GetFirestoreClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection("responses_collection").Add(ctx, map[string]interface{}{
		"author":   response.Author,
		"address":  response.Address,
		"email":    response.Email,
		"Solution": response.Solution,
	})
	if err != nil {
		log.Fatalf("failed to create new response : %v", err)
		return nil, err
	}
	return response, nil
}
