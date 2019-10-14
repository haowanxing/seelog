package seelog

import (
	"github.com/hpcloud/tail"
	"log"
	"os"
)

type msg struct {
	LogName string `json:"logName"`
	Data    string `json:"data"`
}

// 监控日志文件
func monitor() {

	for _, sl := range slogs {
		go func(sl slog) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("[seelog] error:%+v", err)
				}
			}()
			log.Println("开始进行日志监控", sl.Name, sl.Path)

			file, err := os.Open(sl.Path)
			if err != nil {
				log.Printf("[seelog} error:%+v", err)
			}
			fileInfo, _ := file.Stat()
			fileSize := fileInfo.Size()
			if err = file.Close(); err != nil {
				log.Print(err)
			}
			t, _ := tail.TailFile(sl.Path, tail.Config{Follow: true, Location: &tail.SeekInfo{
				Offset: fileSize,
				Whence: 0,
			}})
			for line := range t.Lines {
				manager.broadcast <- msg{sl.Name, line.Text}
			}
			log.Printf("tail lines closed!")
		}(sl)
	}

}
