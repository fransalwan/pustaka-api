# Pustaka API

Pustaka API is a RESTful API for managing a digital library system. It allows users to manage books, authors, categories, and borrowing records.

## Features

- CRUD operations for books, authors, and categories
- User authentication and authorization
- Borrowing and returning books
- Search and filter functionality

## Getting Started

### Prerequisites

- Node.js (v14+)
- npm or yarn
- MongoDB (or your preferred database)

### Installation

```bash
git clone https://github.com/yourusername/pustaka-api.git
cd pustaka-api
npm install
```

### Configuration

Create a `.env` file and set your environment variables:

```env
PORT=3000
MONGODB_URI=mongodb://localhost:27017/pustaka
JWT_SECRET=your_jwt_secret
```

### Running the API

```bash
npm start
```

The API will be available at `http://localhost:3000`.

## API Endpoints

| Method | Endpoint           | Description                |
|--------|--------------------|----------------------------|
| GET    | /books             | List all books             |
| POST   | /books             | Add a new book             |
| GET    | /books/:id         | Get book details           |
| PUT    | /books/:id         | Update a book              |
| DELETE | /books/:id         | Delete a book              |
| ...    | ...                | ...                        |

*See the full API documentation for more endpoints.*

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/feature-name`)
3. Commit your changes
4. Push to the branch
5. Open a pull request

## License

This project is licensed under the MIT License.

---