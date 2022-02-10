mysql -u root << EOF
USE partyholic

drop table users;
drop table parties;
drop table cancelled_parties;
drop table addresses;

EOF