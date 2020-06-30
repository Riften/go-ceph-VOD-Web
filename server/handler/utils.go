package handler

import (
	"fmt"
	"github.com/ceph/go-ceph/rados"
	"io"
	"net/http"
	"os"
)

func fetchCephToHttp(conn *rados.Conn, pool string, objectName string, w http.ResponseWriter) error {
	ioctx, err := conn.OpenIOContext(pool)
	if err != nil {
		fmt.Println("Error when open ceph pool " ,pool, ": ", err)
		return err
	}
	var totalReadBytes uint64
	totalReadBytes = 0
	bytesOut := make([]byte, 1024*1024*10) //10MB buffer
	for {
		bytesRead, err := ioctx.Read(objectName, bytesOut, totalReadBytes)
		if err != nil {
			fmt.Println("Error when read obj from ceph pool: ", err)
			break
		}
		if bytesRead == 0 {
			fmt.Println("Read end")
			break
		}
		_, err = w.Write(bytesOut[:bytesRead])
		if err != nil {
			fmt.Println( "Error when write bytes to http: ", err)
			return err
		}
	}
	return nil
}

func fetchFileToCeph(filePath string, conn *rados.Conn, pool string, objectName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error when open file " + err.Error())
		return err
	}

	defer file.Close()

	ioctx, err := conn.OpenIOContext(pool)
	if err != nil {
		fmt.Println("Error when open ceph pool ", pool, ": ", err)
		return err
	}
	// write some data

	totalBytes := 0
	BufferSize := 1024*1024*10 //10MB
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
		if bytesRead < len(buffer) {
			buffer = buffer[:bytesRead]
		}
		err = ioctx.Append(objectName, buffer)
		if err != nil {
			fmt.Println("Error when write to ceph: ", err)
			return err
		}
	}
	fmt.Println(filePath, " is written to ", objectName)
	return nil
}

func fetchFileToHttp(filePath string, w http.ResponseWriter) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error when open file " + err.Error())
		return err
	}
	defer file.Close()
	totalBytes := 0
	BufferSize := 1024*1024*10 //10MB
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

func directoryExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func fileExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

func secondsToString(seconds int64) string {
	var sPerMinute int64
	var sPerHour int64
	var hour int64
	var minute int64
	sPerMinute = 60
	sPerHour = sPerMinute*60
	hour = seconds / sPerHour
	minute = (seconds % sPerHour) / sPerMinute
	seconds = seconds - hour*sPerHour - minute * sPerMinute
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, seconds)
}
