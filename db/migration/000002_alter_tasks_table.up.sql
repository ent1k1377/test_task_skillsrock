ALTER TABLE tasks RENAME COLUMN update_at TO updated_at;

ALTER TABLE tasks
ALTER COLUMN status SET DEFAULT 'new',
ADD CONSTRAINT status_check CHECK (status IN ('new', 'in_progress', 'done')),
ALTER COLUMN created_at SET DEFAULT now(),
ALTER COLUMN updated_at SET DEFAULT now();