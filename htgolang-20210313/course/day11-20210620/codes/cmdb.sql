create database cmdb default charset utf8mb4;

use cmdb;

create table if not exists asset
(
    id         bigint primary key auto_increment comment "主键",
    ip         varchar(128) not null default "" comment "资产IP",
    addr       varchar(128) not null default "" comment "资产位置",
    created_at datetime     not null,
    updated_at datetime     not null,
    deleted_at datetime
) engine = innodb
  default charset utf8mb4;

create table if not exists module
(
    id         bigint primary key auto_increment,
    name       varchar(128) not null default "",
    created_at datetime     not null,
    updated_at datetime     not null,
    deleted_at datetime
) engine = innodb
  default charset utf8mb4;

create table if not exists role
(
    id         bigint primary key auto_increment,
    name       varchar(128) not null default "",
    created_at datetime     not null,
    updated_at datetime     not null,
    deleted_at datetime
) engine = innodb
  default charset utf8mb4;

create table if not exists rel_role_module
(
    role_id    bigint   not null default 0,
    module_id  bigint   not null default 0,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
) engine = innodb
  default charset utf8mb4;

create table if not exists user
(
    id            bigint primary key auto_increment,
    name          varchar(64)  not null default "",
    password      varchar(512) not null default "",
    birthday      date         not null,
    telephone     varchar(64)  not null default "",
    email         varchar(64)  not null default "",
    addr          varchar(128) not null default "",
    status        tinyint      not null default 0,
    role_id       bigint       not null default 0,
    department_id bigint       not null default 0,
    created_at    datetime     not null,
    updated_at    datetime     not null,
    deleted_at    datetime,
    index idx_name (name(64))
) engine = innodb
  default charset utf8mb4;

create index idx_name_password on user (`name`, `password`);

alter table user
    add column description text not null default "";

create table if not exists department
(
    id         bigint primary key auto_increment,
    name       varchar(64) not null default "",
    created_at datetime    not null,
    updated_at datetime    not null,
    deleted_at datetime
) engine = innodb
  default charset utf8mb4;

# 资产表插入一条数据
insert into asset
values (1, "192.168.1.1", "北京机房", now(), now(), null);

select *
from asset;
select *
from asset\ G;

insert into asset(ip, addr, created_at, updated_at)
values ("192.168.1.2", "北京机房", now(), now());

insert into asset(ip, addr, created_at, updated_at)
values ("192.168.1.3", "北京机房", now(), now()),
       ("192.168.1.4", "北京机房", now(), now()),
       ("192.168.1.5", "北京机房", now(), now()),
       ("192.168.1.6", "北京机房", now(), now()),
       ("192.168.1.7", "北京机房", now(), now()),
       ("192.168.1.8", "北京机房", now(), now()),
       ("192.168.1.9", "北京机房", now(), now());

# 删除数据
delete from asset;
# 等价于 delete * from asset;

# 更新数据
update asset set addr="西安机房"; #更新所有数据
update asset set addr="西安机房",updated_at=now();

# 统计数据量
select count(*) from asset;

# 去重查询
select distinct ip from asset;

# 分页
# limit 限制查询数量
# offset 偏移量
select ip from asset limit 3 offset 0;

# 排序。默认是升序
select * from asset order by ip;
select * from asset order by ip desc; # 降序
select * from asset order by ip asc; # 升序
# 多列排序
select * from asset order by ip asc,addr desc;

# 过滤条件
where
条件：bool
    布尔运算
        与 and
        或 or
        非 not
    关系运算
        <
        >
        <=
        >=
        !=
        =
字符串类型
    like
        以 XXX 开头    like 'abc%'
        以 XXX 结尾    like '%abc'
        包含 XXX       like '%abc%'

        默认使用 % 作为占位符，如果使用非 % 的占位符
        like '|abc|' ESCAPE '|'

        % 代表 0 个或者多个
        — 代表只有一个    abc 或者 adc 可以使用 like 'a_c' 进行匹配
数组
    status=1
    status=1 or status=2
    status in (1,2)
    status not in (1,2)

between and
    status >= 1 and status <= 10
    等价于
    status between 1 and 10;
    status not between 1 and 10;

NULL
    status is NULL
    status is not NULL
    deleted_at is NULL # 可以用来过滤没有删除的数据

组合查询
    and

    select * from asset where id > 13;
    select * from asset where id > 13 and addr like '西安%';
    select * from asset where id > 13 and addr like '西安%' order by id limit 2 offset 1;

聚合查询
    函数：
        count
        max
        min
        avg
        sum

    select 聚合列, 聚合函数 from 表名 group by 聚合列;
    # 统计每个 IP 出现的次数
    select ip, count(*) from asset group by ip;
    # 统计 IP 出现的次数大于 3 次的
    select ip, count(*) from asset group by ip having count(*) > 3;

