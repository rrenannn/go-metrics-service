-- name: InsertNotification :exec
INSERT INTO notifications_log(notification_id, payload, status, attempts, error_text)
VALUES($1, $2, $3, $4, $5);

-- name: UpdateStatus :exec
UPDATE notifications_log
SET status = $2, error_text = $3, updated_at = now()
WHERE notification_id = $1;