package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Helper function execute commands on the commandline.
func shellOut(c []string) error {
	return exec.Command("sh", "-c", strings.Join(c, " ")).Run()
}

// Exists the application.
func fatal(e error) {
	b, _ := json.Marshal(map[string]string{
		"status":  "Failure",
		"message": e.Error(),
	})
	fmt.Println(string(b))
	os.Exit(2)
}
