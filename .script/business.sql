CREATE TABLE `tbl_businesses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `industry_type` varchar(20) NOT NULL,
  `business_type` varchar(20) NOT NULL,
  `description` text,
  `phone_no` varchar(20) NOT NULL,
  `operating_hours` longtext DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `business_image_url` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`name`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `tbl_business_members` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_name` varchar(10) NOT NULL,
  `username` varchar(20) NOT NULL,
  `role` varchar(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`business_name`,`username`),
  UNIQUE KEY `id` (`id`),
  KEY `tbl_business_members_business_name_IDX` (`business_name`) USING BTREE,
  KEY `tbl_business_members_username_IDX` (`username`) USING BTREE,
  CONSTRAINT `tbl_business_members_ibfk_1` FOREIGN KEY (`business_name`) REFERENCES `tbl_businesses` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `tbl_business_joinings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_name` varchar(10) NOT NULL,
  `username` varchar(20) NOT NULL,
  `status` varchar(20) NOT NULL,
  `actioned_by` varchar(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`business_name`,`username`),
  UNIQUE KEY `id` (`id`),
  KEY `tbl_business_joinings_business_name_IDX` (`business_name`) USING BTREE,
  KEY `tbl_business_joinings_username_IDX` (`username`) USING BTREE,
  CONSTRAINT `tbl_business_joinings_ibfk_1` FOREIGN KEY (`business_name`) REFERENCES `tbl_businesses` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
