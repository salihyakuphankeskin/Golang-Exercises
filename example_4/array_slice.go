package  array_slice

import ("fmt")
func Runner() {
    var a [5]int
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
    print('\n')

    var zeroLength []string
    fmt.Println("uninit:", zeroLength, zeroLength == nil, len(zeroLength) == 0)

    threeLength := make([]string, 3)
    fmt.Println("emp:", threeLength, "len:", len(threeLength), "cap:", cap(threeLength))

    alphabet := make([]string, 5)
    for i := range alphabet{
        alphabet[i]="0"
    }
    alphabet[0],alphabet[1],alphabet[2] = "a","b","c"
    fmt.Println("alphabet ", alphabet)
    alphabet = append(alphabet, "f")
    fmt.Println("alphabet appended ", alphabet)

    alphabetBeta:= make([]string, 4)
    copy(alphabetBeta, alphabet)
    fmt.Println("alphabet copy first 4 elements", alphabetBeta)



}