create table tasks (
	id char(36) not null,
	title varchar(50) not null,
	description varchar(100) null,
	done bool default false,
	created_at timestamp default current_timestamp,
	primary key(id)
);
