DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS conversations;

CREATE TABLE conversations (
  id serial,
  subject text,
  created_at timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE messages (
  id serial,
  message_body text,
  created_at timestamp,
  sender_id int,
  conversation_id int references conversations(id),
  PRIMARY KEY (id)
);

-- TEST DATA
INSERT INTO conversations (subject, created_at) VALUES ('Expression of interest.', current_timestamp);
INSERT INTO conversations (subject, created_at) VALUES ('Why hello there.', current_timestamp);

INSERT INTO messages (message_body,created_at,conversation_id) VALUES ('I think you are interesting.', current_timestamp, 1);
INSERT INTO messages (message_body,created_at,conversation_id) VALUES ('I think you are interesting as well what a coincidence.', current_timestamp, 1);
INSERT INTO messages (message_body,created_at,conversation_id) VALUES ('Hello moto! How are you.', current_timestamp, 2);