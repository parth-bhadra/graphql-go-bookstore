# Book Management GraphQL API

This project is a GraphQL API for managing a collection of books. It allows you to query for books and add new books to the collection.

## Schema

The GraphQL schema defines the following types and operations:

### Types

- **Book**
  - `id: ID!` - The unique identifier for the book.
  - `title: String!` - The title of the book.
  - `author: String!` - The author of the book.
  - `publishedYear: Int!` - The year the book was published.

### Queries

- **books**: `[Book!]!` - Retrieves a list of all books.
- **book(id: ID!)**: `Book` - Retrieves a single book by its ID.

### Mutations

- **addBook(title: String!, author: String!, publishedYear: Int!)**: `Book!` - Adds a new book to the collection.

## Getting Started

### Prerequisites

- Node.js
- npm or yarn

### Installation

1. Clone the repository:
    ```sh
    git clone <repository-url>
    ```
2. Navigate to the project directory:
    ```sh
    cd <project-directory>
    ```
3. Install the dependencies:
    ```sh
    npm install
    ```
    or
    ```sh
    yarn install
    ```

### Running the Server

To start the GraphQL server, run:
```sh
npm start