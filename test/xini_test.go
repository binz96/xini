package xini_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/binz96/xini"
)

func TestLoad(t *testing.T) {
	cfg, err := xini.Load("my.ini")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*cfg)
	v := cfg.Section("").Key("k")
	v1 := cfg.Section("").Key("k1")
	v2 := cfg.Section("server").Key("k2")
	v3 := cfg.Section("server").Key("k3")
	fmt.Println(v, v1, v2, v3)
}
