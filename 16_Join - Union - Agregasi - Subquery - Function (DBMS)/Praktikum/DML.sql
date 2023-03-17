-- //Insert 5 operators pada table operators.
insert into operator values('1','TRI', now(), now());
insert into operator values('2','XL AXIATA', now(), now());
insert into operator values('3','TELKOMSEL', now(), now());
insert into operator values('4','SMARTFREN', now(), now());
insert into operator values('5','INDOSAT', now(), now());

-- Insert 3 product type.
insert into product_type values ('1','Paket Murah', now(), now());
insert into product_type values ('2','Paket Mahal', now(), now());
insert into product_type values ('3','Paket Promo', now(), now());

-- Insert 2 product dengan product type id = 1, dan operators id = 3.
insert into product(product_type_id, operator_id,code,nama,status,created_at,update_at) 
values('1','3','0001','TELKOM MURAH 1','1',now(),now());
insert into product (product_type_id, operator_id,code,nama,status,created_at,update_at)
values('1','3','0002','TELKOM MURAH 2','1',now(),now());

-- Insert 3 product dengan product type id = 2, dan operators id = 1.
insert into product(product_type_id, operator_id,code,nama,status,created_at,update_at) 
value('2','1','0003','TRI MAHAL 1','1',now(),now());
insert into product(product_type_id, operator_id,code,nama,status,created_at,update_at) 
value('2','1','0004','TRI MAHAL 2','1',now(),now());
insert into product(product_type_id, operator_id,code,nama,status,created_at,update_at) 
value('2','1','0005','TRI MAHAL 3','1',now(),now());

-- Insert 3 product dengan product type id = 3, dan operators id = 4.
insert into product(product_type_id, operator_id,code,nama,status,created_at,update_at) 
value('3','4','0006','SMART PROMO 1','1',now(),now());
insert into product(product_type_id, operator_id,code,nama,status,created_at,update_at) 
value('3','4','0007','SMART PROMO 2','1',now(),now());
insert into product(product_type_id, operator_id,code,nama,status,created_at,update_at) 
value('3','4','0008','SMART PROMO 3','1',now(),now());

-- Insert product description pada setiap product.
insert into product_description (description, product_id,created_at,update_at)
values ('Berisikan Kuota 1GB','1',now(),now());
insert into product_description (description, product_id,created_at,update_at)
values ('Berisikan Kuota 10GB','2',now(),now());
insert into product_description (description, product_id,created_at,update_at)
values ('Berisikan Kuota 20GB','3',now(),now());
insert into product_description (description, product_id,created_at,update_at)
values ('Berisikan Kuota 30GB','4',now(),now());
insert into product_description (description, product_id,created_at,update_at)
values ('Berisikan Kuota 40GB','5',now(),now());
insert into product_description (description, product_id,created_at,update_at)
values ('Berisikan Kuota 50GB','6',now(),now());
insert into product_description (description, product_id,created_at,update_at)
values ('Berisikan Kuota 60GB','7',now(),now());
insert into product_description (description, product_id,created_at,update_at)
values ('Berisikan Kuota 100GB','8',now(),now());

-- Insert 3 payment methods.
insert into payment_method (nama, status,created_at,update_at) 
values('GOPAY','5',now(),now());
insert into payment_method (nama, status,created_at,update_at) 
values('SHOPPEPAY','5',now(),now());
insert into payment_method (nama, status,created_at,update_at) 
values('OVO','5',now(),now());

-- Insert 5 user pada tabel user.
insert into users (nama,status,dob,gender,created_at,update_at)
values('Subaru','1','2002-03-29','M',now(),now());
insert into users (nama,status,dob,gender,created_at,update_at)
values('Putri','1','2002-03-01','W',now(),now());
insert into users (nama,status,dob,gender,created_at,update_at)
values('Riski','1','2002-03-02','M',now(),now());
insert into users (nama,status,dob,gender,created_at,update_at)
values('Aulia','1','2002-03-06','W',now(),now());
insert into users (nama,status,dob,gender,created_at,update_at)
values('Andika','1','2002-03-19','M',now(),now());

