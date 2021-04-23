package tool

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tedcy/fdfs_client"
)

//上传文件到fastDFS分布式文件存储系统工具包
func UploadFileToDFS(fileName string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	defer client.Destory()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fileId, err := client.UploadByFilename(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return fileId
}

func FileServerAddr() string {
	file, err := os.Open("./config/fastdfs.go")
	if err != nil {
		return ""
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return ""
		}

		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2) //通过=切割字符串
		switch str[0] {
		case "http_server_port":
			return str[1]
		}

	}
}
