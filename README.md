# Database Migrations Script

## Objective

This project provides a simple way to keep databases in sync.

## Structure

- All required code is located inside the `src` folder.
- To use the script, create the required folder structure inside `src`.
- The scripts inside `scripts/migrations` are only examples/test cases.
- You also need a `.env` file containing the database URL.

## How to Use

- When making a database update, create a new script inside the `migrations` folder.
- Migration files should be named using the following format:

  `timestamp_action.sql`

  - `timestamp` is the current date and time in `YYYYMMDDHHMM` format.
  - `action` is a short description of the migration.

- To run the migrations, you can:
  - execute the Go script manually, or
  - automate it using GitHub Actions or webhooks.

## Notes

- Keep migration scripts small and focused.
- To "run" the migrations you can manually run the go script, but the best option is to auto run it with github actions or web hooks
