package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchFileToHttp(filePath string, w http.ResponseWriter) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error when open file " + err.Error())
		return err
	}
	defer file.Close()
	totalBytes := 0
	BufferSize := 1024
	buffer := make([]byte, BufferSize)

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		totalBytes += bytesRead
		fmt.Printf("bytes read: %d\r", totalBytes)
		//fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
		_, err = w.Write(buffer)
		if err != nil {
			fmt.Println("Error when write to http.ResponseWriter: "+err.Error())
			return err
		}
	}
	fmt.Printf("bytes read: %d\n", totalBytes)
	return nil
}
