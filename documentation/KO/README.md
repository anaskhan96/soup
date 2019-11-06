# soup
[![Build Status](https://travis-ci.org/anaskhan96/soup.svg?branch=master)](https://travis-ci.org/anaskhan96/soup)
[![GoDoc](https://godoc.org/github.com/anaskhan96/soup?status.svg)](https://godoc.org/github.com/anaskhan96/soup)
[![Go Report Card](https://goreportcard.com/badge/github.com/anaskhan96/soup)](https://goreportcard.com/report/github.com/anaskhan96/soup)

**파이썬 언어의 Beautifulsoup과 비슷한 GO 언어의 웹 스크래퍼**

*soup*는 Go용 작은 웹 스크레이퍼 패키지로, BeautifulSoup과 인터페이스가 매우 유사합니다.

지금까지 구현하여 작업 완료한 변수 및 함수:
```go
var Headers map[string]string // Header()를 개별적으로 호출하는 대신 키 값 쌍의 맵으로 설정
var Cookies map[string]string // Cookie()를 개별적으로 호출하는 대신 Cookie를 키-값 쌍의 지도로 설정 
func Get(string) (string,error){} // url을 인자로 받아서, HTML 문자열을 반환하는 함수
func GetWithClient(string, *http.Client){} // url과 사용자 지정 HTTP 클라이언트를 인자로 받아서 HTML 문자열을 반환하는 함수
func Header(string, string){} // 키, 값 쌍을 받아서 구현되어 있는 Get() 함수의 HTTP 요청에 대한 헤더로 정의
func Cookie(string, string){} // 키, 값 쌍을 받아서 Get() 함수의 HTTP 요청과 함께 보낼 쿠키로 설정
func HTMLParse(string) Root {} // HTML 문자열을 인수로 사용하여 구성된 DOM에 포인터를 반환
func Find([]string) Root {} // 인자 태그, (키-값 쌍의 속성) 을 인자로 받아서, 첫번째 경우에 대한 포인터를 반환
func FindAll([]string) []Root {} // Find() 함수와 같지만, 모든 경우에 대한 포인터를 반환
func FindStrict([]string) Root {} // 태그 요소, (키-값 쌍의 속성) 을 인자로 사용하여, 첫번째 경우의 포인터를 정확히 매칭되는 값과 함께 반환
func FindAllStrict([]string) []Root {} // FindStrict() 함수와 같지만, 모든 경우의 포인터가 반환
func FindNextSibling() Root {} // DOM에서 요소의 다음 형제에 대한 포인터 반환됨
func FindNextElementSibling() Root {} // DOM에서 요소의 다음 요소 형제에 대한 포인터 반환됨
func FindPrevSibling() Root {} // DOM의 이전 요소 형제에 대한 포인터 반환
func FindPrevElementSibling() Root {} // DOM에서 요소의 이전 요소 형제에 대한 포인터 반환
func Children() []Root {} // 이 DOM 요소에서 모든 직속 자식을 찾기
func Attrs() map[string]string {} // 각 값을 조회하는 요소의 모든 특성을 사용하여 Map이 반환됨
func Text() string {} // 중첩되지 않은 태그의 전체 텍스트, 첫 번째 텍스트가 중첩되지 않은 태그로 반환됨
func FullText() string {} // 중첩되거나/중첩되지 않은 태그에서 텍스트 전체를 반환함
func SetDebug(bool) {} // 디버깅 모드를 true나 false로 설정함; false가 기본 디폴트
```

`Root` 는 세 개의 필드를 가지는 구조입니다 :
* `Pointer` 현재의 html 노드를 포함하고 있는 포인터입니다.
* `NodeValue` 현재의 html 노드의 값을 포함함, 예. ElementNode의 태그 이름 또는 텍스트 노드의 경우 텍스트입니다.
* `Error` 오류 발생 시 오류를 포함하고, 그렇지 않으면 `nil` 이 반환됩니다.

## 설치방법
이 커맨드를 사용하여 이 패키지를 설치 하실 수 있습니다.
```bash
go get github.com/anaskhan96/soup
```

## 사용예시
이 예시 코드는 "Comics I Enjoy"의 (텍스트와 이것들의 링크)의 일부를 [xkcd](https://xkcd.com) 사이트에서 스크래핑 한 것입니다.


[더 많은 예제들](https://github.com/anaskhan96/soup/tree/master/examples)
```go
package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {
	resp, err := soup.Get("https://xkcd.com")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "comicLinks").FindAll("a")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}
```

## 기여방법
이 패키지는 저의 자유시간에 개발되었습니다. 그러나, 오픈소스 사회의 모든 사람들이 이것을 더 나은 웹 스크레퍼로 만들기 위해 기여하는 것은 환영합니다. 패키지에 특정 기능 또는 기능이 포함되어 있어야 한다고 생각되면 언제든지 새로운 Issue를 열거나 PR을 끌어다 놓으십시오.
