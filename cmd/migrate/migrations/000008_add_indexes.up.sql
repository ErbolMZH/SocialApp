CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX IF NOT EXISTS idx_comments_content on comments USING gin(content gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_posts_title on posts USING gin(title gin_trgm_ops);
CREATE INDEX  IF NOT EXISTS idx_posts_tags on posts USING gin(tags);

CREATE INDEX IF NOT EXISTS  idx_users_username on users (username);
CREATE INDEX IF NOT EXISTS  idx_posts_user_id on posts (user_id);
CREATE INDEX IF NOT EXISTS idx_comments_post_id on comments (post_id);



