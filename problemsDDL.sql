create table problems (
  id int primary key AUTO_INCREMENT,
  problem nvarchar(1000),
  platform nvarchar(20),
  description nvarchar(500),
  intiution text,
  link nvarchar(100)
);