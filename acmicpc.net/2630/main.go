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
   
   //white, blue := cut(matrix, n)
   fmt.Fprintln(w, matrix[0:n / 2 - 1][n / 2 :n-1])
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
        
        white, blue = cut(arr[:_size - 1][:_size - 1], _size)
        white += _white
        blue += _blue
        
        white, blue = cut(arr[_size:][:_size - 1], _size)
        white += _white
        blue += _blue
        
        white, blue = cut(arr[:_size - 1][_size:], _size)
        white += _white
        blue += _blue
        
        white, blue = cut(arr[_size:][_size:], _size)
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
