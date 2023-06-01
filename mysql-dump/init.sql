-- TODO - log warnings and errors

create table if not exists gorm.database_sessions (
	session_id INT unsigned auto_increment primary key, 
	session_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

insert into gorm.database_sessions () values ();

