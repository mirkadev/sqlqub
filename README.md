# SQL Sub Util

A simple tool to substitute parameters in SQL queries.

## Usage

1. Run the tool
2. Enter your SQL query with parameters
3. Press enter
4. Get your SQL with substituted parameters

## Example

### Input
```sql
UPDATE "users" SET "bounced" = $1, "updated_at" = CURRENT_TIMESTAMP WHERE ("id" = $2 AND "value" = $3 AND "deleted_at" IS NULL) -- PARAMETERS: [0,"7d1024f7-703d-42f7-bd0b-6e168c6cd0e8","my value"]
```

### Output
```sql
UPDATE "users" SET "bounced" = 0, "updated_at" = CURRENT_TIMESTAMP WHERE ("id" = '7d1024f7-703d-42f7-bd0b-6e168c6cd0e8' AND "value" = 'my value' AND "deleted_at" IS NULL)
```