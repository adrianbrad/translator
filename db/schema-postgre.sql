CREATE TABLE translations
(
  translation_id BIGSERIAL CONSTRAINT translations_pkey PRIMARY KEY,
  created_at     TIMESTAMP DEFAULT now() NOT NULL
);

CREATE TABLE translations_text
(
  translation_id BIGSERIAL NOT NULL CONSTRAINT FK_TranslationsText_Translations REFERENCES translations,
  language_tag   TEXT NOT NULL,
  text           TEXT NOT NULL,
  UNIQUE (translation_id, language_tag)
);