# 将 id 为偶数的行的 ip 设置为 XXX
update asset set ip='XXX' where id % 2 = 0;
# 将所有 ip 的值为 XXX 的行全部删除
delete from asset where ip='XXX';


多表联查
    user department
    user role

    insert into department(name,created_at,updated_at) values("开发",now(),now()),("运维",now(),now()),("测试",now(),now());
    insert into user(name,birthday,created_at,updated_at,description,department_id) values("a",now(),now(),now(),"",1),("b",now(),now(),now(),"",1),("c",now(),now(),now(),"",2),("d",now(),now(),now(),"",0);

    select user.id, user.name,department.name
    from user,department
    where user.department_id = department.id;
    +----+------+------+
    | id | name | name |
    +----+------+------+
    | 1  | a    | 开发 |
    | 2  | b    | 开发 |
    | 3  | c    | 运维 |
    +----+------+------+

    select user.id, user.name,department.name
    from user,department
    where user.department_id = department.id and department.name='运维';
    +----+------+------+
    | id | name | name |
    +----+------+------+
    | 3  | c    | 运维 |
    +----+------+------+

    # 把多个表中的数据连接起来
    inner join => join
    left join
    right join

    select user.id,user.name,department.name from user left join department on user.department_id = department.id;
    +----+------+--------+
    | id | name | name   |
    +----+------+--------+
    | 1  | a    | 开发   |
    | 2  | b    | 开发   |
    | 3  | c    | 运维   |
    | 4  | d    | <null> |
    +----+------+--------+

    select user.id,user.name,department.name from user right join department on user.department_id = department.id;
    +--------+--------+------+
    | id     | name   | name |
    +--------+--------+------+
    | 1      | a      | 开发 |
    | 2      | b      | 开发 |
    | 3      | c      | 运维 |
    | <null> | <null> | 测试 |
    +--------+--------+------+

    # inner join 等价于 where，也等价于 join
    select user.id,user.name,department.name from user inner join department on user.department_id = department.id;
    +----+------+------+
    | id | name | name |
    +----+------+------+
    | 1  | a    | 开发 |
    | 2  | b    | 开发 |
    | 3  | c    | 运维 |
    +----+------+------+

    insert into module(name,created_at,updated_at) values("用户管理",now(),now()),("机房管理",now(),now()),("配置管理",now(),now()),("项目管理",now(),now());
    insert into role(name,created_at,updated_at) values("开发",now(),now()),("测试",now(),now()),("运维",now(),now()),("管理",now(),now());

    select * from module;
    +----+----------+---------------------+---------------------+------------+
    | id | name     | created_at          | updated_at          | deleted_at |
    +----+----------+---------------------+---------------------+------------+
    | 1  | 用户管理 | 2021-11-12 06:48:18 | 2021-11-12 06:48:18 | <null>     |
    | 2  | 机房管理 | 2021-11-12 06:48:18 | 2021-11-12 06:48:18 | <null>     |
    | 3  | 配置管理 | 2021-11-12 06:48:18 | 2021-11-12 06:48:18 | <null>     |
    | 4  | 项目管理 | 2021-11-12 06:48:18 | 2021-11-12 06:48:18 | <null>     |
    +----+----------+---------------------+---------------------+------------+

    select * from role;
    +----+------+---------------------+---------------------+------------+
    | id | name | created_at          | updated_at          | deleted_at |
    +----+------+---------------------+---------------------+------------+
    | 1  | 开发 | 2021-11-12 06:51:44 | 2021-11-12 06:51:44 | <null>     |
    | 2  | 测试 | 2021-11-12 06:51:44 | 2021-11-12 06:51:44 | <null>     |
    | 3  | 运维 | 2021-11-12 06:51:44 | 2021-11-12 06:51:44 | <null>     |
    | 4  | 管理 | 2021-11-12 06:51:44 | 2021-11-12 06:51:44 | <null>     |
    +----+------+---------------------+---------------------+------------+

    # 管理员，具备所有模块的权限
    insert into rel_role_module(role_id,module_id,created_at,updated_at) values(4,1,now(),now()),(4,2,now(),now()),(4,3,now(),now()),(4,4,now(),now());
    # 运维，具备机房管理、 配置管理、项目管理的权限
    insert into rel_role_module(role_id,module_id,created_at,updated_at) values(3,2,now(),now()),(3,3,now(),now()),(3,4,now(),now());
    # 测试，具备配置管理、项目管理的权限
    insert into rel_role_module(role_id,module_id,created_at,updated_at) values(2,3,now(),now()),(2,4,now(),now());
    # 开发，具备项目管理的权限
    insert into rel_role_module(role_id,module_id,created_at,updated_at) values(1,4,now(),now());

    select * from rel_role_module;
    +---------+-----------+---------------------+---------------------+------------+
    | role_id | module_id | created_at          | updated_at          | deleted_at |
    +---------+-----------+---------------------+---------------------+------------+
    | 4       | 1         | 2021-11-12 07:46:53 | 2021-11-12 07:46:53 | <null>     |
    | 4       | 2         | 2021-11-12 07:46:53 | 2021-11-12 07:46:53 | <null>     |
    | 4       | 3         | 2021-11-12 07:46:53 | 2021-11-12 07:46:53 | <null>     |
    | 4       | 4         | 2021-11-12 07:46:53 | 2021-11-12 07:46:53 | <null>     |
    | 3       | 2         | 2021-11-12 07:47:06 | 2021-11-12 07:47:06 | <null>     |
    | 3       | 3         | 2021-11-12 07:47:06 | 2021-11-12 07:47:06 | <null>     |
    | 3       | 4         | 2021-11-12 07:47:06 | 2021-11-12 07:47:06 | <null>     |
    | 2       | 3         | 2021-11-12 07:47:15 | 2021-11-12 07:47:15 | <null>     |
    | 2       | 4         | 2021-11-12 07:47:15 | 2021-11-12 07:47:15 | <null>     |
    | 1       | 4         | 2021-11-12 07:47:23 | 2021-11-12 07:47:23 | <null>     |
    +---------+-----------+---------------------+---------------------+------------+

    # 查看角色及其具备的权限
    select role.name,module.name from role,module,rel_role_module where role.id = rel_role_module.role_id and module.id = rel_role_module.module_id;
    # 完全等价于
    select role.name,module.name from role inner join rel_role_module on role.id = rel_role_module.role_id inner join module on module.id = rel_role_module.module_id;
    +------+----------+
    | name | name     |
    +------+----------+
    | 管理 | 用户管理 |
    | 管理 | 机房管理 |
    | 管理 | 配置管理 |
    | 管理 | 项目管理 |
    | 运维 | 机房管理 |
    | 运维 | 配置管理 |
    | 运维 | 项目管理 |
    | 测试 | 配置管理 |
    | 测试 | 项目管理 |
    | 开发 | 项目管理 |
    +------+----------+

    # 查询开发人员有哪些模块的权限
    select role.name,module.name from role,module,rel_role_module where role.id = rel_role_module.role_id and module.id = rel_role_module.module_id and role_id = 1;
    select role.name,module.name from role,module,rel_role_module where role.id = rel_role_module.role_id and module.id = rel_role_module.module_id and role.name = "开发";
    +------+----------+
    | name | name     |
    +------+----------+
    | 开发 | 项目管理 |
    +------+----------+

    update user set role_id = id;
    select id,name,role_id from user;
    +----+------+---------+
    | id | name | role_id |
    +----+------+---------+
    | 1  | a    | 1       |
    | 2  | b    | 2       |
    | 3  | c    | 3       |
    | 4  | d    | 4       |
    +----+------+---------+

    select user.name,role.name,module.name from user,role,module,rel_role_module where user.role_id = role.id and role.id = rel_role_module.role_id and module.id = rel_role_module.module_id;
    +------+------+----------+
    | name | name | name     |
    +------+------+----------+
    | d    | 管理 | 用户管理 |
    | d    | 管理 | 机房管理 |
    | d    | 管理 | 配置管理 |
    | d    | 管理 | 项目管理 |
    | c    | 运维 | 机房管理 |
    | c    | 运维 | 配置管理 |
    | c    | 运维 | 项目管理 |
    | b    | 测试 | 配置管理 |
    | b    | 测试 | 项目管理 |
    | a    | 开发 | 项目管理 |
    +------+------+----------+

    select user.name,role.name,module.name from user,role,module,rel_role_module where user.role_id = role.id and role.id = rel_role_module.role_id and module.id = rel_role_module.module_id and user.name = 'a';
    +------+------+----------+
    | name | name | name     |
    +------+------+----------+
    | a    | 开发 | 项目管理 |
    +------+------+----------+

