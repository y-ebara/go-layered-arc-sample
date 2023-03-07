-- name: GetTask :one
SELECT
  id,
  description,
  priority,
  start_date,
  due_date,
  done
FROM
  tasks
WHERE
  id = @id
LIMIT 1;

-- name: CreateTask :one
INSERT INTO tasks (
  description,
  priority,
  start_date,
  due_date
)
VALUES (
  @description,
  @priority,
  @start_date,
  @due_date
)
RETURNING id;

-- name: UpdateTask :exec
UPDATE tasks SET
  description = @description,
  priority    = @priority,
  start_date  = @start_date,
  due_date    = @due_date,
  done        = @done
WHERE id = @id;

-- name: GetDownTask :one
SELECT
  id,
  description,
  priority,
  start_date,
  due_date,
FROM
  tasks
WHERE
  id = @id
AND
  done = true
LIMIT 1;