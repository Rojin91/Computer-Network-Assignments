package main

import (
	"fmt"
	"log"
	"os"
	"BitTorrentClient/torrent"
	"BitTorrentClient/utils"
)

func main() {
	utils.InitLogger()
	if len(os.Args) < 2 {
		log.Println("Error: No path to torrent file provided.")
		fmt.Println("Usage: BitTorrentClient <path-to-torrent-file>")
		os.Exit(1)
	}

	torrentPath := os.Args[1]

	// Open and parse the .torrent file
	torrentFile, err := torrent.Open(torrentPath)
	if err != nil {
		log.Fatalf("Failed to open torrent file: %v", err)
	}

	// Start the download process
	err = torrentFile.Download("output.dat")
	if err != nil {
		log.Fatalf("Failed to download torrent: %v", err)
	}
}
