package utils

import (
    "fmt"
    "math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func GetRandomUsername() string {
    return randSeq(8)
}

func GetFullHTTPUrl(host string, uri string, isSecure bool) string {
    scheme := "http"
    if isSecure {
        scheme = "https"
    }
    return fmt.Sprintf("%s://%s/%s", scheme, host, uri)
}
