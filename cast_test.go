package cast

import (
	"crypto/rand"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// getHeap returns the size of the heap
func getHeap(seed int) int {
	runtime.GC()
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	return int(ms.HeapAlloc) - seed
}

// TestCast will create a 100KB []byte and assign the first 50KB
// to a string using ToString() and the copy the second 50KB to a
// string using string(). Both will be assigned to a map[string]bool
// and should result in around 150KB of heap. 100KB for the original
// []byte and 50 for the string() copy.
// We will test read the heap alloc to see if it's around 150KB.
// Then we'll nil the map.
// Then wait up to 10 seconds for the memory to get near zero.
func TestToString(t *testing.T) {
	const sz = 1024 * 500
	var m [2]string
	ch := make(chan bool)
	var start time.Time
	seed := getHeap(0)
	go func() {
		b := make([]byte, sz)
		rand.Read(b)
		m[0] = ToString(b[len(b)/2:])
		m[1] = string(b[:len(b)/2])
		ch <- true
	}()
	<-ch
	if 1.0-float64(getHeap(seed))/(sz+sz/2.0) > 0.05 {
		t.Fatal("failed")
	}
	m[0], m[1] = "", ""
	start = time.Now()
	for {
		if time.Now().Sub(start) > time.Second*10 {
			t.Fatal("failed")
		}
		per := 1.0 - float64(getHeap(seed))/(sz+sz/2.0)
		if per > 0.95 {
			break
		}
	}
}

// TestToBytes is the same as TestToString, but the other way around.
func TestToBytes(t *testing.T) {
	const sz = 1024 * 500
	var m [2][]byte
	ch := make(chan bool)
	var start time.Time
	seed := getHeap(0)
	go func() {
		b := make([]byte, sz)
		rand.Read(b)
		s := string(b)
		b = nil
		m[0] = ToBytes(s[len(s)/2:])
		m[1] = []byte(s[:len(s)/2])
		ch <- true
	}()
	<-ch
	if 1.0-float64(getHeap(seed))/(sz+sz/2.0) > 0.05 {
		t.Fatal("failed")
	}
	m[0], m[1] = nil, nil
	start = time.Now()
	for {
		if time.Now().Sub(start) > time.Second*10 {
			t.Fatal("failed")
		}
		per := 1.0 - float64(getHeap(seed))/(sz+sz/2.0)
		if per > 0.95 {
			break
		}
	}
}

func ExampleToBytes() {
	var s = "Hello Planet"
	b := ToBytes(s)
	fmt.Printf("%s\n", "J"+string(b[1:]))

	// Output:
	// Jello Planet
}

func ExampleToString() {
	var b = []byte("Hello Planet")
	s := ToString(b)
	b[0] = 'J'
	fmt.Printf("%s\n", s)

	// Output:
	// Jello Planet
}
