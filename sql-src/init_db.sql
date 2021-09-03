-- MySQL dump 10.13  Distrib 8.0.26, for Linux (x86_64)
--
-- Host: localhost    Database: test_db
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Facts`
--

DROP TABLE IF EXISTS `Facts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Facts` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `FactType` varchar(255) NOT NULL,
  `Content` varchar(255) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Facts`
--

LOCK TABLES `Facts` WRITE;
/*!40000 ALTER TABLE `Facts` DISABLE KEYS */;
INSERT INTO `Facts` VALUES (1,'Astronomy','The moon moves away from the earth at a rate of 3.78cm a year.'),(2,'Geography','Nemo (Latin for \'no one\') Point is the most isolated place on earth. The closest humans to anyone on Nemo Point is usually the astronauts in the ISS who are ~400km above the surface of earth.'),(3,'Travel','France is the most popular country in the world in terms of tourism.'),(4,'Astronomy','A teaspoon of a neutron star would weigh about 4 billion tons.'),(5,'Misc','Clouds can weigh over 1 million pounds.'),(6,'Misc','The coldest temperature recorded on Earth (from ground level) was −89.2°C, at the Soviet Vostok Station in 1968, on the Antarctic Plateau.'),(7,'Misc','The hottest temperature recorded on Earth is 56.7°C, recorded on 10 July 1913 at Furnace Creek Ranch, in Death Valley in the United States.'),(8,'Math','Pi (π) has been calculated to an accuracy of 62,831,853,071,750 digits, by Fachhochschule Graubünden in Switzerland over the course of 108 days.'),(9,'Math','If there are 23 people in a room there is over a 50% chance that 2 people share the same birthday. At 70 people there is a 99.9% chance.'),(10,'Physics','According to a special solution to the Einstein field equations, wormholes are a possibility and can exist.');
/*!40000 ALTER TABLE `Facts` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-09-03  3:45:57
