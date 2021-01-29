package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keuller/simple-api/internal/models"
	. "github.com/onsi/gomega"
)

func TestPing(t *testing.T) {
	RegisterTestingT(t)

	app := Server()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	app.ServeHTTP(w, req)

	Expect(w.Code).Should(Equal(200))
	Expect(w.Body.String()).Should(ContainSubstring("It works!"))
	Expect(w.Body.String()).Should(Equal("{\"message\":\"It works!\"}"))
}

func TestListTasks(t *testing.T) {
	t.Skip()

	app := Server()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/tasks", nil)
	app.ServeHTTP(w, req)

	Expect(w.Code).Should(Equal(200))
	Expect(w.Body.String()).ShouldNot(Equal(""))

	var tasks []models.TaskResource
	if err := json.NewDecoder(w.Body).Decode(&tasks); err != nil {
		t.Fatal("Cannot deserialize response to TaskResource")
	}

	Expect(len(tasks) > 0).Should(BeTrue())
}
