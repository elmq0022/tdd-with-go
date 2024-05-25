package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	post := Post{
		Title:       readMetaLine(titleSeparator, scanner),
		Description: readMetaLine(descriptionSeparator, scanner),
		Tags:        readTags(scanner),
		Body:        readBody(scanner),
	}

	return post, nil
}

func readMetaLine(tagName string, scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), tagName)
}

func readTags(scanner *bufio.Scanner) []string {
	tags := strings.Split(readMetaLine(tagsSeparator, scanner), ",")
	for i, tag := range tags {
		tags[i] = strings.TrimLeftFunc(
			strings.TrimRightFunc(tag, unicode.IsSpace),
			unicode.IsSpace)
	}
	return tags
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")

	return body
}
