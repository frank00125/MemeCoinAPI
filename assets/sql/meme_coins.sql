CREATE TABLE IF NOT EXISTS meme_coins (
  id SERIAL PRIMARY KEY,
  name text NOT NULL,
  description text,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  popularity_score INT DEFAULT 0
);
-- Set up unique index and constraint for "name" column
CREATE UNIQUE INDEX IF NOT EXISTS meme_coin_name_idx ON meme_coins USING btree (name);