//membuat database alta_online_shop
CREATE DATABASE alta_online_shop;

//masuk kedalam database alta_online_shop
use alta_online_shop;

//membuat tabel users
create table users (
id int(11) AUTO_INCREMENT PRIMARY KEY,
nama varchar(100),
alamat varchar(100),
tanggal_lahir date,
status_user int,
gender varchar(10),
created_at timestamp,
update_at timestamp);

//membuat tabel payment_method
create table payment_method (
id int(11) AUTO_INCREMENT PRIMARY KEY,
nama varchar(100),
status smallint(20),
created_at timestamp,
update_at timestamp);

//membuat tabel transaction
create table transaction (
id int(11) AUTO_INCREMENT PRIMARY KEY,
user_id int(11),
status varchar(10),
total_quanty int(11),
total_price int,
created_at timestamp,
update_at timestamp,
FOREIGN KEY (user_id) REFERENCES users(id));

//membuat tabel operator
create table operator (
id int(11) AUTO_INCREMENT PRIMARY KEY,
nama varchar(255),
created_at timestamp,
update_at timestamp);

//membuat tabel product_type
create table product_type (
id int(11) AUTO_INCREMENT PRIMARY KEY,
nama varchar(100),
created_at timestamp,
updated_at timestamp);

//membuat tabel product
create table product (
id int(11) AUTO_INCREMENT PRIMARY KEY,
product_type_id int(11),
operator_id int(11),
code varchar(100),
nama varchar(100),
status smallint,
created_at timestamp,
update_at timestamp,
FOREIGN KEY (product_type_id) REFERENCES product_type(id),
FOREIGN KEY (operator_id) REFERENCES operator(id));

//membuat tabel transaction_details
create table transaction_details (
transaction_id int(11),
product_id int(11),
status varchar(10),
quanty int(11),
price int,
created_at timestamp,
update_at timestamp,
FOREIGN KEY (transaction_id) REFERENCES transaction(id),
FOREIGN KEY (product_id) REFERENCES product(id));

create table product_description (
id int(11) AUTO_INCREMENT PRIMARY KEY,
description varchar(255),
product_id int,
craeted_at timestamp,
update_at timestamp,
FOREIGN KEY (product_id) REFERENCES product(id));

//membuat tabel kurir
create table kurir (
id int(11) AUTO_INCREMENT PRIMARY KEY,
nama varchar(100),
created_at timestamp,
updated_at timestamp);

//menambahkan kolom ongkos_dasar pada tabel kurir
alter table kurir ADD COLUMN ongkos_dasar int

//mengubah nama tabel kuru menjadi shipping
alter table kurir RENAME shipping;

//menghapus tabel shipping
drop table shipping;

//membuat tabel payment_method_description one to one dengan tabel payment_method
create table payment_method_description (
id int AUTO_INCREMENT PRIMARY KEY,
payment_method_id int(11),
description varchar(100),
created_at timestamp not null default current_timestamp,
updated_at timestamp not null default current_timestamp,
FOREIGN KEY (payment_method_id) REFERENCES payment_method(id));

//membuat tabel user_address one to many dengan tabel users
create table alamat (
id int AUTO_INCREMENT PRIMARY KEY,
alamat varchar(255),
user_id int,
created_at timestamp not null default current_timestamp,
update_at timestamp not null default current_timestamp,
FOREIGN KEY (user_id) REFERENCES users(id));

//membuat tabel user_payment_method_detials many to many dengan tabel users dan payment_method
create table user_payment_method_detials (
user_id int(11),
payment_method_id int(11),
created_at timestamp not null default current_timestamp,
update_at timestamp not null default current_timestamp,
FOREIGN KEY (user_id) REFERENCES users(id),
FOREIGN KEY (payment_method_id) REFERENCES payment_method(id));



