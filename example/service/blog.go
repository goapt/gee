package service

import (
	"context"

	"github.com/goapt/gee/example/proto/demo/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ demo.BlogServiceHTTPServer = (*Blog)(nil)

type Blog struct{}

func (b Blog) Create(ctx context.Context, request *demo.CreateArticleRequest) (*demo.CreateArticleResponse, error) {
	return &demo.CreateArticleResponse{
		Article: &demo.Article{
			Id:        1,
			Title:     "test",
			Content:   "test",
			Like:      123,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}, nil
}

func (b Blog) DeleteArticle(ctx context.Context, request *demo.DeleteArticleRequest) (*demo.DeleteArticleResponse, error) {
	return &demo.DeleteArticleResponse{}, nil
}

func (b Blog) GetArticle(ctx context.Context, request *demo.GetArticleRequest) (*demo.GetArticleResponse, error) {
	return &demo.GetArticleResponse{
		Article: &demo.Article{
			Id:        1,
			Title:     "test",
			Content:   "test",
			Like:      123,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}, nil
}

func (b Blog) ListArticle(ctx context.Context, request *demo.ListArticleRequest) (*demo.ListArticleResponse, error) {
	var articles []*demo.Article

	articles = append(articles, &demo.Article{
		Id:        1,
		Title:     "test",
		Content:   "test",
		Like:      123,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	})

	return &demo.ListArticleResponse{
		Results: articles,
	}, nil
}

func (b Blog) UpdateArticle(ctx context.Context, request *demo.UpdateArticleRequest) (*demo.UpdateArticleResponse, error) {
	return &demo.UpdateArticleResponse{
		Article: &demo.Article{
			Id:        1,
			Title:     "test",
			Content:   "test",
			Like:      123,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}, nil
}
