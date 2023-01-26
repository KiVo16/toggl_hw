// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package ports

// Option defines model for Option.
type Option struct {
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
	Id      *int   `json:"id,omitempty"`
}

// Question defines model for Question.
type Question struct {
	Body string `json:"body"`
	Id   *int   `json:"id,omitempty"`

	// Options Options associated with question
	Options *[]Option `json:"options,omitempty"`
}

// GetQuestionsParams defines parameters for GetQuestions.
type GetQuestionsParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// CreateQuestionJSONRequestBody defines body for CreateQuestion for application/json ContentType.
type CreateQuestionJSONRequestBody = Question

// UpdateQuestionJSONRequestBody defines body for UpdateQuestion for application/json ContentType.
type UpdateQuestionJSONRequestBody = Question
