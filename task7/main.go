/* Что такое интерфейс и как он реализуется в Go?

/ Применяют по месту использования, а не объявления.
Возможна композиция интерфейса и реализация через утиную типизацию.

*/

package task7

import "fmt"

type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error) 
}

type ReadWriter interface {
    Reader
    Writer
}

type File struct{}

func (f File) Write(data []byte) (int, error) {
    fmt.Printf("Writing: %s\n", data)
    return len(data), nil
}

func (f File) Read(data []byte) (int, error) {
    copy(data, "data from file")  
    return len("data from file"), nil
}

func save(w Writer) {
    w.Write([]byte("data to save"))
}

func read(r Reader) {
    data := make([]byte, 50)
    n, _ := r.Read(data)
    fmt.Printf("Read: %s (%d bytes)\n", data[:n], n)
}

func main() {
    f := File{}
    
    save(f)
    read(f)
}