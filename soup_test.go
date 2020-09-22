package soup

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
      <li>To a <a href="hello.jsp">JSP page</a> right?</li>
      <li>To a <a href="hello">servlet</a></li>
    </ul>
    </p>
    <div id="3">
      <div id="4">Last one</div>
    </div>
    <div id="5">
        <h1><span></span></h1>
    </div>
  </body>
</html>
`

const multipleClassesHTML = `
<html>
	<head>
		<title>Sample Application</title>
	</head>
	<body>
		<div class="first second">Multiple classes</div>
		<div class="first">Single class</div>
		<div class="second first third">Multiple classes inorder</div>
		<div>
			<div class="first">Inner single class</div>
			<div class="first second">Inner multiple classes</div>
			<div class="second first">Inner multiple classes inorder</div>
		</div>
	</body>
</html>
`

var doc = HTMLParse(testHTML)
var multipleClasses = HTMLParse(multipleClassesHTML)

func TestFind(t *testing.T) {
	// Find() and Attrs()
	actual := doc.Find("img").Attrs()["src"]
	assert.Equal(t, "images/springsource.png", actual)
	// Find(...) and Text()
	actual = doc.Find("a", "href", "hello").Text()
	assert.Equal(t, "servlet", actual)
	// Nested Find()
	actual = doc.Find("div").Find("div").Text()
	assert.Equal(t, "Just two divs peacing out", actual)
	// Find("") for any
	actual = multipleClasses.Find("body").Find("").Text()
	assert.Equal(t, "Multiple classes", actual)
	// Find("") with attributes
	actual = doc.Find("", "id", "4").Text()
	assert.Equal(t, "Last one", actual)
}

func TestFindNextPrevElement(t *testing.T) {
	// FindNextSibling() and NodeValue field
	actual := doc.Find("div", "id", "0").FindNextSibling().NodeValue
	assert.Equal(t, "check", strings.TrimSpace(actual))
	// FindPrevSibling() and NodeValue field
	actual = doc.Find("div", "id", "2").FindPrevSibling().NodeValue
	assert.Equal(t, "check", strings.TrimSpace(actual))
	// FindNextElementSibling() and NodeValue field
	actual = doc.Find("table").FindNextElementSibling().NodeValue
	assert.Equal(t, "div", strings.TrimSpace(actual))
	// FindPrevElementSibling() and NodeValue field
	actual = doc.Find("p").FindPrevElementSibling().NodeValue
	assert.Equal(t, "div", strings.TrimSpace(actual))
}

func TestFindAll(t *testing.T) {
	// FindAll() and Attrs()
	allDivs := doc.FindAll("div")
	for i := 0; i < len(allDivs); i++ {
		id := allDivs[i].Attrs()["id"]
		actual, _ := strconv.Atoi(id)
		assert.Equal(t, i, actual)
	}
}

func TestFindAllBySingleClass(t *testing.T) {
	actual := multipleClasses.FindAll("div", "class", "first")
	assert.Equal(t, 6, len(actual))
	actual = multipleClasses.FindAll("div", "class", "third")
	assert.Equal(t, 1, len(actual))
}

func TestFindAllByAttribute(t *testing.T) {
	actual := doc.FindAll("", "id", "2")
	assert.Equal(t, 1, len(actual))
}

func TestFindBySingleClass(t *testing.T) {
	actual := multipleClasses.Find("div", "class", "first")
	assert.Equal(t, "Multiple classes", actual.Text())
	actual = multipleClasses.Find("div", "class", "third")
	assert.Equal(t, "Multiple classes inorder", actual.Text())
}

func TestFindAllStrict(t *testing.T) {
	actual := multipleClasses.FindAllStrict("div", "class", "first second")
	assert.Equal(t, 2, len(actual))
	actual = multipleClasses.FindAllStrict("div", "class", "first third second")
	assert.Equal(t, 0, len(actual))
	actual = multipleClasses.FindAllStrict("div", "class", "second first third")
	assert.Equal(t, 1, len(actual))
}

func TestFindStrict(t *testing.T) {
	actual := multipleClasses.FindStrict("div", "class", "first")
	assert.Equal(t, "Single class", actual.Text())
	actual = multipleClasses.FindStrict("div", "class", "third")
	assert.NotNil(t, actual.Error)
}

func TestText(t *testing.T) {
	// <li>To a <a href="hello.jsp">JSP page</a> right?</li>
	li := doc.Find("ul").Find("li")
	assert.Equal(t, "To a ", li.Text())
}
func TestFullText(t *testing.T) {
	// <li>To a <a href="hello.jsp">JSP page</a> right?</li>
	li := doc.Find("ul").Find("li")
	assert.Equal(t, "To a JSP page right?", li.FullText())
}

func TestFullTextEmpty(t *testing.T) {
	// <div id="5"><h1><span></span></h1></div>
	h1 := doc.Find("div", "id", "5").Find("h1")
	assert.Empty(t, h1.FullText())
}

func TestNewErrorReturnsInspectableError(t *testing.T) {
	err := newError(ErrElementNotFound, "element not found")
	assert.NotNil(t, err)
	assert.Equal(t, ErrElementNotFound, err.Type)
	assert.Equal(t, "element not found", err.Error())
}

func TestFindReturnsInspectableError(t *testing.T) {
	r := doc.Find("bogus", "thing")
	assert.IsType(t, Error{}, r.Error)
	assert.Equal(t, "element `bogus` with attributes `thing` not found", r.Error.Error())
	assert.Equal(t, ErrElementNotFound, r.Error.(Error).Type)
}

// Similar test: https://github.com/hashicorp/go-retryablehttp/blob/master/client_test.go#L616
func TestClient_Post(t *testing.T) {
	// Mock server which always responds 200.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("bad method: %s", r.Method)
		}
		if r.RequestURI != "/foo/bar" {
			t.Fatalf("bad uri: %s", r.RequestURI)
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Fatalf("bad content-type: %s", ct)
		}

		// Check the payload
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		expected := []byte(`{"hello":"world"}`)
		if !bytes.Equal(body, expected) {
			t.Fatalf("bad: %v", string(body))
		}

		w.WriteHeader(200)
	}))
	defer ts.Close()

	// Make the request with JSON payload
	_, err := Post(
		ts.URL+"/foo/bar", "application/json", `{"hello":"world"}`)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// Make the request with byte payload
	_, err = Post(
		ts.URL+"/foo/bar", "application/json", []byte(`{"hello":"world"}`))
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// Make the request with string map payload
	_, err = Post(
		ts.URL+"/foo/bar",
		"application/json",
		map[string]string{
			"hello": "world",
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
}

// Similar test: https://github.com/hashicorp/go-retryablehttp/blob/add-circleci/client_test.go#L631
func TestClient_PostForm(t *testing.T) {
	// Mock server which always responds 200.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("bad method: %s", r.Method)
		}
		if r.RequestURI != "/foo/bar" {
			t.Fatalf("bad uri: %s", r.RequestURI)
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/x-www-form-urlencoded" {
			t.Fatalf("bad content-type: %s", ct)
		}

		// Check the payload
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		expected := []byte(`hello=world`)
		if !bytes.Equal(body, expected) {
			t.Fatalf("bad: %v", string(body))
		}

		w.WriteHeader(200)
	}))
	defer ts.Close()

	// Create the form data.
	form1, err := url.ParseQuery("hello=world")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	form2 := url.Values{
		"hello": []string{"world"},
	}

	// Make the request.
	_, err = PostForm(ts.URL+"/foo/bar", form1)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// Make the request.
	_, err = PostForm(ts.URL+"/foo/bar", form2)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
}

func TestHTML(t *testing.T) {
	li := doc.Find("ul").Find("li")
	assert.Equal(t, "<li>To a <a href=\"hello.jsp\">JSP page</a> right?</li>", li.HTML())

}
