
CREATE TABLE IF NOT EXISTS threads (
  thread_id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(50) NOT NULL,
  description varchar(250) DEFAULT NULL,
  PRIMARY KEY (thread_id)
);

CREATE TABLE IF NOT EXISTS replies (
  replie_id int(11) NOT NULL AUTO_INCREMENT,
  threadParent_id int(11),
  content varchar(250) DEFAULT NULL,
  PRIMARY KEY (replie_id),
  FOREIGN KEY (threadParent_id) REFERENCES threads(thread_id) ON DELETE CASCADE
);
