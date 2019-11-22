package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
)

// This is a basic 'soup' example for Korean.
// 이 예제는 'soup'의 가장 기초적인 한국인을 위한 예제입니다.

// This is Easy Coding Test problem for Korean from 'BaekJoon'
// 이 예제는 백준 코딩 테스트의 쉬운 문제입니다.


func main() {	
	url := fmt.Sprintf("https://www.acmicpc.net/problem/1000")
	// url의 마지막 숫자에 1000이외에 다른 숫자를 추가시키면,
	// 다른 문제가 등장합니다.

	resp, _ := soup.Get(url)
	doc := soup.HTMLParse(resp)
	problem := doc.Find("p").Text()
	fmt.Println("문제 : ", problem)
}

// result example
// 결과 예시
// 두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오.

