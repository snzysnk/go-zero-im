CREATE TABLE user
(
    # PRIMARY KEY 要移出去，不然会报错：table user: missing primary key
    id   int unsigned auto_increment,
    name varchar(16) not null default '' comment '名称',
    PRIMARY KEY(id)
) COMMENT '用户表';