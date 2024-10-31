# Kawan Sehat Backend

This project is a back-end service for the Kawan Sehat project, developed as a submission for the [BPJS Healthkaton](https://healthkathon.bpjs-kesehatan.go.id/) in the Innovation System category. Kawan Sehat is a platform similar to Reddit, designed specifically for chronic disease patients to share their experiences and encourage others. Additionally, medical experts can participate by answering questions and providing professional advice. The back-end is written in [Go](https://go.dev/) and utilizes a [PostgreSQL](https://www.postgresql.org/) database to manage and store data.

The project leverages several libraries to enhance its functionality. It uses [sqlc](https://docs.sqlc.dev/) for SQL code generation and [pgx](https://pkg.go.dev/github.com/jackc/pgx/v5) as the PostgreSQL driver. For HTTP routing, the project utilizes the [chi](https://go-chi.io/) library, and for request validation, it employs the [validator](https://pkg.go.dev/github.com/go-playground/validator/v10) library.

## Database Schema

The database schema can be publicly accessed at [dbdocs](https://dbdocs.io/bagashiz/kawan-sehat).

## API Documentation

The API documentation can be accessed at [Postman](https://documenter.getpostman.com/view/23547657/2sAY4vhNRZ).
