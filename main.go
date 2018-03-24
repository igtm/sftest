package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/sync/singleflight"
)

func main() {

	key := "hoge"
	gp := singleflight.Group{}

	e := echo.New()
	e.GET("/:id", func(c echo.Context) error {
		idstr := c.Param("id")
		id, _ := strconv.Atoi(idstr)

		es := &[]Entity{
			Entity{
				ID:   id,
				Name: "iguchi",
			},
		}

		res := <-gp.DoChan(key, func() (interface{}, error) {
			time.Sleep(5000 * time.Millisecond)
			*es = append(*es, Entity{
				ID:   234,
				Name: time.Now().Format("2006-01-02T03:04:05"),
			})
			return *es, nil
		})

		fmt.Printf("v: %q, err: %q, shared: %q\n", res.Val, res.Err, res.Shared)
		s := res.Val.([]Entity)

		return c.JSON(http.StatusOK, s)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type Entity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
