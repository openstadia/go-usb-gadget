package gadget

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetUdcs() []string {
	var udcs []string

	files, err := os.ReadDir("/sys/class/udc")
	if err != nil {
		return nil
	}

	for _, file := range files {
		udcs = append(udcs, file.Name())
	}

	return udcs
}

func WriteBuf(path string, name string, file string, buf []byte) {
	fullPath := filepath.Join(path, name, file)
	fp, err := os.OpenFile(fullPath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}

	defer fp.Close()

	_, err = fp.Write(buf)
	if err != nil {
		return
	}
}

func WriteString(path string, name string, file string, str string) {
	WriteBuf(path, name, file, []byte(str))
}

func WriteInt(path string, name string, file string, value int, format string) {
	buf := fmt.Sprintf(format, value)
	WriteString(path, name, file, buf)
}

func WriteDec(path string, name string, file string, value int) {
	WriteInt(path, name, file, value, "%d\n")
}

func WriteHex16(path string, name string, file string, value uint16) {
	WriteInt(path, name, file, int(value), "0x%04x\n")
}

func WriteHex8(path string, name string, file string, value uint8) {
	WriteInt(path, name, file, int(value), "0x%02x\n")
}
