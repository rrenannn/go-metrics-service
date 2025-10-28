CREATE TABLE notifications_log (
                                   id bigserial PRIMARY KEY,
                                   notification_id uuid NOT NULL, -- id do evento (idempotency key)
                                   payload jsonb,
                                   status varchar(32) NOT NULL, -- queued, processing, success, failed
                                   attempts int DEFAULT 0,
                                   error_text text,
                                   created_at timestamptz DEFAULT now(),
                                   updated_at timestamptz DEFAULT now()
);

CREATE INDEX idx_notifications_notification_id ON notifications_log(notification_id);