子查询
    select where select

    想要查询部门名称中包含“开发”的所有人
    select * from user where user.department_id in (
    select id from department where name like '%开发%'
    );
    +----+------+----------+------------+-----------+-------+------+--------+---------+---------------+---------------------+---------------------+------------+-------------+
    | id | name | password | birthday   | telephone | email | addr | status | role_id | department_id | created_at          | updated_at          | deleted_at | description |
    +----+------+----------+------------+-----------+-------+------+--------+---------+---------------+---------------------+---------------------+------------+-------------+
    | 1  | a    |          | 2021-11-11 |           |       |      | 0      | 1       | 1             | 2021-11-11 09:41:08 | 2021-11-11 09:41:08 | <null>     |             |
    | 2  | b    |          | 2021-11-11 |           |       |      | 0      | 2       | 1             | 2021-11-11 09:41:08 | 2021-11-11 09:41:08 | <null>     |             |
    +----+------+----------+------------+-----------+-------+------+--------+---------+---------------+---------------------+---------------------+------------+-------------+

别名 - 表名的别名
    # 查看角色及其具备的权限
    select role.name,module.name from role,module,rel_role_module where role.id = rel_role_module.role_id and module.id = rel_role_module.module_id;
    # 等价于（有 as）
    select r.name,m.name from role as r,module as m,rel_role_module as rel where r.id = rel.role_id and m.id = rel.module_id;
    # 等价于（没有 as 也可以）
    select r.name,m.name from role r,module m,rel_role_module rel where r.id = rel.role_id and m.id = rel.module_id;
    +------+----------+
    | name | name     |
    +------+----------+
    | 管理 | 用户管理 |
    | 管理 | 机房管理 |
    | 管理 | 配置管理 |
    | 管理 | 项目管理 |
    | 运维 | 机房管理 |
    | 运维 | 配置管理 |
    | 运维 | 项目管理 |
    | 测试 | 配置管理 |
    | 测试 | 项目管理 |
    | 开发 | 项目管理 |
    +------+----------+

