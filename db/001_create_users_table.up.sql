CREATE SEQUENCE unique_user_id;

CREATE TABLE users(
  id integer PRIMARY KEY DEFAULT nextval('unique_user_id'),
  username   varchar(255),
  created_at timestamp
)
