package main

import "fmt"
import "os"
import "bufio"

func main() {
   r := bufio.NewReader(os.Stdin)
   w := bufio.NewWriter(os.Stdout)
   defer w.Flush()
   
   var n int
   
   fmt.Fscan(r, &n)
   matrix := make([][]int, n)
   
   for i := range matrix{
       matrix[i] = make([]int, n)
   }
   
   for i := 0; i < n; i++{
       for j := 0; j < n; j++{
           fmt.Fscan(r, &matrix[i][j])
       }
   }
   
   white, blue := cut(matrix, n)
   fmt.Fprintln(w, white)
   fmt.Fprintln(w, blue)
}

func cut(arr [][]int, size int) (int, int){
    check := check(arr, size * size)
    white := 0
    blue := 0
    
    if check > -1{
        if check == 0{
            white++
        } else {
            blue++
        }
    } else {
        _size := size / 2
        _white := 0
        _blue := 0
        
        _white, _blue = cut(makeArr(0, _size, 0, _size, arr), _size)
        white += _white
        blue += _blue
        
        _white, _blue = cut(makeArr(_size, size, 0, _size, arr), _size)
        white += _white
        blue += _blue
        
        _white, _blue = cut(makeArr(0, _size, _size, size, arr), _size)
        white += _white
        blue += _blue
        
        _white, _blue = cut(makeArr(_size, size, _size, size, arr), _size)
        white += _white
        blue += _blue
    }
    
    return white, blue
}

func check(arr [][]int, goal int) int{
    sum := 0
    
    for _, v := range arr{
        for _, w := range v{
            sum += w
        }
        
        if sum > goal{
            break
        }
    }
    
    if sum == 0 || sum == goal {
        return sum
    }
    
    return -1
}

func makeArr(startY, endY, startX, endX int, arr [][]int) (newArr [][]int){
    newArr = make([][]int, endY - startY)

    for i := startY; i < endY; i++{
        newArr[i - startY] = make([]int, endX - startX)

        for j := startX; j < endX; j++{
            newArr[i - startY][j - startX] = arr[i][j]
        }
    }

    return
}