create table if not exists problems (
  id int primary key AUTO_INCREMENT,
  problem nvarchar(1000),
  platform nvarchar(20),
  description nvarchar(500),
  intiution text,
  link nvarchar(100)
);

insert into problems (problem, platform, description, intiution, link) 
values ("sliding window", "scaler", "slide the fing window", "learn", "leetcode.com"),
("carry forward", "leetcode", "slide the fing window with carry forward", "learn", "leetcode.com"),
("TOH", "scaler", "tower of hanoi", "learn", "scaler.com")