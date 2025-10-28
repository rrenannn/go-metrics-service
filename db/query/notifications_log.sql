-- name: InsertNotification :exec
INSERT INTO notifications_log(notification_id, payload, status, attempts, error_text)
VALUES($1, $2, $3, $4, $5);