package config

import "os"

func GitPushUser() string {
	return os.Getenv("GIT_PUSH_USERNAME")
}
func GitPushUserPwd() string {
	return os.Getenv("GIT_PUASH_PASSWORD")
}
