create table titles(
    id int(3) unsigned not null primary key auto_increment,
    title nvarchar(10) not null
) default charset = utf8 auto_increment = 1;

create table ts(
    id int(3) unsigned not null primary key,
    t datetime not null default now(),
    foreign key(id) references titles(id)
) default charset = utf8;

create table contents(
    id int(3) unsigned not null primary key,
    content nvarchar(10) not null,
    foreign key(id) references titles(id)
) default charset = utf8;