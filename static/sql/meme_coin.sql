CREATE TABLE IF NOT EXISTS meme_coin (
  id SERIAL PRIMARY KEY,
  name text NOT NULL,
  description text,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  popularity_score INT DEFAULT 0
);
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE UNIQUE INDEX IF NOT EXISTS meme_coin_name_idx ON meme_coin USING btree (digest(name, 'sha256'::text));