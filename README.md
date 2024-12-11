# Go Mate

A Go API for dating app.

## Project Structure

- `/cmd` - Entry point of the API.
- `/db` - SQL files for database, including migrations and seeds.
- `/internal` - Internal code for the API:
  - `/appconstant` - Constant values for business logics.
  - `/apperror` - Structured error structs for expected app errors.
  - `/config` - Configurations and settings for the app. ENV vars are loaded here.
