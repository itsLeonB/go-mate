DROP INDEX IF EXISTS user_subscriptions_user_id_idx;

DROP TABLE IF EXISTS user_subscriptions;

DROP TYPE subscription_model;

DROP INDEX IF EXISTS recommendation_logs_user_recommended_created_idx;

DROP INDEX IF EXISTS recommendation_logs_user_id_created_at_idx;

DROP TABLE IF EXISTS recommendation_logs;

DROP TYPE recommendation_log_status;

DROP INDEX IF EXISTS users_email_idx;

DROP TABLE IF EXISTS users;
