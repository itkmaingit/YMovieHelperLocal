PRAGMA foreign_keys=ON;

CREATE TABLE t100_software (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(20) NOT NULL
);

CREATE TABLE t110_character (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t100_id INTEGER NOT NULL,
  is_empty INTEGER NOT NULL DEFAULT 0,
  name VARCHAR(20) NOT NULL,
  FOREIGN KEY (t100_id) REFERENCES t100_software (id) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (name, t100_id)
);

CREATE TABLE t111_emotion (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t110_id INTEGER NOT NULL,
  item_path VARCHAR(255) NOT NULL,
  name VARCHAR(20) NOT NULL,
  FOREIGN KEY (t110_id) REFERENCES t110_character (id) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (name, t110_id)
);

CREATE TABLE t200_project (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t100_id INTEGER NOT NULL,
  name VARCHAR(20) NOT NULL,
  FOREIGN KEY (t100_id) REFERENCES t100_software (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE m301_item_type (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  ymmp_name VARCHAR(255) NOT NULL,
  name VARCHAR(10) NOT NULL,
  UNIQUE (ymmp_name),
  UNIQUE (name)
);

CREATE TABLE t310_single_item (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t200_id INTEGER NOT NULL,
  m301_name VARCHAR(10) NOT NULL,
  item_path VARCHAR(255) NOT NULL,
  name VARCHAR(10) NOT NULL,
  length INTEGER NOT NULL,
  FOREIGN KEY (t200_id) REFERENCES t200_project (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (m301_name) REFERENCES m301_item_type (name) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (item_path)
);

CREATE TABLE t320_multiple_item (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t200_id INTEGER NOT NULL,
  item_path VARCHAR(255) NOT NULL,
  name VARCHAR(20) NOT NULL,
  count_of_items INTEGER NOT NULL,
  FOREIGN KEY (t200_id) REFERENCES t200_project (id) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (item_path)
);

CREATE TABLE t330_dynamic_item (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t200_id INTEGER NOT NULL,
  item_url VARCHAR(255) NOT NULL,
  m301_name VARCHAR(10) NOT NULL,
  item_path_in_pc TEXT NOT NULL,  -- VARCHARのサイズ制限が大きすぎるためTEXTに変更
  name VARCHAR(20) NOT NULL,
  FOREIGN KEY (t200_id) REFERENCES t200_project (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (m301_name) REFERENCES m301_item_type (name) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE (item_url),
  UNIQUE (name, t200_id)
);

CREATE TABLE t400_rule (
  t200_id INTEGER NOT NULL,
  voiceline_layer INTEGER NOT NULL,
  PRIMARY KEY (t200_id),
  FOREIGN KEY (t200_id) REFERENCES t200_project (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t411_character_item_in_rule (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t400_id INTEGER NOT NULL,
  t110_id INTEGER NOT NULL,
  UNIQUE (t110_id, t400_id),
  FOREIGN KEY (t400_id) REFERENCES t400_rule (t200_id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t110_id) REFERENCES t110_character (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t412_empty_item_in_rule (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t400_id INTEGER NOT NULL,
  t110_id INTEGER NOT NULL,
  sentence VARCHAR(255) NOT NULL DEFAULT '空白',
  FOREIGN KEY (t400_id) REFERENCES t400_rule (t200_id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t110_id) REFERENCES t110_character (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t420_dynamic_item_in_rule (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t400_id INTEGER NOT NULL,
  layer INTEGER NOT NULL,
  t330_id INTEGER NOT NULL,
  UNIQUE (t400_id, layer),
  FOREIGN KEY (t400_id) REFERENCES t400_rule (t200_id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t330_id) REFERENCES t330_dynamic_item (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t430_static_item_in_rule (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  t400_id INTEGER NOT NULL,
  layer INTEGER NOT NULL,
  FOREIGN KEY (t400_id) REFERENCES t400_rule (t200_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t431_single_item_in_rule (
  t430_id INTEGER NOT NULL,
  t310_id INTEGER NOT NULL,
  is_fixed_start INTEGER NOT NULL,
  is_fixed_end INTEGER NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t310_id) REFERENCES t310_single_item (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t432_multiple_item_in_rule (
  t430_id INTEGER NOT NULL,
  t320_id INTEGER NOT NULL,
  is_fixed_start INTEGER NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t320_id) REFERENCES t320_multiple_item (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE m440_fixed (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(10) NOT NULL,
  UNIQUE (name)
);

CREATE TABLE t441_fixed_start (
  t430_id INTEGER NOT NULL,
  insert_place VARCHAR(10) NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (insert_place) REFERENCES m440_fixed (name) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t442_flexible_start (
  t430_id INTEGER NOT NULL,
  t110_id INTEGER NOT NULL,
  adjustment_value INTEGER NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (t110_id) REFERENCES t110_character (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t451_fixed_end (
  t430_id INTEGER NOT NULL,
  length INTEGER NOT NULL,
  is_unique INTEGER NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE t452_flexible_end (
  t430_id INTEGER NOT NULL,
  how_many_aheads INTEGER NOT NULL,
  adjustment_value INTEGER NOT NULL,
  PRIMARY KEY (t430_id),
  FOREIGN KEY (t430_id) REFERENCES t430_static_item_in_rule (id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO
  m301_item_type (ymmp_name, name)
VALUES
  ("YukkuriMovieMaker.Project.Items.VoiceItem, YukkuriMovieMaker", "ボイスアイテム"),
  ("YukkuriMovieMaker.Project.Items.TextItem, YukkuriMovieMaker", "テキストアイテム"),
  ("YukkuriMovieMaker.Project.Items.VideoItem, YukkuriMovieMaker", "ビデオアイテム"),
  ("YukkuriMovieMaker.Project.Items.AudioItem, YukkuriMovieMaker", "オーディオアイテム"),
  ("YukkuriMovieMaker.Project.Items.ImageItem, YukkuriMovieMaker", "画像アイテム"),
  ("YukkuriMovieMaker.Project.Items.ShapeItem, YukkuriMovieMaker", "図形アイテム"),
  ("YukkuriMovieMaker.Project.Items.TachieItem, YukkuriMovieMaker", "立ち絵アイテム"),
  ("YukkuriMovieMaker.Project.Items.TachieFaceItem, YukkuriMovieMaker", "表情アイテム"),
  ("YukkuriMovieMaker.Project.Items.EffectItem, YukkuriMovieMaker", "エフェクトアイテム"),
  ("YukkuriMovieMaker.Project.Items.FrameBufferItem, YukkuriMovieMaker", "画面の複製"),
  ("YukkuriMovieMaker.Project.Items.GroupItem, YukkuriMovieMaker", "グループ制御");

INSERT INTO
  m440_fixed (name)
VALUES
  ("最初"),("最後");
