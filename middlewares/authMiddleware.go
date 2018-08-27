package middlewares

import (
  "net/http"
  "os"
  "fmt"
  "bytes"
  "io/ioutil"
)


func AuthMiddleware(next http.Handler) http.Handler {

  AUTH_PROXY := os.Getenv("AUTH_PROXY")
  AUTH_PORT := os.Getenv("AUTH_PORT")
  AUTH_ENDPOINT := os.Getenv("AUTH_ENDPOINT")

  url := fmt.Sprintf("%s://%s:%s/%s", "https", AUTH_PROXY,  AUTH_PORT, AUTH_ENDPOINT)

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    r.Body = ioutil.NopCloser(bytes.NewReader(body))

    validationRequest, err := http.NewRequest("GET", url, bytes.NewReader(body))

    validationRequest.Header.Set("Host", r.Host)
    validationRequest.Header.Set("X-Forwarded-For", r.RemoteAddr)

    for header, values := range r.Header {
        for _, value := range values {
            validationRequest.Header.Add(header, value)
        }
    }

    client := &http.Client{}
    _, validationError := client.Do(validationRequest)
    if validationError != nil {
      http.Error(w, validationError.Error(), http.StatusForbidden)
    }

    if err != nil {
      next.ServeHTTP(w, r)

    } else {
      next.ServeHTTP(w, r)
    }
  })
}
