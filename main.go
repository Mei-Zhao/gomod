package main

import (
    "crypto/rand"
    "fmt"
    "io"
    "io/ioutil"

    "github.com/cheggaaa/pb/v3"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
    var limit int64 = 1024 * 1024 * 500
    // we will copy 200 Mb from /dev/rand to /dev/null
    reader := io.LimitReader(rand.Reader, limit)
    writer := ioutil.Discard

    // start new bar
    bar := pb.Full.Start64(limit)
    // create proxy reader
    barReader := bar.NewProxyReader(reader)
    // copy from proxy reader
    io.Copy(writer, barReader)
    // finish bar
    bar.Finish()

}
