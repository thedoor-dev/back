# BACK

![]()

```SQL
CREATE TABLE `posts` (
	`pid` BIGINT(11) NOT NULL,
	`abstract` CHAR(100) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	`article` TEXT NOT NULL COLLATE 'utf8mb4_general_ci',
	`ctime` DATETIME NULL DEFAULT current_timestamp(),
	`public` TINYINT(1) NULL DEFAULT '0',
	PRIMARY KEY (`pid`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE tags(
    `tid` BIGINT(11) NOT NULL,
    `pid` BIGINT(11) NOT NULL,
    `name` CHAR()
	PRIMARY KEY (`pid`) USING BTREE
)
```
