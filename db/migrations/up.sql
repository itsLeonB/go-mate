CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS users_email_idx ON users (email);

CREATE TYPE recommendation_log_status AS ENUM ('viewed', 'liked', 'passed');

CREATE TABLE IF NOT EXISTS recommendation_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    recommended_user_id UUID NOT NULL REFERENCES users(id),
    status recommendation_log_status NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX recommendation_logs_user_id_created_at_idx
ON recommendation_logs (user_id, created_at);

CREATE INDEX recommendation_logs_user_recommended_created_idx
ON recommendation_logs (user_id, recommended_user_id, created_at);

CREATE TYPE subscription_model AS ENUM ('extra_appearance', 'extra_recommendations');

CREATE TABLE IF NOT EXISTS user_subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    model subscription_model NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expired_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS user_subscriptions_user_id_idx ON user_subscriptions (user_id);