package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	const PUT = 1
	const DEL = 2
	var n, m, b int
	fmt.Fscan(r, &n, &m, &b)
	
	matrix := make([][]int, n)
	sum := 0.0

	for i := 0; i < n; i++{
		matrix[i] = make([]int, m)

		for j := 0; j < m; j++{
			fmt.Fscan(r, &matrix[i][j])
			sum += float64(matrix[i][j])
		}
	}

	flatHeight := math.Round(sum / float64(n * m))

	// 놓는 시간 1초
	// 제거 시간 2초
	// ex) 놓을게 5개 있고, 제거할게 3개가 있다. 머할래?
	// 제거할게 적더라도, 놓는게 이득이다. 밤에 무서운 몬스터 나옴 빨리해야됨.

	// 평평한 기준이 얼마인지?
	// => 빈번하게 나타난 숫자를 기준으로 삼는다. (x)
	// => 평균 높이를 구한다.

	// 놓을게 많은지? 제거할게 많은지? 
	// >> 시간 체크 시작
	// if 제거할 때 시간 이득이다?
	// 바로 제거

	// if 놓을 때 인벤 충분 ?
	// YES => 바로 시간체크
	// NO => 제거해서 놓아본다. => 시간체크
	
	fmt.Println(int(flatHeight))
}