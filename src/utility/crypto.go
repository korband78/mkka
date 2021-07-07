package utility

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// UUID : RFC4122
func UUID(isDash bool) string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	if isDash {
		return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	}
	return fmt.Sprintf("%x%x%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// MD5FromFilepath : 파일 경로 => MD5 해시
func MD5FromFilepath(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	return MD5FromFile(file)
}

// MD5FromFile : 파일 객체 => MD5 해시
func MD5FromFile(file *os.File) (string, error) {
	var returnMD5String string
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

// MD5FromString : 스트링 => MD5 해시
func MD5FromString(text string) (string, error) {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// MD5FromBytes : 바이트 => MD5 해시
func MD5FromBytes(byt []byte) (string, error) {
	hash := md5.New()
	hash.Write(byt)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