别名 - 列名的别名（查询结果中显示）
    select r.name as roleName,m.name as moduleName from role as r,module as m,rel_role_module as rel where r.id = rel.role_id and m.id = rel.module_id;
    +----------+------------+
    | roleName | moduleName |
    +----------+------------+
    | 管理     | 用户管理   |
    | 管理     | 机房管理   |
    | 管理     | 配置管理   |
    | 管理     | 项目管理   |
    | 运维     | 机房管理   |
    | 运维     | 配置管理   |
    | 运维     | 项目管理   |
    | 测试     | 配置管理   |
    | 测试     | 项目管理   |
    | 开发     | 项目管理   |
    +----------+------------+

集合
    并集 union（去重）/union all（不去重）
    交集 intersect
    差集 except

    并集
        select name from role union select name from department;
        +------+
        | name |
        +------+
        | 开发 |
        | 测试 |
        | 运维 |
        | 管理 |
        +------+

        select name from role union all select name from department;
        +------+
        | name |
        +------+
        | 开发 |
        | 测试 |
        | 运维 |
        | 管理 |
        | 开发 |
        | 运维 |
        | 测试 |
        +------+

    交集
        select name from role intersect select name from department;
        +------+
        | name |
        +------+
        | 开发 |
        | 测试 |
        | 运维 |
        +------+

    差集
        select name from role except select name from department;
        +------+
        | name |
        +------+
        | 管理 |
        +------+

函数
    字符串：
        upper()
            select upper('abc');
            +--------------+
            | upper('abc') |
            +--------------+
            | ABC          |
            +--------------+
        lower()
            select lower('ABC');
            +--------------+
            | lower('ABC') |
            +--------------+
            | abc          |
            +--------------+
        trim()
        to_base64()
            select to_base64('abc');
            +------------------+
            | to_base64('abc') |
            +------------------+
            | YWJj             |
            +------------------+

        from_base64()
            select from_base64('abc');
            +--------------------+
            | from_base64('abc') |
            +--------------------+
            | <null>             |
            +--------------------+
    时间：
        now()
            select now();
            +---------------------+
            | now()               |
            +---------------------+
            | 2021-11-12 10:45:30 |
            +---------------------+

        date_format()
            select date_format(now(),'%y');
            +-------------------------+
            | date_format(now(),'%y') |
            +-------------------------+
            | 21                      |
            +-------------------------+

            select date_format(now(),'%Y');
            +-------------------------+
            | date_format(now(),'%Y') |
            +-------------------------+
            | 2021                    |
            +-------------------------+

            select date_format(now(),'%Y-%m-%d %H:%m:%s');
            +----------------------------------------+
            | date_format(now(),'%Y-%m-%d %H:%m:%s') |
            +----------------------------------------+
            | 2021-11-12 10:11:17                    |
            +----------------------------------------+

    编解码：
        md5
            select md5('abc');
            +----------------------------------+
            | md5('abc')                       |
            +----------------------------------+
            | 900150983cd24fb0d6963f7d28e17f72 |
            +----------------------------------+

        sha1
            select sha1('abc');
            +------------------------------------------+
            | sha1('abc')                              |
            +------------------------------------------+
            | a9993e364706816aba3e25717850c26c9cd0d89d |
            +------------------------------------------+

连接字符串：
    select "abc" "def";
    +--------+
    | abc    |
    +--------+
    | abcdef |
    +--------+

计算：
    select 1+2;
    +-----+
    | 1+2 |
    +-----+
    | 3   |
    +-----+

alter table alarm add column status int not null default 0;