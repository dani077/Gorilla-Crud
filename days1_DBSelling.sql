-- MySQL dump 10.15  Distrib 10.0.34-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: DBSelling
-- ------------------------------------------------------
-- Server version	10.0.34-MariaDB-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tblItem`
--

DROP TABLE IF EXISTS `tblItem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tblItem` (
  `tblItemID` int(11) NOT NULL AUTO_INCREMENT,
  `ItemCode` varchar(13) NOT NULL,
  `ItemName` varchar(30) NOT NULL,
  `BuyingPrice` decimal(18,0) NOT NULL,
  `SellingPrice` decimal(18,0) NOT NULL,
  `ItemAmount` int(11) NOT NULL,
  `Pieces` varchar(15) NOT NULL,
  PRIMARY KEY (`tblItemID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tblItem`
--

LOCK TABLES `tblItem` WRITE;
/*!40000 ALTER TABLE `tblItem` DISABLE KEYS */;
INSERT INTO `tblItem` VALUES (1,'IC001','Laptop',1200,1500,8,'Unit'),(2,'IC002','Headset',200,230,7,'Unit'),(3,'IC003','Mouse',260,300,8,'Unit'),(4,'IC004','PenDisk 32Gb',150,200,5,'PCS'),(5,'IC005','Cable HDMI',100,120,7,'PCS'),(6,'IC006','Headset Extra Bass',350,400,9,'Unit');
/*!40000 ALTER TABLE `tblItem` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tblOfficer`
--

DROP TABLE IF EXISTS `tblOfficer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tblOfficer` (
  `tblOfficerID` int(11) NOT NULL AUTO_INCREMENT,
  `OfficerCode` varchar(5) NOT NULL,
  `OfficerName` varchar(30) NOT NULL,
  `OfficerPassword` varchar(10) NOT NULL,
  `OfficerStatus` varchar(15) NOT NULL,
  PRIMARY KEY (`tblOfficerID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tblOfficer`
--

LOCK TABLES `tblOfficer` WRITE;
/*!40000 ALTER TABLE `tblOfficer` DISABLE KEYS */;
INSERT INTO `tblOfficer` VALUES (1,'OC001','Fullan Banyu','Fullan123','Active'),(2,'OC002','Heru Farouk','Heru123','Active'),(3,'OC003','Ingrid Wijaya','Ingrid123','Active'),(4,'OC004','Budi Cahaya','Budi123','Active'),(5,'OC005','Alex Timor','Alex123','Non Active'),(6,'OC006','Anton Kurniawan','Anton123','Active');
/*!40000 ALTER TABLE `tblOfficer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tblSelling`
--

DROP TABLE IF EXISTS `tblSelling`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tblSelling` (
  `tblSellingID` int(11) NOT NULL AUTO_INCREMENT,
  `Invoice` varchar(5) NOT NULL,
  `InvoiceDate` date NOT NULL,
  `Item` int(11) NOT NULL,
  `Total` decimal(18,0) NOT NULL,
  `Paid` decimal(18,0) NOT NULL,
  `Pengembalian` decimal(18,0) NOT NULL,
  `OfficerCode` varchar(5) NOT NULL,
  PRIMARY KEY (`tblSellingID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tblSelling`
--

LOCK TABLES `tblSelling` WRITE;
/*!40000 ALTER TABLE `tblSelling` DISABLE KEYS */;
INSERT INTO `tblSelling` VALUES (1,'Iv001','2017-01-02',3,2900,3000,100,'OC002'),(2,'Iv002','2017-01-02',3,1960,5000,3040,'OC001'),(3,'Iv003','2017-01-03',2,4000,10000,6000,'OC002'),(4,'Iv004','2017-01-05',1,300,300,0,'OC002'),(5,'Iv005','2017-01-10',3,5500,10000,4500,'OC003'),(6,'Iv006','2017-01-11',2,6080,10000,3020,'OC003');
/*!40000 ALTER TABLE `tblSelling` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tblSellingDetail`
--

DROP TABLE IF EXISTS `tblSellingDetail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tblSellingDetail` (
  `tblSellingDetailID` int(11) NOT NULL AUTO_INCREMENT,
  `Invoice` varchar(5) NOT NULL,
  `ItemCode` varchar(30) NOT NULL,
  `ItemName` varchar(10) NOT NULL,
  `ItemPrice` varchar(15) NOT NULL,
  `SubTotal` decimal(18,0) NOT NULL,
  PRIMARY KEY (`tblSellingDetailID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tblSellingDetail`
--

LOCK TABLES `tblSellingDetail` WRITE;
/*!40000 ALTER TABLE `tblSellingDetail` DISABLE KEYS */;
INSERT INTO `tblSellingDetail` VALUES (1,'Iv001','IC001','Laptop','1500',1500),(2,'Iv001','IC003','Mouse','300',600),(3,'Iv001','IC004','PenDisk 32','200',800),(4,'Iv002','IC002','Headset','230',460),(5,'Iv002','IC003','Mouse','300',300),(6,'Iv002','IC006','Headset Ba','400',1200);
/*!40000 ALTER TABLE `tblSellingDetail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `viewItem`
--

DROP TABLE IF EXISTS `viewItem`;
/*!50001 DROP VIEW IF EXISTS `viewItem`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `viewItem` (
  `ItemName` tinyint NOT NULL,
  `InvoiceDate` tinyint NOT NULL,
  `OfficerName` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Final view structure for view `viewItem`
--

/*!50001 DROP TABLE IF EXISTS `viewItem`*/;
/*!50001 DROP VIEW IF EXISTS `viewItem`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `viewItem` AS select `tblItem`.`ItemName` AS `ItemName`,`tblSelling`.`InvoiceDate` AS `InvoiceDate`,`tblOfficer`.`OfficerName` AS `OfficerName` from ((`tblItem` join `tblSelling`) join `tblOfficer`) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-03-14  7:47:49
