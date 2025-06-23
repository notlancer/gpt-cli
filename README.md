# GPT-CLI

A simple CLI tool.

---

## Setup: Environment Variables

Create a `.env` file in the project root with the following content:

```env
BearerToken=YOUR_TOKEN_HERE
```

---

## Running with Docker

If you have Docker installed, you can quickly build and run the CLI with these commands:

```bash
docker build -t cli .
docker run -it cli
```

---

## Running Locally with Go

You can also build and run the project directly with Docker (using Go):

```bash
docker build -t cli .
docker run -it cli
```

---

### Notes:

* Make sure your `.env` file is present before running the container.
* The container will run in interactive mode (`-it`) for CLI access.
