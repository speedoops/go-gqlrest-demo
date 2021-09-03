package graph

import (
	"fmt"
	"testing"
	"time"

	gqlcli "github.com/99designs/gqlgen/client"
	"github.com/speedoops/go-gqlrest-demo/graph/engine"
	"github.com/speedoops/go-gqlrest-demo/graph/model"
	restcli "github.com/speedoops/go-gqlrest/client"
	"github.com/stretchr/testify/require"
)

const T9527 = "T9527" // 这是一个特定的值，第一个创建的 Todo 的 ID

type Todo struct {
	ID       string           `json:"id"`
	Text     string           `json:"text"`
	Done     bool             `json:"done"`
	User     model.User       `json:"user"`
	Category []model.Category `json:"category"`
}

func TestTodo(t *testing.T) {
	srv := engine.NewMockServer(&Resolver{})
	c := gqlcli.New(srv, gqlcli.Path("/query"))

	t.Run("mutation.createTodo", func(t *testing.T) {
		var resp struct {
			CreateTodo Todo `json:"createTodo"`
		}

		mutation := `
		mutation createTodo($text:String!) {
			createTodo(input: {userID:"uid", text:$text}){
			  id,text,done
			}
		  }
		`

		genRandomText := func() string {
			return fmt.Sprintf("text_%s", time.Now().Format("2006-01-02 15:04:05"))
		}

		c.MustPost(mutation, &resp, gqlcli.Var("text", genRandomText()))
		c.MustPost(mutation, &resp, gqlcli.Var("text", genRandomText()))
		c.MustPost(mutation, &resp, gqlcli.Var("text", genRandomText()))

		require.NotEmpty(t, resp.CreateTodo)
		t.Logf("%+v", resp.CreateTodo)
	})

	t.Run("mutation.updateTodo", func(t *testing.T) {
		var resp struct {
			UpdateTodo Todo `json:"updateTodo"`
		}

		mutation := `
		mutation updateTodo($id: ID!) {
			updateTodo(input: {id: $id, userID:"uid", text:"9527.Updated"}){
			  id,text,done
			}
		  }
		`
		c.MustPost(mutation, &resp, gqlcli.Var("id", T9527))

		require.NotEmpty(t, resp.UpdateTodo)
		t.Logf("%+v", resp.UpdateTodo)
	})

	t.Run("query.todos should contain T9527", func(t *testing.T) {
		var resp struct {
			Todos []Todo `json:"todos"`
		}
		query := `
		query todos {
			todos(ids:["T9527"],userId2:"userID2",text2:"text2",done2:true) {
				id,text,done
			}
		}
		`
		c.MustPost(query, &resp)

		require.NotEmpty(t, resp.Todos)
		for _, v := range resp.Todos {
			if v.ID == T9527 {
				return
			}
		}
		require.Fail(t, "T9527 not found")
		t.Logf("%+v", resp.Todos)
	})

	t.Run("mutation.deleteTodo", func(t *testing.T) {
		var resp struct {
			DeleteTodo bool `json:"deleteTodo"`
		}

		mutation := `
		mutation deleteTodo($id: ID!) {
			deleteTodo(id: $id)
		  }
		`
		c.MustPost(mutation, &resp, gqlcli.Var("id", T9527))

		require.NotEmpty(t, resp.DeleteTodo)
		t.Logf("%+v", resp.DeleteTodo)
	})

	t.Run("query.todos should not contain T9527", func(t *testing.T) {
		var resp struct {
			Todos []Todo `json:"todos"`
		}
		query := `
		query todos {
			todos(ids:["T9527"],userId2:"userID2",text2:"text2",done2:true) {
				id,text,done
			}
		}
		`
		c.MustPost(query, &resp)

		require.NotEmpty(t, resp.Todos)
		for _, v := range resp.Todos {
			require.NotEqual(t, v.ID, T9527)
		}
		t.Logf("%+v", resp.Todos)
	})
}

