package main

import (
    "fmt"
    "log"
    "os"
    "TorrentClient/torrent"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: TorrentClient <path-to-torrent-file>")
        os.Exit(1)
    }

    torrentFilePath := os.Args[1]
    tFile, err := torrent.LoadTorrentFile(torrentFilePath)
    if err != nil {
        log.Fatalf("Failed to load torrent file: %v", err)
    }

    outputFilePath := "downloaded_file" // Set your output file path here
    err = tFile.Download(outputFilePath)
    if err != nil {
        log.Fatalf("Download failed: %v", err)
    }

    fmt.Println("Download completed successfully.")
}