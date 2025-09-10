package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	if err := StartServerProducts(); err != nil {
		panic(err)
	}
}

// curl http://127.0.0.1:8080/shop/10/products
// curl -X POST -H "Content-Type: application/json" -d "{\"item\":\"pen\"}" http://127.0.0.1:8080/orders
// Требования:
//   - GET /shops/{id}/products -> список
//   - POST /shops/{id}/products {"name":"pen"} -> 201 Created, Location
//   - GET /shops/{id}/products/{pid} -> 200 или 404
//   - /shopProducts?shop=1 -> 404
func StartServerProducts() error {
	mu := http.NewServeMux()

	mu.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "pong")
	})

	mu.HandleFunc("GET /shop/{id}/products/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Products for shop ID: %s", id)
	})

	fmt.Println("Server starting on :8080")
	return http.ListenAndServe(":8080", mu)
}

// Требования:
//   - GET /articles -> 200 JSON-массив в форме {"data":[...]} (минимум один элемент, например {"id":1,"title":"t0"})
//   - POST /articles {"title":"t1"} -> 201 Created, Location: /articles/2, тело {"data":{"id":2,"title":"t1"}}
//   - GET /articles/{id} -> 200 (если есть) или 404 {"error":"not found"}
//   - Любые глагольные пути (например /getArticle?id=1) должны возвращать 404
//   - Content-Type: application/json для ответов с телом
func StartServerArticles() error {
	mux := http.NewServeMux()
	articles := []map[string]interface{}{
		{"id": 1, "title": "t0"},
	}
	nextID := 2

	mux.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(map[string]interface{}{"data": articles})
			return
		}
		if r.Method == http.MethodPost {
			var body map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, `{"error":"bad request"}`, http.StatusBadRequest)
				return
			}
			newArticle := map[string]interface{}{
				"id":    nextID,
				"title": body["title"],
			}
			articles = append(articles, newArticle)
			w.Header().Set("Location", fmt.Sprintf("/articles/%d", nextID))
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{"data": newArticle})
			nextID++
			return
		}
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		id := 0

		_, err := fmt.Sscan(strings.TrimPrefix(r.URL.Path, "/articles/"), &id)
		if err != nil || id < 1 {
			http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
			return
		}
		for _, article := range articles {
			if int(article["id"].(int)) == id {
				json.NewEncoder(w).Encode(map[string]interface{}{"data": article})
				return
			}
		}
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
	})

	return http.ListenAndServe(":8080", mux)
}