func TestTodos_REST(t *testing.T) {
	s := engine.NewMockServer(&Resolver{})
	c := restcli.New(s, restcli.Prefix(""))

	t.Run("rest.createTodo", func(t *testing.T) {
		var resp struct {
			//nolint:staticcheck,revive // ignore SA5008: unknown JSON option "squash"
			Todo `json:",squash"`
		}

		payload := `{"input": {"userID":"uid", "text":"$text"}}`
		err := c.Post("/api/v1/todo", &resp, restcli.Body(payload))
		require.Nil(t, err)
		require.NotEmpty(t, resp.Todo)

		t.Logf("%+v", resp.Todo)
	})

	t.Run("rest.updateTodo", func(t *testing.T) {
		var resp struct {
			//nolint:staticcheck,revive // ignore SA5008: unknown JSON option "squash"
			Todo `json:",squash"`
		}

		payload := `{"input": {"userID":"uid", "text":"$text.Updated"}}`
		err := c.Put("/api/v1/todo/T9527", &resp, restcli.Body(payload))
		require.Nil(t, err)
		require.NotEmpty(t, resp.Todo)

		t.Logf("%+v", resp.Todo)
	})

	t.Run("rest.deleteTodo", func(t *testing.T) {
		err := c.Delete("/api/v1/todo/T9527", nil)
		require.Nil(t, err)
	})

	t.Run("rest.todos", func(t *testing.T) {
		var resp struct {
			Todos []Todo `json:"results"`
		}

		err := c.Get("/api/v1/todos?ids=T9527&userId2=userId2&text2=text2&done2=true", &resp)
		require.Nil(t, err)

		for _, v := range resp.Todos {
			if v.ID == T9527 {
				t.Fail()
			}
		}
		t.Logf("%+v", resp.Todos)
	})
}

func TestTodos_POST(t *testing.T) {
	s := engine.NewMockServer(&Resolver{})
	c := restcli.New(s, restcli.Prefix(""))

	t.Run("rest.createTodo", func(t *testing.T) {
		var resp struct {
			//nolint:staticcheck,revive // ignore SA5008: unknown JSON option "squash"
			Todo `json:",squash"`
		}

		payload := `{"userID":"uid", "text":"$text"}` // 方式1
		//payload = fmt.Sprintf(`{"input": %s}`, payload) // 方式2

		err := c.Post("/api/v1/todo", &resp, restcli.Body(payload))
		require.Nil(t, err)
		require.NotEmpty(t, resp.Todo)

		t.Logf("%+v", resp.Todo)
	})
}

func TestTodos_GET(t *testing.T) {
	s := engine.NewMockServer(&Resolver{})
	c := restcli.New(s, restcli.Prefix(""))

	t.Run("rest.todos", func(t *testing.T) {
		var resp struct {
			Todos []Todo `json:"results"`
		}

		err := c.Get("/api/v1/todos?ids=T9527&userId2=userId2&text2=text2&done2=true", &resp)
		require.Nil(t, err)

		for _, v := range resp.Todos {
			if v.ID == T9527 {
				t.Fail()
			}
		}
		t.Logf("%+v", resp.Todos)
	})
}

func TestTodos_PUT(t *testing.T) {
	s := engine.NewMockServer(&Resolver{})
	c := restcli.New(s, restcli.Prefix(""))

	t.Run("rest.updateTodo", func(t *testing.T) {
		var resp struct {
			//nolint:staticcheck,revive // ignore SA5008: unknown JSON option "squash"
			Todo `json:",squash"`
		}

		payload := `{"text":"$text.Updated"}`

		err := c.Put("/api/v1/todo/T9527?userID=uid", &resp, restcli.Body(payload))
		require.NotNil(t, err)
		t.Logf("%+v", err)
		require.Contains(t, err.Error(), "code 404")
	})
}

func TestTodos_DELETE(t *testing.T) {
	s := engine.NewMockServer(&Resolver{})
	c := restcli.New(s, restcli.Prefix(""))

	t.Run("rest.deleteTodo", func(t *testing.T) {
		err := c.Delete("/api/v1/todo/T9527", nil)
		require.Nil(t, err)
	})

	t.Run("rest.deleteTodoByUser", func(t *testing.T) {
		err := c.Delete("/api/v1/todos/?userID=T9527", nil)
		require.NotNil(t, err)
		t.Logf("%+v", err)
		require.Contains(t, err.Error(), "http 404")
	})
}
