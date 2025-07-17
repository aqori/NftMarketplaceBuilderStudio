// cmd/nftmarketplacebuilderstudio/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacebuilderstudio/internal/nftmarketplacebuilderstudio"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacebuilderstudio.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
