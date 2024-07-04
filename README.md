POC Gin Framework with Go 1.22

How to Run 
- create database
  - postgres DB
  - database name : gin
  - insert first data
    - role -> USER -> INSERT INTO public.roles (id, role, created_at, updated_at, deleted_at) VALUES (1, 'USER', '2024-07-04 00:00:00.000000 +00:00', '2024-07-04 00:00:00.000000 +00:00', null);
    - role -> ADMIN -> INSERT INTO public.roles (id, role, created_at, updated_at, deleted_at) VALUES (2, 'ADMIN', '2024-07-04 00:00:00.000000 +00:00', '2024-07-04 00:00:00.000000 +00:00', null);
    - user -> user1 role USER
      - INSERT INTO public.users (id, name, email, password, status, role_id, created_at, updated_at, deleted_at)
        VALUES (1, 'user1', 'user@test.test', 'password', 1, 1, '2024-07-04 00:00:00.000000 +00:00', '2024-07-04 00:00:00.000000 +00:00', null);
    - user -> user2 role ADMIN
      - INSERT INTO public.users (id, name, email, password, status, role_id, created_at, updated_at, deleted_at)
        VALUES (2, 'user2', 'admin@test.test', 'password', 1, 2, '2024-07-04 00:00:00.000000 +00:00', '2024-07-04 00:00:00.000000 +00:00', null);
- run -> go mod tidy
- run main.go
