-- Stage 1:
-- Games and offers are "raw": they have been fetched by "game providers" and
-- "offers providers".
-- They still need another processing (e.g. deduplicating) before becoming
-- available data.

CREATE TABLE IF NOT EXISTS raw_games (
  game_id TEXT,
  provider_id TEXT,

  title TEXT,
  description TEXT,
  image_url TEXT,

  PRIMARY KEY (game_id, provider_id)
);

CREATE TABLE IF NOT EXISTS raw_offers (
  game_id TEXT,
  game_provider_id TEXT,

  offer_provider_id TEXT,

  regular_price REAL,
  discount_price REAL,
  discount_start TIMESTAMP,
  discount_end TIMESTAMP,
  buy_link TEXT,

  price REAL GENERATED ALWAYS AS (LEAST(regular_price, discount_price)) STORED,
  fetched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_game
    FOREIGN KEY (game_id, game_provider_id)
    REFERENCES raw_games(game_id, provider_id)
);

CREATE INDEX IF NOT EXISTS game_idx ON raw_offers (game_id);

-- Stage 2:
-- This is the data available to use for applications, generated starting from
-- raw data.

-- SLUGS
CREATE TABLE IF NOT EXISTS slug_seq (
  base_slug TEXT PRIMARY KEY,
  counter integer NOT NULL DEFAULT 1
);

CREATE OR REPLACE FUNCTION get_next_slug(TEXT)
RETURNS text VOLATILE LANGUAGE sql AS
$$
  INSERT INTO slug_seq (base_slug) VALUES ($1)
  ON CONFLICT (base_slug) DO UPDATE
    SET counter=slug_seq.counter+1 WHERE slug_seq.base_slug=$1
    RETURNING
      CASE WHEN counter > 1 THEN format('%s-%s', base_slug, counter)
      ELSE base_slug
      END
$$;

CREATE EXTENSION IF NOT EXISTS "unaccent";

CREATE OR REPLACE FUNCTION slugify("value" TEXT)
RETURNS TEXT AS $$
  -- removes accents (diacritic signs) from a given string --
  WITH "unaccented" AS (
    SELECT unaccent("value") AS "value"
  ),
  -- lowercases the string
  "lowercase" AS (
    SELECT lower("value") AS "value"
    FROM "unaccented"
  ),
  -- replaces anything that's not a letter, number, hyphen('-'), or underscore('_') with a hyphen('-')
  "hyphenated" AS (
    SELECT regexp_replace("value", '[^a-z0-9\\-_]+', '-', 'gi') AS "value"
    FROM "lowercase"
  ),
  -- trims hyphens('-') if they exist on the head or tail of the string
  "trimmed" AS (
    SELECT regexp_replace(regexp_replace("value", '-+$', ''), '^-+', '') AS "value"
    FROM "hyphenated"
  ),
  -- adds unique number at the end
  "unique" AS (
    SELECT get_next_slug("value") AS "value"
    FROM "trimmed"
  )

  SELECT "value" FROM "unique";
$$ LANGUAGE SQL STRICT IMMUTABLE;

CREATE TABLE IF NOT EXISTS games (
  id SERIAL PRIMARY KEY,

  title TEXT,
  description TEXT,
  image_url TEXT,

  slug TEXT GENERATED ALWAYS AS (slugify(title)) STORED
);

CREATE INDEX IF NOT EXISTS game_slugs ON games (slug);

-- Map "raw" ID to new internal IDs
CREATE TABLE IF NOT EXISTS id_associations (
  original_game_id TEXT,
  original_game_provider_id TEXT,
  game_id INT,

  CONSTRAINT fk_game
    FOREIGN KEY (game_id)
    REFERENCES games (id),

  CONSTRAINT fk_raw_game
    FOREIGN KEY (original_game_id, original_game_provider_id)
	REFERENCES raw_games (game_id, provider_id)
);

CREATE TABLE IF NOT EXISTS offers (
  game_id INT,
  offer_provider_id TEXT,

  regular_price REAL,
  discount_price REAL,
  discount_start TIMESTAMP,
  discount_end TIMESTAMP,
  buy_link TEXT,

  price REAL GENERATED ALWAYS AS (LEAST(regular_price, discount_price)) STORED,
  fetched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_game
    FOREIGN KEY (game_id)
    REFERENCES games (id)
);
