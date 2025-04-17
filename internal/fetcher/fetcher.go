package fetcher

import (
    "net/http"
    "log"
)

func GetHTML(url string) *http.Response {
    res, err := http.Get(url)
    if err != nil {
        log.Fatalf("Ошибка при загрузке: %v", err)
    }

    if res.StatusCode != 200 {
        log.Fatalf("Статус код: %d", res.StatusCode)
    }

    return res
}