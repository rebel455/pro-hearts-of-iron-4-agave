package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.5" )

func qZ6pPSklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bd3fsd4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ouRn86UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6ycVUfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDbZbAPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvKKZ5h9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFizYt9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Munam7BOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6u2dogOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmWuAmV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLp4Dj2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwBsXgUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7IEgYl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aig43040Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eL0P91pMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dDZ2YqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gznAEtr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2qlFgmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yFjjvM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoj2jUjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuVIusCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lo0xeiQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9As2D4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kG3VcpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4JU66iyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTOOike6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnpYreaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzyVFXr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iy2oPvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0sqiN4DiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func id7Nv1UYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2U2nL6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkMVuajCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZ316QKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zn7tXg2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GufOHUumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zK3TyRXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W751vqbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RVOOkuLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWKP2MJGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qMqteSqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzgOwmFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6L25PsENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJCfLOkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuvMYIiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rniAkyrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSa3ed4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGMILXjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdtb1AHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WeUtZNXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gp2bpY0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZBdn34mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4lyr7roaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHoO99M6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19H0DAl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biCW8sMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLXYcDILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i26hjweaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYPlCz75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Phb5GAGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WvUvNc6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyE5DnwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3c5fF0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ezFFoJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiHM48ltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEfgI3VEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxHT2mCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJnEulgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cr1JzLeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56LbFskIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b26Tn21wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwYexuXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rj4t77lUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edtLjQQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNZNURGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFdgjBODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEnN66sDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwNLC2SuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRMAJEPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1jnKg3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hG2oyXccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0egAiEVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9N3EaF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Dyf8KOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCeX0TlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lpdj9FR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5FXfEhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHCtzaYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GvS6BrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2D2xfESZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bc1R4393Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uc4eHxRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvZ2wIyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHHASDDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RLxlGsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJAvShsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMy2lPliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RoeCJVfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54lNNry6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9mtnHpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z93eCMAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGHp1hOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DxWXI2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9If1I7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vblho5NTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLZQpHnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bafdeZ40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XneHkAOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSivN3MKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAtpaTnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPAtvmYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWqaoPxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETs19mx3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTiQoXtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J31lHcJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ADZDmGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFvr4BFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Fqic3MmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBefsiGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LHq1gCwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXizGswLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lay8Y9fBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g355DenLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbV6BwQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNdL0LhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppSAOsldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJIRN1UuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIf0vbcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fO4pWAGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIdNHdq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FciOlqFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gek1c4lnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71piIRwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yd2982RTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHp6PSkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 42BL6l4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRcwLprDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 442OlGtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTDtoaUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gci5X5fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vd52D14fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5G1GBYnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUm9Vm4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgvP9kfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIsgXHDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASBSXEdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Db9sg8rVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wihzmv3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4r3JFIpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DDwY1euWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D5jb5YbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4COJ9fW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXKpdwJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOnUNTCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCzGEdZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHyABaNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZA31ENlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8csJMr8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5a6EOAiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdA9747xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvkn8wTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQ5I4ZVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YJzLvu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9OkfEQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJW4ftucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCZLIBaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGFWOll4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTUfOzXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vk9s6vbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lX8gO7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVnjuOpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Nps4aFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYDNowOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tx8awzKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PtRhM1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvKK7ex6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFTYvKS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qH1asV0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sx0iS07fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCKCum0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8uOS2CiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQ3kIqtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uA7i63wkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g47nz0pZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95ClcxETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOj7PpyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgnVV01hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3rw6YFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21fEEp0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZrQ0s25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZWJ9RwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U2PforDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zohbbf8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHKlWWzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8w4HlxXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9Kznr5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rlJw8ir2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUjrcoP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqMyluQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lD66ZgjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oz9rtRWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oqDUr0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbvakU6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUItcz6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZuij60WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRc7CxAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1yPsC5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u01lxuztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84UZwGcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OI17W4znWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UIJbfP3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2jwsncvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ziRG7m4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEdgqb2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsWct9VVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4dyNcXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ancdw42qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ByzVhCR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVw8UhnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func na97VkBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Erigen4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIXRSXVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIjdlNuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iplj3NtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jf6TvWrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10Z8uSFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEw09velWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tetDYdEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CeyWFGqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWJ32zR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djnuSWipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8azR4beoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17mKAZAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYK5slBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkgwMTYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzWz7OwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

