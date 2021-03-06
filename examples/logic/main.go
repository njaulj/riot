package main

import (
	"log"

	"github.com/go-ego/riot/engine"
	"github.com/go-ego/riot/types"
)

var (
	// searcher is coroutine safe
	searcher = engine.Engine{}
)

func main() {
	// Init engine
	searcher.Init(types.EngineInitOptions{
		Using:             5,
		NotUsingSegmenter: true})
	defer searcher.Close()

	// Add the document to the index, docId starts at 1
	searcher.IndexDocument(1, types.DocIndexData{Content: "Google Is Experimenting With Virtual Reality Advertising"}, false)
	searcher.IndexDocument(2, types.DocIndexData{Content: "Google accidentally pushed Bluetooth update for Home speaker early"}, false)
	searcher.IndexDocument(3, types.DocIndexData{Content: "Google is testing another Search results layout with rounded cards, new colors, and the 4 mysterious colored dots again"}, false)

	// Wait for the index to refresh
	searcher.FlushIndex()

	// var strArr []string
	strArr := []string{"accidentally"}
	query := "google testing"

	// The search output format is found in the types.SearchResponse structure
	docs := searcher.Search(types.SearchRequest{
		Text: query,
		Logic: types.Logic{
			ShouldLabels: true,
			LogicExpression: types.LogicExpression{
				NotInLabels: strArr,
			},
		},
		RankOptions: &types.RankOptions{
			OutputOffset: 0,
			MaxOutputs:   100,
		}})

	log.Println("search...", docs)
}
