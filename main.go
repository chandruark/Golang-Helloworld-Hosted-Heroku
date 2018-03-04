package main
import (
  "log"
  "fmt"
  "net/http"
  "os"
)
func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8000"
    return port, fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}
func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello  bala how are you World")
}
func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }
  http.HandleFunc("/", hello)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}