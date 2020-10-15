# Microservices_exercise_GoAndNode

Server 1 
From golang to akses api BMKG

Server 2 
From Node JS to DB

main.go Controller access in postman

before start  sample main > Server 2
must npm i and createdb db_nodejs and table 
CREATE TABLE `customer_bank` (
  `customer_id` int(11) NOT NULL,
  `nama` varchar(50) NOT NULL,
  `alamat` varchar(200) NOT NULL,
  `kode_pos` char(5) DEFAULT NULL,
  `no_hp` varchar(15) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL
)
and Insert one data

run server 2 and run main.go
