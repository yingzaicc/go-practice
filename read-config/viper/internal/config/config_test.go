package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	InitConfig()
	fmt.Printf("Database config: %v\n", GetConfig())
}
