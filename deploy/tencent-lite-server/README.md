## mysql

```shell
docker run -d --name mysql \
-p 3306:3306 \
-e TZ=Asia/Shanghai \
-e MYSQL_ROOT_PASSWORD=123456 \
--privileged=true \
--restart=always \
--user=root \
mysql/mysql-server:8.0.28-1.2.7-server \
--default-authentication-plugin=mysql_native_password \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_general_ci \
--explicit_defaults_for_timestamp=true \
--lower_case_table_names=1

mysql -uroot -p123456
# 执行mysql命令
CREATE USER 'root'@'%' IDENTIFIED BY 'root'; 
GRANT ALL ON . TO 'root'@'%'; 
flush privileges; 
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456'; 
flush privileges;
grant all privileges on *.* to 'root'@'%' with grant option;
flush 
```