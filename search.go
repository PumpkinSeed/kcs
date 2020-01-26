package kcs

import (
	"fmt"
	"os"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/rs/xid"
)

const idxIDDelimiter = "::"

func Search(q string) map[string]map[string]struct{} {
	idx := createIdx()

	query := bleve.NewMatchQuery(q)
	search := bleve.NewSearchRequest(query)
	searchResult, err := idx.Search(search)
	if err != nil {
		panic(err)
	}
	return getResult(searchResult)
}

func getResult(sr *bleve.SearchResult) map[string]map[string]struct{} {
	var result = make(map[string]map[string]struct{})

	for _, hit := range sr.Hits {
		categoryID, commandID := decomposeIdxID(hit.ID)
		if result[categoryID] == nil {
			result[categoryID] = make(map[string]struct{})
		}
		result[categoryID][commandID] = struct{}{}
	}

	return result
}

func composeIdxID(categoryID, commandID string) string {
	return categoryID + idxIDDelimiter + commandID
}

func decomposeIdxID(id string) (string, string) {
	parts := strings.Split(id, idxIDDelimiter)
	if len(parts) != 2 {
		fmt.Println("Invalid bleve id")
		return "", ""
	}

	return parts[0], parts[1]
}

func idxIDGenerator() string {
	idxID := xid.New().String()
	return idxID + "-kcs.bleve"
}

func createIdx() bleve.Index {
	idxID := idxIDGenerator()
	defer func() {
		if err := os.RemoveAll(idxID); err != nil {
			fmt.Println(err)
		}
	} ()

	mapping := bleve.NewIndexMapping()
	idx, err := bleve.New(idxID, mapping)
	if err != nil {
		panic(err)
	}

	for categoryID, category := range Data.Categories {
		for commandID, command := range category.Commands {
			if err := idx.Index(composeIdxID(categoryID, commandID), command); err != nil {
				panic(err)
			}
		}
	}

	return idx
}
