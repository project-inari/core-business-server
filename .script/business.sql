-- business.tbl_businesses definition

CREATE TABLE `tbl_businesses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `industry_type` varchar(20) NOT NULL,
  `business_type` varchar(20) NOT NULL,
  `description` text,
  `phone_no` varchar(20) NOT NULL,
  `operating_hours` longtext,
  `address` varchar(255) DEFAULT NULL,
  `business_image_url` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_channels definition

CREATE TABLE `tbl_channels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_business_joinings definition

CREATE TABLE `tbl_business_joinings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_id` int NOT NULL,
  `username` varchar(20) NOT NULL,
  `status` varchar(20) NOT NULL,
  `actioned_by` varchar(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `business_id` (`business_id`),
  CONSTRAINT `tbl_business_joinings_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_business_members definition

CREATE TABLE `tbl_business_members` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_id` int NOT NULL,
  `username` varchar(20) NOT NULL,
  `role` varchar(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_business_user` (`business_id`,`username`),
  CONSTRAINT `tbl_business_members_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_categories definition

CREATE TABLE `tbl_categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_id` int NOT NULL,
  `category_name` varchar(255) NOT NULL,
  `category_picture_url` varchar(255) DEFAULT NULL,
  `description` text,
  `parent_category_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `business_id` (`business_id`),
  KEY `parent_category_id` (`parent_category_id`),
  CONSTRAINT `tbl_categories_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_categories_ibfk_2` FOREIGN KEY (`parent_category_id`) REFERENCES `tbl_categories` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_customers definition

CREATE TABLE `tbl_customers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_id` int NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `type` varchar(50) DEFAULT NULL,
  `address` text,
  `phone_no` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `business_id` (`business_id`),
  CONSTRAINT `tbl_customers_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_suppliers definition

CREATE TABLE `tbl_suppliers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `type` varchar(50) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `business_id` (`business_id`),
  CONSTRAINT `tbl_suppliers_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_tags definition

CREATE TABLE `tbl_tags` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_id` int NOT NULL,
  `tag_name` varchar(100) NOT NULL,
  `color` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_tag_name` (`business_id`,`tag_name`),
  CONSTRAINT `tbl_tags_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_warehouses definition

CREATE TABLE `tbl_warehouses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `warehouse_picture_url` text DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_warehouse_name` (`business_id`,`name`),
  CONSTRAINT `tbl_warehouses_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_category_tags definition

CREATE TABLE `tbl_category_tags` (
  `category_id` int NOT NULL,
  `tag_id` int NOT NULL,
  PRIMARY KEY (`category_id`,`tag_id`),
  KEY `tag_id` (`tag_id`),
  CONSTRAINT `tbl_category_tags_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `tbl_categories` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_category_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tbl_tags` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_customer_orders definition

CREATE TABLE `tbl_customer_orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` varchar(50) DEFAULT NULL,
  `business_id` int NOT NULL,
  `customer_id` int NOT NULL,
  `channel_id` int NOT NULL,
  `shipping_method` varchar(100) DEFAULT NULL,
  `shipping_fee` decimal(10,2) DEFAULT NULL,
  `shipping_cost` decimal(10,2) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_order_id` (`business_id`,`order_id`),
  KEY `customer_id` (`customer_id`),
  KEY `channel_id` (`channel_id`),
  CONSTRAINT `tbl_customer_orders_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_customer_orders_ibfk_2` FOREIGN KEY (`customer_id`) REFERENCES `tbl_customers` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_customer_orders_ibfk_3` FOREIGN KEY (`channel_id`) REFERENCES `tbl_channels` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_products definition

CREATE TABLE `tbl_products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_id` int NOT NULL,
  `supplier_id` int DEFAULT NULL,
  `item_name` varchar(255) DEFAULT NULL,
  `brand` varchar(100) DEFAULT NULL,
  `category_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `supplier_id` (`supplier_id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `tbl_products_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_products_ibfk_2` FOREIGN KEY (`supplier_id`) REFERENCES `tbl_suppliers` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_products_ibfk_3` FOREIGN KEY (`category_id`) REFERENCES `tbl_categories` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- business.tbl_product_variants definition

CREATE TABLE `tbl_product_variants` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `variant_name` varchar(255) DEFAULT NULL,
  `sku_no` varchar(50) DEFAULT NULL,
  `picture_url` text DEFAULT NULL,
  `base_selling_price` decimal(10,2) DEFAULT NULL,
  `base_purchase_price` decimal(10,2) DEFAULT NULL,
  `note` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `tbl_product_variants_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `tbl_products` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- business.tbl_stock_movements definition

CREATE TABLE `tbl_stock_movements` (
  `id` int NOT NULL AUTO_INCREMENT,
  `from_warehouse_id` int NOT NULL,
  `to_warehouse_id` int NOT NULL,
  `remarks` text,
  `move_date` date DEFAULT NULL,
  `move_time` time DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `from_warehouse_id` (`from_warehouse_id`),
  KEY `to_warehouse_id` (`to_warehouse_id`),
  CONSTRAINT `tbl_stock_movements_ibfk_1` FOREIGN KEY (`from_warehouse_id`) REFERENCES `tbl_warehouses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_stock_movements_ibfk_2` FOREIGN KEY (`to_warehouse_id`) REFERENCES `tbl_warehouses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- business.tbl_supplier_contacts definition

CREATE TABLE `tbl_supplier_contacts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `supplier_id` int NOT NULL,
  `full_name` varchar(100) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone_no` varchar(50) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `remarks` text DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `supplier_id` (`supplier_id`),
  CONSTRAINT `tbl_supplier_contact_ibfk_1` FOREIGN KEY (`supplier_id`) REFERENCES `tbl_suppliers` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_supplier_orders definition

CREATE TABLE `tbl_supplier_orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `receive_id` varchar(50) DEFAULT NULL,
  `business_id` int NOT NULL,
  `supplier_id` int NOT NULL,
  `status` varchar(50) DEFAULT NULL,
  `shipping_method` varchar(100) DEFAULT NULL,
  `shipping_cost` decimal(10,2) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_receive_id` (`business_id`,`receive_id`),
  KEY `supplier_id` (`supplier_id`),
  CONSTRAINT `tbl_supplier_orders_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_supplier_orders_ibfk_2` FOREIGN KEY (`supplier_id`) REFERENCES `tbl_suppliers` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_customer_order_products definition

CREATE TABLE `tbl_customer_order_products` (
  `customer_order_id` int NOT NULL,
  `variant_id` int NOT NULL,
  `selling_price_per_unit` decimal(10,2) DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `discount` decimal(10,2) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`customer_order_id`, `variant_id`),
  KEY `variant_id` (`variant_id`),
  CONSTRAINT `tbl_customer_order_products_ibfk_1` FOREIGN KEY (`customer_order_id`) REFERENCES `tbl_customer_orders` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_customer_order_products_ibfk_2` FOREIGN KEY (`variant_id`) REFERENCES `tbl_product_variants` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_inventory definition

CREATE TABLE `tbl_inventory` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_id` int NOT NULL,
  `warehouse_id` int NOT NULL,
  `variant_id` int NOT NULL,
  `quantity` int NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_inventory` (`business_id`, `warehouse_id`, `variant_id`),
  KEY `warehouse_id` (`warehouse_id`),
  KEY `variant_id` (`variant_id`),
  CONSTRAINT `tbl_inventory_ibfk_1` FOREIGN KEY (`business_id`) REFERENCES `tbl_businesses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_inventory_ibfk_2` FOREIGN KEY (`warehouse_id`) REFERENCES `tbl_warehouses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_inventory_ibfk_3` FOREIGN KEY (`variant_id`) REFERENCES `tbl_product_variants` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_product_stock_movement definition

CREATE TABLE `tbl_product_stock_movement` (
  `stock_movement_id` int NOT NULL,
  `variant_id` int NOT NULL,
  `quantity` int NOT NULL,
  PRIMARY KEY (`stock_movement_id`, `variant_id`),
  KEY `variant_id` (`variant_id`),
  CONSTRAINT `tbl_product_stock_movement_ibfk_1` FOREIGN KEY (`stock_movement_id`) REFERENCES `tbl_stock_movements` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_product_stock_movement_ibfk_2` FOREIGN KEY (`variant_id`) REFERENCES `tbl_product_variants` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_variant_tags definition

CREATE TABLE `tbl_variant_tags` (
  `variant_id` int NOT NULL,
  `tag_id` int NOT NULL,
  PRIMARY KEY (`variant_id`,`tag_id`),
  KEY `tag_id` (`tag_id`),
  CONSTRAINT `tbl_variant_tags_ibfk_1` FOREIGN KEY (`variant_id`) REFERENCES `tbl_product_variants` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_variant_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tbl_tags` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- business.tbl_supplier_order_products definition

CREATE TABLE `tbl_supplier_order_products` (
  `supplier_order_id` int NOT NULL,
  `variant_id` int NOT NULL,
  `price_per_unit` decimal(10,2) DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  PRIMARY KEY (`supplier_order_id`, `variant_id`),
  KEY `variant_id` (`variant_id`),
  CONSTRAINT `tbl_supplier_order_products_ibfk_1` FOREIGN KEY (`supplier_order_id`) REFERENCES `tbl_supplier_orders` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tbl_supplier_order_products_ibfk_2` FOREIGN KEY (`variant_id`) REFERENCES `tbl_product_variants` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;