package soup

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

const testHTML = `
<html>
  <head>
    <title>Sample "Hello, World" Application</title>
  </head>
  <body bgcolor=white>

    <table border="0" cellpadding="10">
      <tr>
        <td>
          <img src="images/springsource.png">
        </td>
        <td>
          <h1>Sample "Hello, World" Application</h1>
        </td>
      </tr>
    </table>
    <div id="0">
      <div id="1">Just two divs peacing out</div>
    </div>
    check
    <div id="2">One more</div>
    <p>This is the home page for the HelloWorld Web application. </p>
    <p>To prove that they work, you can execute either of the following links:
    <ul>
      <li>To a <a href="hello.jsp">JSP page</a></li>
      <li>To a <a href="hello">servlet</a></li>
    </ul>
    </p>
    <div id="3">
      <div id="4">Last one</div>
    </div>

  </body>
</html>
`

var doc = HTMLParse(testHTML)

func TestFind(t *testing.T) {
	// Find() and Attrs()
	actual := doc.Find("img").Attrs()["src"]
	if !reflect.DeepEqual(actual, "images/springsource.png") {
		t.Error("Instead of `images/springsource.png`, got", actual)
	}
	// Find(...) and Text()
	actual = doc.Find("a", "href", "hello").Text()
	if !reflect.DeepEqual(actual, "servlet") {
		t.Error("Instead of `servlet`, got", actual)
	}
	// Nested Find()
	actual = doc.Find("div").Find("div").Text()
	if !reflect.DeepEqual(actual, "Just two divs peacing out") {
		t.Error("Instead of `Just two divs peacing out`, got", actual)
	}
}

func TestFindNextPrevElement(t *testing.T) {
	// FindNextSibling() and NodeValue field
	actual := doc.Find("div", "id", "0").FindNextSibling().NodeValue
	if !reflect.DeepEqual(strings.TrimSpace(actual), "check") {
		t.Error("Instead of `check`, got", actual)
	}
	// FindPrevSibling() and NodeValue field
	actual = doc.Find("div", "id", "2").FindPrevSibling().NodeValue
	if !reflect.DeepEqual(strings.TrimSpace(actual), "check") {
		t.Error("Instead of `check`, got", actual)
	}
	// FindNextElementSibling() and NodeValue field
	actual = doc.Find("table").FindNextElementSibling().NodeValue
	if !reflect.DeepEqual(actual, "div") {
		t.Error("Instead of `div`, got", actual)
	}
	// FindPrevElementSibling() and NodeValue field
	actual = doc.Find("p").FindPrevElementSibling().NodeValue
	if !reflect.DeepEqual(actual, "div") {
		t.Error("Instead of `div`, got", actual)
	}
}

func TestFindAll(t *testing.T) {
	// FindAll() and Attrs()
	allDivs := doc.FindAll("div")
	for i := 0; i < len(allDivs); i++ {
		id := allDivs[i].Attrs()["id"]
		actual, _ := strconv.Atoi(id)
		if !reflect.DeepEqual(actual, i) {
			t.Error("Instead of", i, "got", actual)
		}
	}
}