// Реализуйте ТОЛЬКО эту функцию.
// Требования:
//   - Начальный профиль: {"name":"Ann","age":30}
//   - GET /profile -> текущее состояние (JSON)
//   - PATCH /profile {"age":31} -> меняет только age
//   - PUT /profile {"name":"Bob","age":20} -> полная замена
//   - DELETE /profile -> 204; повторный DELETE тоже 204 (идемпотентно)
func StartServerProfile() error {
	mu := http.NewServeMux()

	type Profile struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := Profile{Name: "Ann", Age: 30}

	mu.HandleFunc("GET /profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})
	mu.HandleFunc("PATCH /profile", func(w http.ResponseWriter, r *http.Request) {
		profile := new(Profile)
		json.NewDecoder(r.Body).Decode(&profile)
		user.Age = profile.Age
	})
	mu.HandleFunc("PUT /profile", func(w http.ResponseWriter, r *http.Request) {
		profile := new(Profile)
		json.NewDecoder(r.Body).Decode(&profile)
		user = *profile
	})
	mu.HandleFunc("DELETE /profile", func(w http.ResponseWriter, r *http.Request) {
		user = Profile{}
		w.WriteHeader(http.StatusNoContent)
	})

	return http.ListenAndServe(":8080", mu)
}

// Реализуйте ТОЛЬКО эту функцию.
// Требования:
//   - Коллекция пользователей: ids 1,2,3 (хардкод допустим)
//   - GET /users?limit=2&offset=1 ->
//     {"meta":{"limit":2,"offset":1,"total":3},"items":[{"id":2},{"id":3}]}
//   - JSON и корректные заголовки
func StartServerUsers() error {
	mux := http.NewServeMux()

	type User struct {
		ID int `json:"id"`
	}

	type Response struct {
		Meta  map[string]int `json:"meta"`
		Items []User         `json:"items"`
	}

	Users := []User{{ID: 1}, {ID: 2}, {ID: 3}}

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

		body, _ := json.Marshal(Response{
			Meta: map[string]int{
				"limit":  limit,
				"offset": offset,
				"total":  len(Users),
			},
			Items: Users[offset : limit+offset],
		})

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	return http.ListenAndServe(":8080", mux)
}

// Реализуйте ТОЛЬКО эту функцию.
// Требования:
//   - GET /about с Accept: application/json -> CT: application/json; {"service":"shop"}
//   - GET /about с Accept: text/plain -> CT: text/plain; shop
//   - Без Accept допустимо вернуть JSON
func StartServerAbout() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("Content-Type") {
		case "text/plain":
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("shop"))
		default:
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"service":"shop"}`))
		}
	})

	return http.ListenAndServe(":8080", mux)
}

// curl http://127.0.0.1:8080/orders
// curl -X POST -H "Content-Type: application/json" -d "{\"item\":\"pen\"}" http://127.0.0.1:8080/orders
// Требования:
//   - GET /orders/42 -> 200 {"id":42}
//   - GET /orders/999 -> 404 {"error":"not found"}
//   - Любой другой метод на /orders -> 405
func StartServerOrders() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Header().Set("Content-Type", "application/json")

		var response strings.Builder
		if id == "42" {
			response.WriteString(`{"id":42}`)
			w.WriteHeader(http.StatusOK)
		} else {
			response.WriteString(`{"error":"not found"}`)
			w.WriteHeader(http.StatusNotFound)
		}

		fmt.Fprint(w, response.String())
	})

	mux.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		var buf strings.Builder
		buf.WriteString(`{"error":"method not allowed"}`)
		fmt.Fprint(w, buf.String())
	})

	return http.ListenAndServe(":8080", mux)
}

// curl http://localhost:8080/user/Alice
func StartServerName() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		name := r.PathValue("name")
		fmt.Fprintf(w, "Hello, %s!\n", name)
	})

	return http.ListenAndServe(":8080", mux)
}

// curl -X POST -H "Content-Type: application/json" -d "{\"text\":\"hello\"}" http://127.0.0.1:8080/upper
func StartServerAdd() error {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /upper", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Input string `json:"text"`
		}

		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		err := json.NewEncoder(w).Encode(map[string]string{"result": strings.ToUpper(req.Input)})
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	})

	return http.ListenAndServe(":8080", mux)
}

// Реализуйте ТОЛЬКО эту функцию.
// Требования:
//   - 127.0.0.1:8080
//   - GET /ping -> "pong"
//   - middleware добавляет заголовок X-MW: on
func StartServerMiddlerware() error {
	testMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-MW", "on")
			next.ServeHTTP(w, r)
		})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	return http.ListenAndServe(":8080", testMiddleware(mux))
}

// Реализуйте ТОЛЬКО эту функцию.
// Требования:
//   - 127.0.0.1:8080
//   - GET /secure -> ok (если Authorization: Bearer secret123)
//   - иначе -> 401 и тело "unauthorized"
//
// curl -H "Authorization: Bearer secret123" http://localhost:8080/secure
func StartServerBearer() error {
	testMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if auth := r.Header.Get("Authorization"); auth != "Bearer secret123" {
				// http.Error(w, "unauthorized", http.StatusUnauthorized)
				w.Write([]byte("unauthorized"))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /secure", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
		// w.WriteHeader(http.StatusOK)
	})

	return http.ListenAndServe(":8080", testMiddleware(mux))
}

// Реализуйте ТОЛЬКО эту функцию.
// Требования:
//   - 127.0.0.1:8080
//   - GET /rid -> тело "rid"
//   - если есть заголовок X-Request-ID во входящем запросе,
//     добавьте такой же X-Request-ID в ответ с тем же значением
func StartServerRid() error {
	testMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if rid := r.Header.Get("X-Request-ID"); rid != "" {
				w.Header().Set("X-Request-ID", rid)
			}

			next.ServeHTTP(w, r)
		})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /rid", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "rid")
	})

	return http.ListenAndServe(":8080", testMiddleware(mux))
}

// Реализуйте ТОЛЬКО эту функцию.
// Требования:
//   - 127.0.0.1:8080
//   - GET /slow -> тело "done"
//   - middleware добавляет X-Response-Time-ms (целое число, >=50)
func StartServerTime() error {
	testMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Response-Time-ms", fmt.Sprintf("%d", time.Now().UnixMilli()))

			next.ServeHTTP(w, r)
		})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /slow", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "done")
	})

	return http.ListenAndServe(":8080", testMiddleware(mux))
}
