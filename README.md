# BookTracker-CLI-App

One-line description: A Go-based command-line interface (CLI) application for tracking and managing book reading progress using the Cobra library.

Summary: This Go program, `BookTracker`, leverages the Cobra library to provide a structured and user-friendly CLI for managing a reading list. It allows users to add books, update reading progress, and view reading statistics. Each book is stored as a JSON object in a file, allowing persistent storage across sessions.

**Core Features:**
- **Add a Book:** Users can add a new book to their tracker by specifying the title, author, and total number of pages. This information is saved in a JSON file called `books.json`.
- **Update Reading Progress:** Users can update the current page number they've reached for a specific book. The program searches for the book by title in the JSON data and updates the `currentPage` value.
- **View Statistics:** The statistics command provides a detailed look at each tracked book, showing the current progress, total days since the book was added, and average reading pace calculated in pages per day.

**Technical Implementation:**
- **Data Storage:** Book data is serialized into JSON and stored in a file, simplifying CRUD operations and ensuring data persistence.
- **Error Handling:** Basic error handling is implemented for file operations and JSON marshaling, though improvements could be made to handle specific errors more gracefully.
- **Modular Commands:** Using Cobra, the application structures each operation as a command, making it easy to extend with new features such as removing a book or listing all books.

This application serves as a practical example of building CLI tools in Go that are capable of performing file I/O operations, handling JSON data, and providing an interactive and easy-to-use command interface.
