package firestore

import (
	"context"
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
	Author   string
	Address  string
	Email    string
	Solution string
}

func GetFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(`{
		"type": "service_account",
		"project_id": "someapiproject",
		"private_key_id": "8d98f52858030672165531ac14bd80498bee92e8",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC0XGlPz3D0H7n6\ncmNfOGfEkqtbgbUex7JISzL3rkLJRQe0uROI/EX6CUK6ksl7Emrq1kY5t/xGx6k+\nxYGWics704k15a1/AL5X/WRRaGxtkQhuQazqCwHqjl6xGuO6klTdvFCZPAsEbxNH\nGGY5NPCi4VnNE1F7R8hfr48KOQ6FXWBM2hqZRODMe6xJP7GA0hH95V2kW5QDveJs\nMU9Hd9O9rO1uhD58o0jAk9mbIF268xrpjOEYqxvTVDZb7iZWI8m12fu3eV0SbfIF\nKOLjf1/NCfZyQ7OttdkqT30I6oBHUo4fTPNVkBaj69Sj9K6qDqE9GL6gXdZkJZ4n\njR/LEklDAgMBAAECggEAR3RFFKVubokTOTGQRO0nzyz8Tmh4xRUAgLuqGY4kT3DV\nuLeKEb9ASerZUOlOgT+utBLoB33oqHH5jzDYQjedGLqZpYy0y5gT0PBGiioAqvfG\ni0fhpWdu/uoggbHRftzyWlZ85/httPf8fzIfbZKXsy/sT97TbS/nJmF7HeW05wiI\nEq23byOm1PXj4nmJLZyVgm6wccdk5PoCQ2UAjfAQUlv74u39JRQzA8bRNwJKNtTD\ndxctmjYJBmDTsRpfgSHwWa/AXlJ9YjuPLLgWc7AU5UEO5a3x7EtHjLRYmjpINMQO\nbpd464IX1HN+DDcIMPyCanyh+xYmjxrd88pTVHzLYQKBgQDvuNzxGC9EjqzwnAgE\nckMmzAojW5vCi/6GjCpJCEwe5x+ar7D6g2SNbawvH1qdyol4YaPctkMrasiHSEgl\ncr9cvMuPXNQ3W9yj436yTmv5+UcKSAv1vJzANVW1cMCsr68O0G81+GDfCnm1Rpb+\nclOzK98Ss0qyV4/G72YxfOvivwKBgQDAm6lzhfHHsopgaSGiBbTeMINycOZnIs5k\nibPEycEKW53xxmslLkcNb8GjAu+A+6mzc6be50fDpAqKGmlhH/TqlyrMmtLwbHO2\nP5Enb0SctqLRR6GPsmrRLX4+PdvKsD8Dmb4TJ4RzrgAUPuiCqktSpa8T2efwSuKe\nUmBc8anufQKBgQCAz1WxGuyzKvMUatMICJm6rCK6cwwUEoNWqtRB3/p/FHPv+33e\nbmHGeOrveyqG5QDPNbAF9c3L85oCzz0tGiZnX28F/rxtbqf1TFWU2/y7Gk4o4SPE\nDHAx+7atQwPVBqXLEQbg+jCbSJazaFXULXx6JxW7h6mYgOJZ4+OGrfhWIwKBgQCc\n36vaPaQ5ZD+0WqxcDI3N0nGdSjs+kWjNFiLnCvRBfXFdNKCb/d89IGL0ZDWyNkd3\ns6CcOH+I5xj2dqCRzLdsQodHcmqQC6ULMScGmWemxFJEZjU+lrDNgmIqS7OymG4a\nfqQDcdI9beD+nGY/1nfW7r90SazRWAzPqoR17xba+QKBgAt91dbBe8pVXorRIykx\nSacqCBrUxyiDo2csOOiWAlrDTJyK2ccLMjXBOeNVQXleRN3Ijg5qvV4mz8bYkskA\nMHARfqPgrw7+nM+EBJcarwFl1AxmuJvgkbsoMRlMzSj0kFiE/VVx3bmoXFeph9Ow\nREnGAY3X11tCz7nw+sA+mHQi\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-2bnp3@someapiproject.iam.gserviceaccount.com",
		"client_id": "113246810133394875986",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-2bnp3%40someapiproject.iam.gserviceaccount.com"
	  }
	  `))
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
