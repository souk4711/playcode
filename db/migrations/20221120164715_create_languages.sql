-- migrate:up
CREATE TABLE languages (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  name VARCHAR(255),
  langcode VARCHAR(255),
  cr_langcode VARCHAR(255)
);

-- migrate:down
DROP TABLE languages;
