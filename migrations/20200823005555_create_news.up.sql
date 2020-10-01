CREATE TABLE IF NOT EXISTS news (
		id int auto_increment primary key,
		title varchar(255),
		text TEXT,
		news_date varchar(255)
	) DEFAULT CHARACTER SET UTF8