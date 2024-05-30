package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

//Version 4
func NewUUID() string {
	return uuid.NewString()
}

//Version 1
func NewV1UUID() string {
	uuidV1, _ := uuid.NewUUID()
	return uuidV1.String()
}

//Version 3
func NewV3UUID(name string) string {
	uuidV3 := uuid.NewMD5(uuid.NameSpaceDNS, []byte(name))
	return uuidV3.String()
}

//Version 5
func NewV5UUID(name string) string {
	uuidV5 := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(name))
	return uuidV5.String()
}

//生成唯一ID
func GenId(mode, code int) string {
	var builder strings.Builder
	now := time.Now()
	// mode
	modeStr := strconv.Itoa(mode % 10)
	if ln := len(modeStr); ln < 2 {
		modeStr = strings.Repeat("0", 2-ln)
	}
	builder.WriteString(modeStr)
	// code
	if code == 0 {
		code = int(now.UnixNano())
	}
	codeStr := strconv.Itoa(code % 10000)
	if ln := len(codeStr); ln < 4 {
		codeStr = strings.Repeat("0", 4-ln)
	}
	builder.WriteString(codeStr)
	// date
	builder.WriteString(now.Format("060102150405"))
	// rand
	rand.Seed(now.UnixNano())
	randNum := rand.Intn(9999-1000+1) + 1000
	builder.WriteString(strconv.Itoa(randNum))

	return builder.String()
}
