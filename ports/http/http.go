package ports

import (
	"base/internal/app"
	"base/internal/app/handlers"
	"base/internal/domain/model"
	e "base/pkg/http/errors"
	"base/pkg/pagination"
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
		e.NewHttpError(err).
			WithCode(http.StatusInternalServerError).
			Handle(w)

		return
	}

	q := NewQuestionToModel(question)

	fQuestion, err := s.app.Handlers.CreateQuestion.Handle(r.Context(), handlers.CreateQuestionRequest{
		Question: q,
	})
	if err != nil {
		e.NewHttpError(err).Handle(w)
		return
	}

	render.Respond(w, r, NewQuestionFromModel(*fQuestion))
	w.WriteHeader(http.StatusCreated)
}

func (s HttpServer) GetQuestions(w http.ResponseWriter, r *http.Request, params GetQuestionsParams) {
	req := handlers.GetQuestionsRequest{Pages: pagination.Pagination{}}

	if params.PageSize != nil {
		req.Pages.PageSize = *params.PageSize
	}

	if params.Page != nil {
		req.Pages.Page = *params.Page
	}

	questions, err := s.app.Handlers.GetQuestions.Handle(r.Context(), req)
	if err != nil {
		e.NewHttpError(err).Handle(w)
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
		e.NewHttpError(err).Handle(w)
		return
	}

	render.Respond(w, r, struct{}{})
	w.WriteHeader(http.StatusOK)
}

func (s HttpServer) UpdateQuestion(w http.ResponseWriter, r *http.Request, id int) {
	ref := QuestionRef{}
	err := render.Decode(r, &ref)
	if err != nil {
		e.NewHttpError(err).
			WithCode(http.StatusInternalServerError).
			Handle(w)

		return
	}

	req := handlers.UpdateQuestionRequest{
		ID:   id,
		Body: ref.Body,
	}

	if ref.Options != nil {
		options := make([]model.Option, len(*ref.Options))
		for i, o := range *ref.Options {
			options[i] = NewOptionToModel(o)
		}

		req.Options = &options
	}

	fQuestion, err := s.app.Handlers.UpdateQuestion.Handle(r.Context(), req)
	if err != nil {
		e.NewHttpError(err).Handle(w)
		return
	}

	render.Respond(w, r, NewQuestionFromModel(*fQuestion))
	w.WriteHeader(http.StatusOK)
}
