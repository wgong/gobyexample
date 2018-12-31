// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.

package main

import "fmt"
import "time"

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
    fmt.Println("working...")
    time.Sleep(2 * time.Second)
    fmt.Println("done after 2 sec")

    // Send a value to notify that we're done.
    done <- true
}

func main() {

    // Start a worker goroutine, giving it the channel to
    // notify on.
    done := make(chan bool, 1)
    fmt.Println(time.Now())
    go worker(done)
    fmt.Println(time.Now())

    // Block until we receive a notification from the
    // worker on the channel.
    <-done
    fmt.Println(time.Now())
}
