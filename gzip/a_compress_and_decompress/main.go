package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

var originJsonData = `
{
    "apm_id": "20000001",
    "header":
    {
        "custom":
        {
            "sdk_version": "6.4.0.1",
            "host_app_id": "8025677"
        },
        "os": "Android",
        "os_version": "12",
        "device_model": "SM-A217F",
        "device_brand": "samsung",
        "sdk_version_name": "0.0.5",
        "aid": "10000000",
        "update_version_code": 6401,
        "bd_did": "63186367869893"
    },
    "local_time": 1735628334587,
    "launch":
    [
        {
            "local_time_ms": 1735628334587
        }
    ]
}
`

// 压缩
func compressJSONToGzip(jsonData []byte) ([]byte, error) {
	var buf bytes.Buffer

	// 使用最佳压缩级别（可选）
	gz, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	if _, err := gz.Write(jsonData); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 解压
func decompressGzipToJSON(gzipData []byte) ([]byte, error) {
	// 1. 创建一个 bytes.Reader 读取压缩数据
	reader := bytes.NewReader(gzipData)
	// 2. 创建 Gzip 读取器
	gz, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	// 3. 读取解压后的数据
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, gz); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	// 压缩
	originJsonDataByte := []byte(originJsonData)
	compressedData, err := compressJSONToGzip(originJsonDataByte)
	if err != nil {
		panic(err)
	}

	// 压缩前后的len对比
	fmt.Printf("Original: %d bytes\n", len(originJsonData))
	fmt.Printf("Compressed: %d bytes\n", len(compressedData))

	// 解压
	decompressedBytes, err := decompressGzipToJSON(compressedData)
	if err != nil {
		panic(err)
	}
	// 解压数据的长度和内容.
	fmt.Printf("decompressed Data: %d bytes\n", len(decompressedBytes))
	fmt.Printf("Original Date : %s", string(decompressedBytes))
}
