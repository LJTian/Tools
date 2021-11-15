package tools

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// LocalTime 获取时间
func LocalTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

// GetCurrentDirectory 获取执行目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}

// OsExe 执行系统命令函数
func OsExe(CmdName, action, SvrName string) (err error) {

	cmd := exec.Command(CmdName, action, SvrName)
	Cmderr := cmd.Run()
	if err != nil {
		err = Cmderr
	}
	return err
}

//CreateUUID 生成UUID
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// 报文长度计算
func StatisticalLen(msg string, MsgLenNum int) []byte {

	var MsgLen [2]byte
	var msgInfo bytes.Buffer

	// 添加报文头
	iLen := len(hex.EncodeToString([]byte(msg))) / 2
	MsgLen[0] = byte(iLen >> 8)
	MsgLen[1] = byte(iLen & 0xff)

	msgInfo.WriteByte(MsgLen[0])
	msgInfo.WriteByte(MsgLen[1])
	msgInfo.WriteString(msg)

	return msgInfo.Bytes()
}
