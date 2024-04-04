create table if not exists problems (
  id int primary key AUTO_INCREMENT,
  problem text,
  platform text,
  description text,
  intiution text,
  link text
);

insert into problems (problem, platform, description, intiution, link) 
values ("sliding window", "scaler", "slide the fing window", "learn", "https://leetcode.com"),
("carry forward", "leetcode", "slide the fing window with carry forward", "learn", "https://leetcode.com"),
("TOH", "scaler", "tower of hanoi", "learn", "https://scaler.com")