# Database Migrations Script

## Overview

This project provides a simple way to keep databases in sync by tracking and
applying SQL migrations.

## Project Structure

- All required code is located inside the `src` folder.
- To use the script, create the required folder structure inside `src`.
- The files inside `scripts/migrations` are only examples/test cases.

## Migration Naming Convention

When adding a database update, create a new file inside the `migrations`
folder using this format:

`timestamp_action.sql`

- `timestamp` must use the `YYYYMMDDHHMM` format.
- `action` should be a short, descriptive name for the migration.

Example:

`202607061430_create_users_table.sql`

## How to Run Migrations

You can run migrations in one of the following ways:

- manually, by executing the Go script
- automatically, using GitHub Actions

## GitHub Actions Workflow

A GitHub Actions workflow is included in the `src` folder. It automatically runs
the migration script when changes are pushed to the `dev` or `main` branch.

To make the workflow work correctly, you must configure the required environment
secrets for each environment. Refer to `.env.example` for the full list of
required values.

> Note: The ideal approach is to run migrations after deployment, once the
> application is already updated. However, for simplicity and broader
> compatibility, this workflow runs automatically on push.

### Setup Notes

- Create the necessary secrets in GitHub for both the `dev` and `main`
  environments.
- Make sure the secrets match the names expected by the workflow.
- Update `.env.example` if new configuration values are added in the future.

## How It Works

1. The script creates a table to store applied migrations.
2. It reads all migration files from the `migrations` folder.
3. It compares the available files with the migrations already recorded in the
   database.
4. It runs any migrations that have not been applied yet.
5. It records newly applied migrations in the `migrations` table.

## Best Practices

- Keep migration scripts small and focused.
- Avoid editing migrations after they have been applied.
- Use GitHub Actions for automated deployment whenever possible.
- Use manual execution only for local testing or one-off runs.

## License

This project is free to use, including for personal and work purposes.

You may not sell, redistribute for profit, or claim it as your own.ou may not sell, redistribute for profit, or claim it as your own.
