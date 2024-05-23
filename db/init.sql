DROP DATABASE IF EXISTS ymoviehelper_db;

CREATE DATABASE ymoviehelper_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE ymoviehelper_db;


CREATE TABLE t100_software (
  id int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE t110_character (
  id int NOT NULL AUTO_INCREMENT,
  t100_id int NOT NULL,
  is_empty boolean NOT NULL DEFAULT FALSE,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (t100_id) REFERENCES t100_software (id) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (`name`, t100_id)
);

CREATE TABLE t111_emotion (
  id int NOT NULL AUTO_INCREMENT,
  t110_id int NOT NULL,
  item_path varchar(255) NOT NULL,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (t110_id) REFERENCES t110_character (id) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (`name`, t110_id)
);

CREATE TABLE t200_project (
  id int NOT NULL AUTO_INCREMENT,
  t100_id int NOT NULL,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (t100_id) REFERENCES t100_software (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE m301_item_type (
  id int NOT NULL AUTO_INCREMENT,
  ymmp_name varchar(255) NOT NULL,
  `name` varchar(40) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (ymmp_name),
  UNIQUE (`name`)
);

CREATE TABLE t310_single_item (
  id int NOT NULL AUTO_INCREMENT,
  t200_id int NOT NULL,
  m301_name varchar(10) NOT NULL,
  item_path varchar(255) NOT NULL,
  `name` varchar(10) NOT NULL,
  `length` int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (t200_id) REFERENCES t200_project (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (m301_name) REFERENCES m301_item_type (`name`) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (item_path)
);

CREATE TABLE t320_multiple_item (
  id int NOT NULL AUTO_INCREMENT,
  t200_id int NOT NULL,
  item_path varchar(255) NOT NULL,
  `name` varchar(20) NOT NULL,
  count_of_items int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (t200_id) REFERENCES t200_project (id) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (item_path)
);

CREATE TABLE t330_dynamic_item (
  id int NOT NULL AUTO_INCREMENT,
  t200_id int NOT NULL,
  item_url varchar(255) NOT NULL,
  m301_name varchar(10) NOT NULL,
  item_path_in_pc text NOT NULL,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (t200_id) REFERENCES t200_project (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (m301_name) REFERENCES m301_item_type (`name`) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (item_url),
  UNIQUE (`name`, t200_id)
);

CREATE TABLE t400_rule (
  t200_id int NOT NULL,
  voiceline_layer int NOT NULL,
  PRIMARY KEY (t200_id),
  FOREIGN KEY (t200_id) REFERENCES t200_project (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE t411_character_item_in_rule (
  id int NOT NULL AUTO_INCREMENT,
  t400_id int NOT NULL,
  t110_id int NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (t110_id, t400_id),
  FOREIGN KEY (t400_id) REFERENCES t400_rule (t200_id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t110_id) REFERENCES t110_character (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t412_empty_item_in_rule (
  id int NOT NULL AUTO_INCREMENT,
  t400_id int NOT NULL,
  t110_id int NOT NULL,
  sentence varchar(255) NOT NULL DEFAULT '空白',
  PRIMARY KEY (id),
  FOREIGN KEY (t400_id) REFERENCES t400_rule (t200_id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t110_id) REFERENCES t110_character (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t420_dynamic_item_in_rule (
  id int NOT NULL AUTO_INCREMENT,
  t400_id int NOT NULL,
  layer int NOT NULL,
  t330_id int NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (t400_id, layer),
  FOREIGN KEY (t400_id) REFERENCES t400_rule (t200_id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t330_id) REFERENCES t330_dynamic_item (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t430_static_item_in_rule (
  id int NOT NULL AUTO_INCREMENT,
  t400_id int NOT NULL,
  layer int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (t400_id) REFERENCES t400_rule (t200_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t431_single_item_in_rule (
  t430_id int NOT NULL,
  t310_id int NOT NULL,
  is_fixed_start bool NOT NULL,
  is_fixed_end bool NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t310_id) REFERENCES t310_single_item (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t432_multiple_item_in_rule (
  t430_id int NOT NULL,
  t320_id int NOT NULL,
  is_fixed_start bool NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t320_id) REFERENCES t320_multiple_item (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE m440_fixed (
  id int NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE(`name`)
);

CREATE TABLE t441_fixed_start (
  t430_id int NOT NULL,
  insert_place varchar(10) NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (insert_place) REFERENCES m440_fixed (`name`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t442_flexible_start (
  t430_id int NOT NULL,
  t110_id int NOT NULL,
  adjustment_value int NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t110_id) REFERENCES t110_character (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t451_fixed_end (
  t430_id int NOT NULL,
  `length` int NOT NULL,
  is_unique bool NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t452_flexible_end (
  t430_id int NOT NULL,
  how_many_aheads int NOT NULL,
  adjustment_value int NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO
  m301_item_type(ymmp_name, `name`)
VALUES
  (
    "YukkuriMovieMaker.Project.Items.VoiceItem, YukkuriMovieMaker",
    "ボイスアイテム"
  ),
  (
    "YukkuriMovieMaker.Project.Items.TextItem, YukkuriMovieMaker",
    "テキストアイテム"
  ),
  (
    "YukkuriMovieMaker.Project.Items.VideoItem, YukkuriMovieMaker",
    "ビデオアイテム"
  ),
  (
    "YukkuriMovieMaker.Project.Items.AudioItem, YukkuriMovieMaker",
    "オーディオアイテム"
  ),
  (
    "YukkuriMovieMaker.Project.Items.ImageItem, YukkuriMovieMaker",
    "画像アイテム"
  ),
  (
    "YukkuriMovieMaker.Project.Items.ShapeItem, YukkuriMovieMaker",
    "図形アイテム"
  ),
  (
    "YukkuriMovieMaker.Project.Items.TachieItem, YukkuriMovieMaker",
    "立ち絵アイテム"
  ),
  (
    "YukkuriMovieMaker.Project.Items.TachieFaceItem, YukkuriMovieMaker",
    "表情アイテム"
  ),
  (
    "YukkuriMovieMaker.Project.Items.EffectItem, YukkuriMovieMaker",
    "エフェクトアイテム"
  ),
  (
    "YukkuriMovieMaker.Project.Items.FrameBufferItem, YukkuriMovieMaker",
    "画面の複製"
  ),
  (
    "YukkuriMovieMaker.Project.Items.GroupItem, YukkuriMovieMaker",
    "グループ制御"
  );

INSERT INTO
  m440_fixed(`name`)
VALUES
  ("最初"),("最後");
