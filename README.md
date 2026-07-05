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

## How It Works

- The script starts by creating a table to store applied migrations.
- It then retrieves all scripts from the `migrations` folder, along with the migrations already recorded in the table.
- It compares both lists and runs any scripts that have not yet been applied.
- Finally, it updates the `migrations` table with the newly applied scripts.

## Notes

- Keep migration scripts small and focused.
- To "run" the migrations you can manually run the go script, but the best option is to auto run it with github actions or web hooks

## License

This project is free to use, including for personal and work purposes.  
You may not sell, redistribute for profit, or claim it as your own.
