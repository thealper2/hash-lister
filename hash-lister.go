package main

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func printUsage() {
    fmt.Printf(`
Usage:
    ` + os.Args[0] + ` --md5
    ` + os.Args[0] + ` --sha1
    ` + os.Args[0] + ` --sha256
    ` + os.Args[0] + ` --sha512

Example:
    # Print MD5 checksum of a file.
    ` + os.Args[0] + ` --md5 <filename>

    # Print SHA1 checksum of a file.
    ` + os.Args[0] + ` --sha1 <filename>

    # Print SHA256 checksum of a file.
    ` + os.Args[0] + ` --sha256 <filename>

    # Print SHA512 checksum of a file.
    ` + os.Args[0] + ` --sha512 <filename>

`)
}

func checkArgs() (string, string) {
    if len(os.Args) != 3 {
        printUsage()
        os.Exit(1)
    }
    return os.Args[1], os.Args[2]
}

func main() {
    filename, hashFlag := checkArgs()

    file, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    switch hashFlag {
    case "--md5": fmt.Printf("MD5: %x\n", md5.Sum(file))
    case "--sha1": fmt.Printf("SHA1: %x\n", sha1.Sum(file))
    case "--sha256": fmt.Printf("SHA256: %x\n", sha256.Sum256(file))
    case "--sha512": fmt.Printf("SHA512: %x\n", sha512.Sum512(file))
    }

}