-- user subaru --  
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('1','1','OK','1','20000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('1','2','OK','2','40000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('1','3','OK','3','60000',now(),now());

-- user putri --  
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('2','1','OK','1','20000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('2','2','OK','2','40000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('2','3','OK','3','60000',now(),now());

-- user riski --  
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('3','1','OK','1','20000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('3','2','OK','2','40000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('3','3','OK','3','60000',now(),now());

-- user aulia --  
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('4','1','OK','1','20000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('4','2','OK','2','40000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('4','3','OK','3','60000',now(),now());

-- user putri --  
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('5','1','OK','1','20000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('5','2','OK','2','40000',now(),now());
insert into transaction (users_id,payment_method_id,status,total_quanty,total_price,created_at,update_at)
values('5','3','OK','3','60000',now(),now());

-- insert product --

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('1','1','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('1','2','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('1','3','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('2','4','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('2','5','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('2','6','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('3','7','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('3','8','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('3','1','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('4','2','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('4','3','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('4','4','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('5','5','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('5','6','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('5','7','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('6','8','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('6','1','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('6','2','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('7','3','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('7','4','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('7','5','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('8','6','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('8','7','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('8','8','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('9','1','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('9','2','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('9','3','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('10','4','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('10','5','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('10','6','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('11','7','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('11','8','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('11','1','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('12','2','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('12','3','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('12','4','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('13','5','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('13','6','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('13','7','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('14','8','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('14','1','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('14','2','OK','3','60000',now(),now());

insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('15','3','OK','1','20000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('15','4','OK','2','40000',now(),now());
insert into transaction_details (transaction_id,product_id,status,quanty,price,created_at,update_at) 
values('15','5','OK','3','60000',now(),now());


-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M. -- 

select * from users where gender='M';



-- Tampilkan product dengan id = 3. --

select * from product where id='3';



-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’. --

select * from users where created_at > date_sub(now(), interval 1 week) AND nama LIKE '%a%';



-- Hitung jumlah user / pelanggan dengan status gender Perempuan. --

 select sum(gender='W') from users;
 
 
 
 -- Tampilkan data pelanggan dengan urutan sesuai nama abjad --
 
 select * from users order by nama ASC;


-- Tampilkan 5 data pada data product -- 

select * from product limit 5;



-- Ubah data product id 1 dengan nama ‘product dummy’. --

update product set nama='product dummy' where id='1'; 




-- Update qty = 3 pada transaction detail dengan product id = 1. --

update transaction_details set quanty='3' where product_id='1'; 




-- Delete data pada tabel product dengan id = 1. --
-- Delete pada pada tabel product dengan product type id 1. -- 

update product set deleted_at = now() where id = 1 OR product_type_id =1;






-- Gabungkan data transaksi dari user id 1 dan user id 2. --

 select * from transaction where users_id in(1,2);
 
 
 
 -- Tampilkan jumlah harga transaksi user id 1. -- 
 
 select users_id, Sum(total_price) as total_price from transaction where users_id = 1;


-- Tampilkan semua field table product dan field name table product type yang saling berhubungan. -- 

select count(*) from transaction t
join transaction_details td on t.id = td.transaction_id
join product p on p.id = td.product_id
where p.product_type_id = 2;


-- Tampilkan semua field table transaction, field name table product dan field name table user. --

select t.*, p.nama as product, u.nama as users from transaction_details td
join transaction t on t.id = td.transaction_id
join users u on u.id = t.users_id
join product p on p.id = td.product_id
where p.deleted_at is null
 
 
 -- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud. --
 -- delete row dorign key -- 
 delimiter $$
 create trigger delete_transaction_details
 before delete on transaction for each row
 begin
 declare trans_id int;
 set trans_id = old.id;
 delete from transaction_details where transaction_id = trans_id;
 end $$
 
 delete from transaction where id = 4;
 
 
 
 -- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus. --
 
 delimiter $$
 create trigger update_transaction_id
 after delete on transaction_details for each row
 begin
 declare trans_id int;
 set trans_id = old.transaction_id;
 update transaction set total_quanty = (select sum(quanty) from transaction_details where transaction_id = trans_id)
 where id = trans_id;
 end $$
 
 select * from transaction;
 
 select * from transaction_details where transaction_id = 15;
 
 delete from transaction_details where transaction_id = 15 AND product_id = 4;
 
 
 
 -- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query. -- 
 
 select * from product where id not in(
 select product_id from transaction_details
 );
 
select * from product where id in (
select p.id from product p
left outer join transaction_details td on p.id = td.product_id
where td.transaction_id is null
);