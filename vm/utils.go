package vm

import (
	"bytes"
	"crypto/md5"
	"encoding/base32"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

func GetMacAddr(id bson.ObjectId, index int) string {
	hash := md5.New()
	hash.Write([]byte(id.Hex() + strconv.Itoa(index)))
	macHash := fmt.Sprintf("%x", hash.Sum(nil))
	macHash = macHash[:10]
	macBuf := bytes.Buffer{}

	for i, run := range macHash {
		macBuf.WriteRune(run)
		if i%2 == 1 && i != len(macHash)-1 {
			macBuf.WriteRune(':')
		}
	}

	return "00:" + macBuf.String()
}

func GetIface(id bson.ObjectId, n int) string {
	hash := md5.New()
	hash.Write([]byte(id.Hex()))
	hashSum := base32.StdEncoding.EncodeToString(hash.Sum(nil))[:12]
	return fmt.Sprintf("p%s%d", strings.ToLower(hashSum), n)
}
