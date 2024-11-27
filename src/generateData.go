package src

import (
	"fmt"
	"log"
	"math/rand/v2"
	"github.com/bxcodec/faker/v4"
)

func SeedData(numRecords int) {
	if db == nil {
		log.Fatal("Database not initialized, put InitDb() first.")
	}

	query := `INSERT INTO Products (name, description, price, image_url ) VALUES (?,?,?,?)`

	for i := 0; i < numRecords; i++ {
		name := faker.Name()
		description := faker.Sentence()
		price := rand.Float32()*100 + 1
		imageUrl := fmt.Sprintf("https://picsum.photos/seed/%d/200/300", i)
		_, err := db.Exec(query, name, description, price, imageUrl)
		if err != nil {
			log.Fatalf("Failed to insert record: %v", err)
		}
	}
	fmt.Printf("%d records inserted into Products table. \n", numRecords)
}

