package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// 读取当前目录下的所有文件和文件夹
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		println(entry)
		if entry.IsDir() {
			// 拼接 .git/config 文件的路径
			gitConfigPath := filepath.Join(dir, entry.Name(), ".git", "config")
			if _, err := os.Stat(gitConfigPath); err == nil {
				// 读取 config 文件
				input, err := ioutil.ReadFile(gitConfigPath)
				if err != nil {
					fmt.Println("Error reading config file:", err)
					continue
				}

				// 删除包含指定字符串的行
				output := ""
				scanner := bufio.NewScanner(strings.NewReader(string(input)))
				for scanner.Scan() {
					line := scanner.Text()
					if !strings.Contains(line, `puttykeyfile = D:\\OneDrive\\密钥\\aliyun\\id_rsa.ppk`) {
						output += line + "\n"
					}
				}

				if err := scanner.Err(); err != nil {
					fmt.Println("Error scanning config file:", err)
					continue
				}

				// 写回 config 文件
				err = ioutil.WriteFile(gitConfigPath, []byte(output), 0644)
				if err != nil {
					fmt.Println("Error writing config file:", err)
					continue
				}

				fmt.Println("Updated config file:", gitConfigPath)
			}
		}
	}
}
