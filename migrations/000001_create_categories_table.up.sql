CREATE TABLE categories (
  id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  parent_id SMALLINT NULL,
  path VARCHAR(255) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL UNIQUE,
  slug VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (parent_id) REFERENCES categories(id) ON DELETE SET NULL, -- 부모 카테고리 삭제 시 null 처리 추가
  CONSTRAINT uq_parent_slug UNIQUE (parent_id, slug) -- 같은 자식 Slug 로 생성 금지
);