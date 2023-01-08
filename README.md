# Database

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(128),
    last_name VARCHAR(128),
    age INT CHECK(age > 0),
    email TEXT UNIQUE NOT NULL
)
```