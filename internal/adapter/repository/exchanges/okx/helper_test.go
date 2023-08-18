package okx

import (
	"hamgit.ir/novin-backend/trader-bot/config"
	"log"
	"testing"
)

func Test_createOKXSignature(t *testing.T) {
	config.Init("../../../../../dev/config/trader/")
	log.Println(createOKXSignature("", "", config.C().OKX.SecretKey, nil))
}
