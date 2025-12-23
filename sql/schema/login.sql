CREATE TABLE login_log (
	id TEXT PRIMARY KEY,
	user_uuid TEXT DEFAULT "",
	phone TEXT DEFAULT "",
	timestamp INT DEFAULT 0,
	error_cause TEXT DEFAULT "",
	action TEXT DEFAULT "",
	error BOOLEAN DEFAULT FALSE
);


CREATE TABLE otp_log (
	id TEXT PRIMARY KEY,
	user_uuid TEXT DEFAULT "",
	timestampt INT DEFAULT 0,
	error_cause TEXT DEFAULT "",
	action TEXT DEFAULT "",
	error BOOLEAN DEFAULT FALSE,
	otp_id TEXT DEFAULT ""
);


CREATE TABLE vehicle_log (
	id TEXT PRIMARY KEY,
	nopol TEXT DEFAULT "",
	noka TEXT DEFAULT "",
	nosing TEXT DEFAULT "",
	cabang TEXT DEFAULT ""
);


CREATE TABLE action_log (
	id TEXT PRIMARY KEY,
	user_uuid TEXT,
	phone TEXT,
	action TEXT DEFAULT "",
	action_id TEXT DEFAULT "",
	error_cause TEXT DEFAULT ""
);
