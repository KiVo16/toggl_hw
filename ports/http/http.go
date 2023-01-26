package ports

import (
	"base/internal/app"
	"base/internal/app/handlers"
	"base/pkg/pagination"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type HttpServer struct {
	app app.App
}

func NewHttpServer(app app.App) HttpServer {
	return HttpServer{
		app: app,
	}
}

func (s HttpServer) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	question := Question{}
	err := render.Decode(r, &question)
	if err != nil {
		http.Error(w, "invalid-request", http.StatusBadRequest)
		return
	}

	q := NewQuestionToModel(question)

	_, err = s.app.Handlers.CreateQuestion.Handle(r.Context(), handlers.CreateQuestionRequest{
		Question: q,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("internal: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s HttpServer) GetQuestions(w http.ResponseWriter, r *http.Request, params GetQuestionsParams) {
	questions, err := s.app.Handlers.GetQuestions.Handle(r.Context(), handlers.GetQuestionsRequest{
		Pages: pagination.Pagination{
			// Limit:  *params.Limit,
			// Offset: *params.Offset,
		},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("internal: %v", err), http.StatusInternalServerError)
		return
	}

	fQuestions := make([]Question, len(questions))
	for i, q := range questions {
		fQuestions[i] = NewQuestionFromModel(q)
	}

	render.Respond(w, r, fQuestions)
	w.WriteHeader(http.StatusOK)
}

func (s HttpServer) DeleteQuestion(w http.ResponseWriter, r *http.Request, id int) {
	_, err := s.app.Handlers.DeleteQuestion.Handle(r.Context(), handlers.DeleteQuestionRequest{
		ID: id,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("internal: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s HttpServer) UpdateQuestion(w http.ResponseWriter, r *http.Request, id int) {
	_, err := s.app.Handlers.DeleteQuestion.Handle(r.Context(), handlers.DeleteQuestionRequest{
		ID: id,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("internal: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
