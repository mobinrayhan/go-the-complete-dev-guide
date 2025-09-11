Perfect 👍 — let’s go deeper and make it **core-level**, with exact fields, roles, and auth rules spelled out.

---

# 📝 Secure Notes API (Go Learning Project)

## 📂 Core Entities & Fields

### 1. **User**

- `id` (UUID or bigint auto-increment)
- `email` (string, unique, required)
- `username` (string, unique, required)
- `password_hash` (string, required, stored using bcrypt)
- `created_at` (timestamp)
- `updated_at` (timestamp)

👉 **Input when registering:**

```json
{
  "email": "test@example.com",
  "username": "john_doe",
  "password": "MyStrongPassword123"
}
```

👉 **Stored internally (hashed password, no plaintext):**

```json
{
  "id": 1,
  "email": "test@example.com",
  "username": "john_doe",
  "password_hash": "$2a$10$.......",
  "created_at": "2025-09-11T12:00:00Z"
}
```

---

### 2. **Note**

- `id` (UUID or bigint auto-increment)
- `user_id` (foreign key → users.id)
- `title` (string, required, max length \~255)
- `body` (text, required)
- `tags` (array of strings or JSON field, optional)
- `created_at` (timestamp)
- `updated_at` (timestamp)

👉 **Input when creating a note:**

```json
{
  "title": "My first note",
  "body": "This is a secure note only I can see.",
  "tags": ["personal", "todo"]
}
```

---

### 3. **Auth Tokens**

- Access Token (short-lived, e.g., 15m)
- Refresh Token (long-lived, e.g., 7d)
- Stored in DB (optional, for invalidation / logout tracking).

---

## 🔑 Authentication & Authorization Rules

1. **Registration**

   - Anyone (public, unauthenticated).
   - Requires `email + username + password`.
   - Password hashed before saving.

2. **Login**

   - Anyone (public).
   - Requires `email + password`.
   - Valid credentials → issue Access + Refresh tokens.

3. **Authenticated Requests**

   - Must send JWT Access Token in `Authorization: Bearer <token>` header.
   - Middleware validates:

     - Token signature.
     - Expiry.
     - Extracts `user_id`.

4. **Notes**

   - Only the **authenticated user** can create, read, update, delete **their own** notes.
   - User **A** cannot access notes of User **B** (check `note.user_id == auth.user_id`).

5. **Password Reset**

   - User provides `email`.
   - System generates `reset_token` (random string, expires in X min).
   - Token emailed (in real project; here, just log/print).
   - User calls `POST /auth/reset-password` with:

     ```json
     {
       "token": "RESET123",
       "new_password": "NewSecret123"
     }
     ```

6. **Password Update (while logged in)**

   - Authenticated user.
   - Must provide old password + new password.
   - System validates old password hash, then updates.

---

## 🔐 Example Endpoints & Inputs/Outputs

### `POST /auth/register`

Request:

```json
{
  "email": "test@example.com",
  "username": "john_doe",
  "password": "password123"
}
```

Response:

```json
{
  "message": "user registered successfully"
}
```

---

### `POST /auth/login`

Request:

```json
{
  "email": "test@example.com",
  "password": "password123"
}
```

Response:

```json
{
  "access_token": "<JWT_ACCESS>",
  "refresh_token": "<JWT_REFRESH>"
}
```

---

### `GET /notes` (with `Authorization: Bearer <token>`)

Response:

```json
[
  {
    "id": 1,
    "title": "My first note",
    "body": "This is a secure note.",
    "tags": ["personal"],
    "created_at": "2025-09-11T12:00:00Z"
  }
]
```

---

### `POST /notes`

Request:

```json
{
  "title": "Groceries",
  "body": "Milk, Eggs, Bread",
  "tags": ["shopping", "personal"]
}
```

Response:

```json
{
  "id": 2,
  "title": "Groceries",
  "body": "Milk, Eggs, Bread",
  "tags": ["shopping", "personal"],
  "created_at": "2025-09-11T13:00:00Z"
}
```

---

## 🎯 Core Learning Concepts You’ll Practice

- Password hashing with `bcrypt`.
- JWT generation/validation middleware.
- Role-based access (user’s own resources only).
- Secure DB queries (SQL injection prevention).
- Refresh tokens & token expiry handling.
- Folder structure: `cmd/server`, `internal/auth`, `internal/notes`, `internal/models`.

---

👉 Do you want me to **sketch the Go folder structure + minimal starter code (handlers, middleware, DB setup)** so you can copy-paste and start coding right away?
