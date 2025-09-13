-- MySQL dump 10.13  Distrib 9.4.0, for macos15 (arm64)
--
-- Host: 127.0.0.1    Database: notes_app
-- ------------------------------------------------------
-- Server version	9.4.0

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
-- Table structure for table `notes`
--

DROP TABLE IF EXISTS `notes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `notes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `title` varchar(255) NOT NULL,
  `body` text NOT NULL,
  `tags` json DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `notes_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `notes`
--

LOCK TABLES `notes` WRITE;
/*!40000 ALTER TABLE `notes` DISABLE KEYS */;
INSERT INTO `notes` VALUES (1,15,'Note Title 1','This is the body of note 1.','[\"tag1\", \"tag2\", \"note1\"]','2024-06-01 06:47:14','2025-02-18 15:01:36'),(2,2,'Note Title 2','This is the body of note 2.','[\"tag1\", \"tag2\", \"note2\"]','2024-10-08 13:32:07','2024-12-09 07:30:57'),(3,13,'Note Title 3','This is the body of note 3.','[\"tag1\", \"tag2\", \"note3\"]','2024-01-26 14:13:03','2024-08-12 02:12:38'),(4,11,'Note Title 4','This is the body of note 4.','[\"tag1\", \"tag2\", \"note4\"]','2023-11-12 00:31:21','2024-02-14 03:16:11'),(5,15,'Note Title 5','This is the body of note 5.','[\"tag1\", \"tag2\", \"note5\"]','2024-09-23 01:26:53','2025-02-13 07:11:15'),(6,9,'Note Title 6','This is the body of note 6.','[\"tag1\", \"tag2\", \"note6\"]','2025-04-03 11:14:55','2023-11-25 15:33:55'),(7,14,'Note Title 7','This is the body of note 7.','[\"tag1\", \"tag2\", \"note7\"]','2024-10-01 13:19:08','2024-04-17 15:07:53'),(8,11,'Note Title 8','This is the body of note 8.','[\"tag1\", \"tag2\", \"note8\"]','2025-05-27 06:35:52','2024-01-23 09:26:44'),(9,14,'Note Title 9','This is the body of note 9.','[\"tag1\", \"tag2\", \"note9\"]','2024-11-27 07:19:33','2024-07-11 02:10:15'),(10,4,'Note Title 10','This is the body of note 10.','[\"tag1\", \"tag2\", \"note10\"]','2024-08-02 00:10:50','2025-03-21 02:55:11'),(11,10,'Note Title 11','This is the body of note 11.','[\"tag1\", \"tag2\", \"note11\"]','2024-02-13 23:54:39','2024-08-07 14:21:42'),(12,3,'Note Title 12','This is the body of note 12.','[\"tag1\", \"tag2\", \"note12\"]','2024-07-19 11:27:21','2025-03-05 23:54:07'),(13,1,'Note Title 13','This is the body of note 13.','[\"tag1\", \"tag2\", \"note13\"]','2025-02-25 10:59:19','2025-07-19 09:22:22'),(14,12,'Note Title 14','This is the body of note 14.','[\"tag1\", \"tag2\", \"note14\"]','2025-03-30 23:39:31','2025-03-19 08:41:56'),(15,2,'Note Title 15','This is the body of note 15.','[\"tag1\", \"tag2\", \"note15\"]','2023-10-06 18:58:55','2024-07-22 16:46:54');
/*!40000 ALTER TABLE `notes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (3,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_table`
--

DROP TABLE IF EXISTS `test_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `test_table` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_table`
--

LOCK TABLES `test_table` WRITE;
/*!40000 ALTER TABLE `test_table` DISABLE KEYS */;
INSERT INTO `test_table` VALUES (1,'TestName1','This is description for test record 1.','2025-03-25 20:25:29','2023-11-22 02:45:50'),(2,'TestName2','This is description for test record 2.','2024-09-05 03:20:29','2025-02-26 18:53:37'),(3,'TestName3','This is description for test record 3.','2023-12-15 03:49:14','2024-02-24 21:32:27'),(4,'TestName4','This is description for test record 4.','2025-02-12 17:18:35','2024-06-16 17:03:31'),(5,'TestName5','This is description for test record 5.','2024-05-15 02:06:47','2023-12-24 00:08:24'),(6,'TestName6','This is description for test record 6.','2025-05-21 01:25:21','2024-11-28 03:07:57'),(7,'TestName7','This is description for test record 7.','2023-11-08 08:57:03','2025-08-31 06:32:11'),(8,'TestName8','This is description for test record 8.','2024-05-07 20:36:19','2025-09-11 08:32:09'),(9,'TestName9','This is description for test record 9.','2024-10-26 13:22:10','2024-04-26 15:31:49'),(10,'TestName10','This is description for test record 10.','2025-04-11 22:06:35','2024-04-12 12:57:14'),(11,'TestName11','This is description for test record 11.','2025-06-25 18:20:42','2025-05-07 22:22:44'),(12,'TestName12','This is description for test record 12.','2025-07-01 08:54:36','2025-02-08 18:59:19'),(13,'TestName13','This is description for test record 13.','2024-11-11 14:50:23','2025-08-30 12:25:41'),(14,'TestName14','This is description for test record 14.','2025-02-28 20:17:28','2025-04-05 11:59:04'),(15,'TestName15','This is description for test record 15.','2025-03-17 11:44:45','2024-10-10 03:46:51');
/*!40000 ALTER TABLE `test_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  KEY `username_2` (`username`,`email`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'user1','user1@example.com','hashed_password_1','2025-09-01 20:21:41','2024-03-20 13:49:28'),(2,'user2','user2@example.com','hashed_password_2','2024-01-13 03:45:35','2024-02-05 05:37:31'),(3,'user3','user3@example.com','hashed_password_3','2025-04-03 13:06:54','2025-03-03 03:15:08'),(4,'user4','user4@example.com','hashed_password_4','2025-07-26 17:36:51','2024-06-26 13:02:44'),(5,'user5','user5@example.com','hashed_password_5','2024-07-04 14:42:33','2023-09-29 05:47:12'),(6,'user6','user6@example.com','hashed_password_6','2024-12-19 12:18:36','2024-02-28 00:25:16'),(7,'user7','user7@example.com','hashed_password_7','2025-04-12 03:57:05','2024-09-24 16:28:56'),(8,'user8','user8@example.com','hashed_password_8','2024-12-13 06:35:48','2024-11-15 21:53:53'),(9,'user9','user9@example.com','hashed_password_9','2025-08-25 07:07:59','2025-04-29 20:17:44'),(10,'user10','user10@example.com','hashed_password_10','2025-04-28 20:42:56','2025-04-07 14:33:27'),(11,'user11','user11@example.com','hashed_password_11','2024-05-27 11:51:22','2024-02-24 23:16:22'),(12,'user12','user12@example.com','hashed_password_12','2024-01-26 12:30:24','2023-09-27 03:41:32'),(13,'user13','user13@example.com','hashed_password_13','2025-05-30 10:57:54','2025-07-20 13:32:05'),(14,'user14','user14@example.com','hashed_password_14','2024-05-10 20:54:09','2024-12-25 19:30:56'),(15,'user15','user15@example.com','hashed_password_15','2023-11-21 01:14:24','2025-03-19 06:19:57');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-09-13 16:06:56
