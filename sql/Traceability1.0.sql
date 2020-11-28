/*==============================================================*/
/* DBMS name:      MYSQL50                                      */
/* Created on:     2020/11/28 0:09:44                           */
/*==============================================================*/


drop table if exists admin;

drop table if exists batch_link;

drop table if exists goods;

drop table if exists goods_batch;

/*==============================================================*/
/* Table: admin                                                 */
/*==============================================================*/
create table admin
(
   a_id                 int,
   a_name               varchar(64),
   a_account            varchar(16),
   a_password           varchar(32),
   a_imageUrl           varchar(255)
);

alter table admin comment '管理员表';

/*==============================================================*/
/* Table: batch_link                                            */
/*==============================================================*/
create table batch_link
(
   bl_id                int not null,
   gd_id                int,
   bl_info1             varchar(100),
   bl_info2             varchar(100),
   bl_time              timestamp,
   bl_info4             varchar(100),
   bl_type              int,
   bl_state             int,
   primary key (bl_id)
);

alter table batch_link comment '环节表

type 环节类型
state 环节状态 ‘查收’,''退回'',''丢弃';

/*==============================================================*/
/* Table: goods                                                 */
/*==============================================================*/
create table goods
(
   g_id                 int not null,
   g_name               varchar(100),
   g_size               varchar(100),
   g_info               text,
   g_imgUrl             varchar(100),
   primary key (g_id)
);

alter table goods comment '商品表';

/*==============================================================*/
/* Table: goods_batch                                           */
/*==============================================================*/
create table goods_batch
(
   gd_id                int not null,
   g_id                 int,
   gd_number            varchar(100),
   gd_num               int,
   QR_codeurl           varchar(255),
   gd_creationTime      timestamp,
   gd_state             int,
   primary key (gd_id)
);

alter table goods_batch comment '商品批次

';

