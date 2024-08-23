// resolver.go

package graph

import (
	"github.com/parth-bhadra/graphql-go-bookstore/graph/model"
)

type Resolver struct {
	books []*model.Book
}

// func (r *mutationResolver) AddBook(ctx context.Context, title string, author string, publishedYear int) (*model.Book, error) {
// 	newBook := &model.Book{
// 		ID:            strconv.Itoa(len(r.books) + 1),
// 		Title:         title,
// 		Author:        author,
// 		PublishedYear: publishedYear,
// 	}

// 	r.books = append(r.books, newBook)
// 	return newBook, nil
// }

// func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
// 	for _, book := range r.books {
// 		if book.ID == id {
// 			return book, nil
// 		}
// 	}
// 	return nil, nil
// }

// func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
// 	return r.books, nil
// }

// Mutation returns the mutation resolver implementation.
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query returns the query resolver implementation.
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
