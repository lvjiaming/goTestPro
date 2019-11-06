/**
 日志系统
 */
package customLog

import (
	"errors"
	"fmt"
	"os"
)

type LogWriter interface { // 写入的接口
	Write(data interface{}) error
}

type Logger struct { // 日志器
	writeList []LogWriter
}

func (l *Logger) RegisterWriter (writer LogWriter)  { // 日志器的注册方法
	l.writeList = append(l.writeList, writer)
}

func (l *Logger) Log (data interface{})  { // 日志器的写入方法
	for _, writer := range l.writeList{
		err := writer.Write(data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func NewLogger() *Logger { // 日志器的初始化方法
	return &Logger{}
}

type fileWriter struct {
	file *os.File // 文件
}

func (f *fileWriter) SetFile (name string) (err error) { // 设置写入的文件
	if f.file != nil { // 如果此文件存在，则表示文件已经被打开，需要先关闭
		f.file.Close()
	}
	f.file, err = os.Create(name)
	return err
}

func (f *fileWriter) Write (data interface{}) error {  // 实现日志写入器接口的写入函数
	if f.file == nil {
		return errors.New("file is not create")
	}
	str := fmt.Sprintf("%v\n", data) // 将数据序列化为字符串
	_, err := f.file.Write([]byte(str))  // 将日志写入文件
	return err
}

func NewFileWriter() *fileWriter { // 文件写入器的初始化函数
	return &fileWriter{}
}

type consoleWriter struct {  // 命令行日志器

}

func (c *consoleWriter) Write (date interface{}) error {
	str := fmt.Sprintf("%v\n", date)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

func NewConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}