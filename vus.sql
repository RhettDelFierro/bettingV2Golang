CREATE TABLE vus (
	id INT NOT NULL,
	team_id INT NOT NULL,
	opponent_id INT,
	period TEXT,
	season TEXT,
	orbc TEXT,
	drbc INT,
	sast INT,
	ftast INT,
	cfgm INT,
	cfga INT,
	ufgm INT,
	ufga INT,
	dfgm INT,
	dfga INT,
	PRIMARY KEY(id, team_id)
);