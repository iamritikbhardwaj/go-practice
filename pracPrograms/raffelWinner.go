package pracPrograms

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type RaffelEntry struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RaffelEnties []RaffelEntry

func importData() RaffelEnties {
	data, err := os.ReadFile("entries.json")
	if err != nil {
		panic(err)
	}

	var entries RaffelEnties
	json.Unmarshal(data, &entries)
	return entries
}

func winner(entries RaffelEnties) RaffelEntry {
	wi := rand.Intn(len(entries))

	return entries[wi]
}

func (r *RaffelEntry) RaffelWinner() {
	entries := importData()
	winner := winner(entries)
	fmt.Println("RaffelWinner", winner.Name)
}
