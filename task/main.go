package main

import (
    "fmt"
    "io"
    "net"
    "os"
)

func main() {
   
    conn, err := net.Dial("tcp", "localhost:80")
    if err != nil {
        fmt.Println("serverda ulanish xatolike buldii", err)
        return
    }
    defer conn.Close()

   
    files := []string{"newfile.txt"} 
    for _, file := range files {
        err = sendFile(conn, file)
        if err != nil {
            fmt.Println("faylni  xatolik", err)
            return
        }
    }
    fmt.Println("baarcha fayllar muvafaqiyatli yuborildi.")
}

func sendFile(conn net.Conn, filename string) error {
  
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    
    _, err = fmt.Fprintf(conn, "%s\n", filename)
    if err != nil {
        return err
    }


    _, err = io.Copy(conn, file)
    if err != nil {
        return err
    }
    fmt.Printf("%s fayli muvaffaqiyatli yuborildi.\n", filename)
    return nil
}
