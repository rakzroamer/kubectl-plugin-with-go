package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	repoName := "rakzroamer"
	remoteURL, err := getRemoteURL()
	if err != nil {
		fmt.Printf("Unable to get the remote URL at %s", err)
		return
	}

	fmt.Println("Remote URL: ", remoteURL)

	// if !strings.Contains(remoteURL, "origin") {
	// 	fmt.Println("Remote origin not set. Exiting")
	// 	return
	// }

	if err := gitAdd(); err != nil {
		fmt.Printf("Error in Git add: %s\n", err)
		return
	}
	if err := gitCommit(); err != nil {
		fmt.Printf("Error in Git Commit: %s\n", err)
		return
	}
	if err := gitPush(remoteURL); err != nil {
		fmt.Printf("Error in Git Push: %s\n", err)
		return
	}
	fmt.Printf("Git automation is complete, refer your repo at %s\n", repoName)
}
func getRemoteURL() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
func gitAdd() error {
	gitAddCmd := exec.Command("git", "add", ".")
	return gitAddCmd.Run()
}

func gitCommit() error {
	timestamp := time.Now().Format("2006-01-02 01:01:01")
	commitMsg := fmt.Sprintf("changes at %s", timestamp)
	gitCommitCmd := exec.Command("git", "commit", "-m", commitMsg)
	return gitCommitCmd.Run()
}

func gitPush(remoteURL string) error {
	gitPushCmd := exec.Command("git", "push", remoteURL, "main")
	return gitPushCmd.Run()
}
