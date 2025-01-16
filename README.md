# SnippetBox

**SnippetBox** is an open-source web application for managing and sharing code snippets. This project is inspired by the work of [Alex Edwards](https://www.alexedwards.net/), based on his book "Let's Go!" and his original SnippetBox application. This version is modified to use **SQLite3** instead of MySQL and minimal dependencies to provide a lightweight and simple solution.

## Features

- Simple and efficient session management using SQLite.
- Built with minimal external dependencies to make it easier to learn and understand.

## Project Structure

## Project Structure

- `/cmd/web/` - Contains the main entry point for the application, with the `main` package and HTTP request handlers for managing snippets.
- `/internal/models/` - Defines data models and handles interactions with SQLite3 for storing and retrieving snippets.
- `/internal/sessions/` - A simplified, custom session management system, based on the core concepts from Alex Edwards' original `scs` package, but further reduced for educational purposes.
- `/internal/store/` - Contains the logic for interacting with the SQLite3 database, including storing and retrieving data for sessions and snippets.
- `/internal/validator/` - Provides validation functionality for user input, ensuring that data adheres to the correct format before being processed.

## Motivation

This project was inspired by Alex Edwards' "Let's Go!" book and his original **SnippetBox** project. It is a simplified, functional implementation of the concepts he introduced, with a specific focus on:

- Using **SQLite3** for data storage instead of MySQL.
- Stripping down unnecessary dependencies and code to make the project easy to understand and focus on the core concepts.
- Custom session management, based on the core ideas from Alex Edwards' `scs` package, but reduced and modified to work with SQLite3 for educational purposes.

## Getting Started

### Prerequisites

- Go 1.18 or higher.
- SQLite3 database.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/BenaliOssama/myforum.git
   cd myforum
   ```

2. Install the dependencies:

   ```bash
   go mod tidy
   ```

3. Create and migrate the SQLite database: (not implemented yet)

   ```bash
   go run ./cmd/migrate/main.go
   ```

4. Run the application:

   ```bash
   make run
   ```

5. Open your browser and go to `http://localhost:9999` to access the application.

## Acknowledgments

This project is based on the original work by **[Alex Edwards](https://www.alexedwards.net/)**, whose book *"Let's Go!"* and code from the **SnippetBox** project inspired much of this work. The session management in this project is heavily influenced by the [Alex Edwards' `scs` package](https://github.com/alexedwards/scs), which has been modified for this project to work with SQLite3.

Many thanks to Alex Edwards for his excellent resources and contributions to the Go community!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
