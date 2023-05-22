package main

import (
	"fmt"
	"log"

	"github.com/EtoniaThea/googledocs"
)

func main() {
	client, err := googledocs.NewClient("credentials.json", "token.json")
	if err != nil {
		log.Fatalf("Unable to create Google Docs client: %v", err)
	}

	// Create a document
	documentID, err := client.CreateDocument("My New Document")
	if err != nil {
		log.Fatalf("Unable to create document: %v", err)
	}
	fmt.Printf("Created document with ID: %s\n", documentID)

	// Read a document
	doc, err := client.ReadDocument(documentID)
	if err != nil {
		log.Fatalf("Unable to read document: %v", err)
	}
	fmt.Printf("The title of the doc is: %s\n", doc.Title)

	// Update a document
	err = client.UpdateDocument(documentID, "Updated Document")
	if err != nil {
		log.Fatalf("Unable to update document: %v", err)
	}
	fmt.Println("Document updated successfully.")

	// Delete a document
	err = client.DeleteDocument(documentID)
	if err != nil {
		log.Fatalf("Unable to delete document: %v", err)
	}
	fmt.Println("Document deleted successfully.")
}
