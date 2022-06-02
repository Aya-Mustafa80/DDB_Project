select user,host from mysql.user;
create user 'RemoteProject'@'%' identified with mysql_native_password by '12345678';
FLUSH privileges;
grant all on *.* to 'RemoteProject' @'%';



