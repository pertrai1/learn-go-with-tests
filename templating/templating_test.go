package templating

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestTemplating(t *testing.T) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "this is a post",
			Description: "this is a description",
			Tags:        []string{"tag1", "tag2"},
		}
	)

	postrenderer, err := NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postrenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "this is a post",
			Description: "this is a description",
			Tags:        []string{"tag1", "tag2"},
		}
	)

	postrenderer, err := NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postrenderer.Render(io.Discard, aPost)
	}
}
