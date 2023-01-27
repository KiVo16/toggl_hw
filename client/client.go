package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "base/ports/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var testToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.hSuLs1RfF4gg5jwjKk-xRBrr1NKfOnr7aWHcVTWnex0"

// test223
func main() {
	fmt.Println("Sending request")

	cc, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := pb.NewQuestionsServiceClient(cc)
	start := time.Now()

	// CreateQuestion(c)
	// DeleteQuestion(c)
	GetQuestions(c)
	// UpdateQuestion(c)

	elapsed := time.Since(start)
	log.Printf("Call took %s", elapsed)
}

func CreateQuestion(c pb.QuestionsServiceClient) {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", testToken)

	req := &pb.CreateQuestionRequest{
		Question: &pb.Question{
			Body: "test",
			Options: []*pb.Option{
				&pb.Option{
					Body:    "t1",
					Correct: true,
				},
				&pb.Option{
					Body:    "t2",
					Correct: false,
				},
			},
		},
	}
	res, err := c.CreateQuestion(ctx, req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())

		} else {
			log.Fatalf("Error while getting response: %v", err)
		}
	}

	fmt.Println("Response: ", res)
}

func DeleteQuestion(c pb.QuestionsServiceClient) {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", testToken)

	req := &pb.DeleteQuestionRequest{
		Id: 5,
	}
	res, err := c.DeleteQuestion(ctx, req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())

		} else {
			log.Fatalf("Error while getting response: %v", err)
		}
	}

	fmt.Println("Response: ", res)
}

func GetQuestions(c pb.QuestionsServiceClient) {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", testToken)

	req := &pb.GetQuestionsRequest{
		Pagination: &pb.Pagination{
			PageSize: 5,
			Page:     1,
		},
	}
	res, err := c.GetQuestions(ctx, req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())

		} else {
			log.Fatalf("Error while getting response: %v", err)
		}
	}

	fmt.Println("Response: ", res)
	fmt.Println("Response len(res.Questions): ", len(res.Questions))
}

func UpdateQuestion(c pb.QuestionsServiceClient) {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", testToken)

	req := &pb.UpdateQuestionRequest{
		Id: 9,
		Ref: &pb.QuestionRef{
			Body: "tetatastas2222211111",
			Options: []*pb.Option{
				&pb.Option{
					Body:    "t55",
					Correct: false,
				},
				&pb.Option{
					Body:    "t2222",
					Correct: true,
				},
				&pb.Option{
					Body:    "t55222",
					Correct: true,
				},
			},
		},
		PropsToUpdate: []pb.UpdateQuestionRequest_UpdateProps{
			pb.UpdateQuestionRequest_BODY,
			pb.UpdateQuestionRequest_OPTIONS,
		},
	}
	res, err := c.UpdateQuestion(ctx, req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())

		} else {
			log.Fatalf("Error while getting response: %v", err)
		}
	}

	fmt.Println("Response: ", res)
}
