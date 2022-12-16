table "d_platform" {
  schema  = schema.unity
  comment = "data_平台表"
  column "pk" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "name" {
    null = true
    type = varchar(20)
  }
  column "status" {
    null = true
    type = tinyint
  }
  column "created_time" {
    null = true
    type = timestamp
  }
  column "updated_time" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.pk]
  }
}
table "d_rank" {
  schema  = schema.unity
  comment = "data_头衔表"
  column "pk" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "name" {
    null    = true
    type    = varchar(20)
    default = ""
    comment = "头衔名称"
  }
  column "experience" {
    null    = true
    type    = bigint
    comment = "晋级所需经验"
  }
  column "description" {
    null    = true
    type    = varchar(100)
    default = ""
    comment = "头衔描述"
  }
  column "rights_pks" {
    null    = true
    type    = varchar(100)
    default = ""
    comment = "权益PK, 逗号(,)分隔"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建时间"
  }
  column "updated_time" {
    null      = true
    type      = timestamp
    comment   = "更新时间"
    on_update = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.pk]
  }
}
table "d_rank_rights" {
  schema  = schema.unity
  comment = "data_头衔权益关联表1N"
  column "pk" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "rights_pk" {
    null    = true
    type    = int
    comment = "权益PK"
  }
  column "rank_pk" {
    null    = true
    type    = int
    comment = "头衔PK"
  }
  primary_key {
    columns = [column.pk]
  }
  index "rank_rights_pks" {
    columns = [column.rank_pk, column.rights_pk]
  }
}
table "d_rights" {
  schema  = schema.unity
  comment = "data_权益表"
  column "pk" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "name" {
    null    = true
    type    = varchar(20)
    comment = "权益名称"
  }
  column "rights_classify_pk" {
    null    = true
    type    = int
    comment = "权益类目PK"
  }
  column "description" {
    null    = true
    type    = varchar(100)
    comment = "权益描述"
  }
  column "status" {
    null    = true
    type    = tinyint
    comment = "权益状态 1:启用 2：禁用"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建时间"
  }
  column "updated_time" {
    null    = true
    type    = timestamp
    comment = "更新时间"
  }
  primary_key {
    columns = [column.pk]
  }
}
table "d_rights_classify" {
  schema  = schema.unity
  comment = "data_权益类目表"
  column "pk" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "name" {
    null    = true
    type    = varchar(20)
    comment = "名称"
  }
  column "supperior_pk" {
    null    = true
    type    = int
    comment = "上级PK"
  }
  column "description" {
    null    = true
    type    = varchar(100)
    comment = "描述"
  }
  primary_key {
    columns = [column.pk]
  }
}
table "d_task" {
  schema  = schema.unity
  comment = "data_任务表"
  column "pk" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "id" {
    null    = true
    type    = bigint
    comment = "任务ID"
  }
  column "task_classify_pk" {
    null    = true
    type    = int
    comment = "任务类目PK"
  }
  column "owner_user_id" {
    null    = true
    type    = bigint
    comment = "任务负责人"
  }
  column "name" {
    null    = true
    type    = varchar(20)
    default = ""
    comment = "名称"
  }
  column "title" {
    null    = true
    type    = varchar(60)
    default = ""
    comment = "标题"
  }
  column "description" {
    null    = true
    type    = varchar(200)
    default = ""
    comment = "描述"
  }
  column "claim" {
    null    = true
    type    = json
    comment = "要求"
  }
  column "task_score" {
    null    = true
    type    = int
    comment = "任务积分/单"
  }
  column "task_number" {
    null    = true
    type    = int
    comment = "单数"
  }
  column "task_start_time" {
    null    = true
    type    = datetime
    comment = "任务开始时间"
  }
  column "task_end_time" {
    null    = true
    type    = datetime
    comment = "任务结束时间"
  }
  column "status" {
    null    = true
    type    = tinyint
    comment = "状态：1:草稿， 3：审核中， 5：进行中， 7：下架， 9：结束"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建时间"
  }
  column "updated_time" {
    null      = true
    type      = timestamp
    comment   = "更新时间"
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "auditor" {
    null    = true
    type    = varchar(20)
    default = ""
    comment = "审核人"
  }
  primary_key {
    columns = [column.pk]
  }
}
table "d_task_classify" {
  schema  = schema.unity
  comment = "data_任务类目表"
  column "pk" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "platform_pk" {
    null    = true
    type    = int
    comment = "平台PK"
  }
  column "name" {
    null    = true
    type    = varchar(20)
    default = ""
    comment = "任务类目名称"
  }
  column "description" {
    null    = true
    type    = varchar(200)
    default = ""
    comment = "描述"
  }
  column "status" {
    null    = true
    type    = tinyint
    comment = "状态：1：启用， 2：禁用"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建时间"
  }
  column "updated_time" {
    null      = true
    type      = timestamp
    comment   = "更新时间"
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "created_user" {
    null    = true
    type    = varchar(20)
    default = ""
    comment = "创建人"
  }
  primary_key {
    columns = [column.pk]
  }
}
table "d_task_order" {
  schema  = schema.unity
  comment = "data_任务订单表"
  column "pk" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "worker_id" {
    null    = true
    type    = int
    comment = "接单人ID"
  }
  column "task_id" {
    null    = true
    type    = bigint
    comment = "任务ID"
  }
  column "order_number" {
    null    = true
    type    = varchar(64)
    comment = "任务接单号（任务ID+自定义）"
  }
  column "score" {
    null    = true
    type    = int
    comment = "任务积分"
  }
  column "step" {
    null    = true
    type    = json
    comment = "完成步骤"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建接单时间"
  }
  column "updated_time" {
    null      = true
    type      = timestamp
    comment   = "更新接单时间"
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "evaluate" {
    null    = true
    type    = tinyint
    comment = "评价 1：未完成 2：未达标 10：已达标"
  }
  column "auditor" {
    null    = true
    type    = varchar(20)
    comment = "评审人"
  }
  primary_key {
    columns = [column.pk]
  }
}
table "d_task_order_snapshoot" {
  schema = schema.unity
  column "pk" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "order_number" {
    null    = true
    type    = varchar(64)
    comment = "任务订单号"
  }
  column "task_owner_name" {
    null    = true
    type    = varchar(20)
    comment = "任务负责人"
  }
  column "task_id" {
    null    = true
    type    = bigint
    comment = "任务ID"
  }
  column "task_name" {
    null    = true
    type    = varchar(200)
    comment = "任务名称"
  }
  column "start_time" {
    null    = true
    type    = timestamp
    comment = "开始时间"
  }
  column "end_time" {
    null    = true
    type    = timestamp
    comment = "结束时间"
  }
  column "worker_id" {
    null    = true
    type    = varchar(20)
    comment = "接单人"
  }
  column "claim" {
    null    = true
    type    = json
    comment = "要求"
  }
  column "step" {
    null    = true
    type    = json
    comment = "步骤"
  }
  column "score" {
    null    = true
    type    = int
    comment = "积分"
  }
  column "evaluate" {
    null    = true
    type    = tinyint
    comment = "评价"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建时间"
  }
  primary_key {
    columns = [column.pk]
  }
}
table "d_user_address" {
  schema  = schema.unity
  comment = "data_用户地址表"
  column "pk" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "user_pk" {
    null    = false
    type    = bigint
    comment = "用户PK"
  }
  column "state_code" {
    null    = true
    type    = int
    comment = "省"
  }
  column "city_code" {
    null    = true
    type    = int
    comment = "市"
  }
  column "street_code" {
    null    = true
    type    = int
    comment = "街道"
  }
  column "address" {
    null    = true
    type    = varchar(60)
    default = ""
    comment = "详细地址"
  }
  column "postal_code" {
    null    = true
    type    = int
    comment = "邮编"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建时间"
  }
  column "updated_time" {
    null      = true
    type      = timestamp
    comment   = "更新时间"
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "is_default" {
    null    = true
    type    = tinyint
    comment = "1: 默认"
  }
  primary_key {
    columns = [column.pk]
  }
  index "user_pk_default" {
    columns = [column.user_pk, column.is_default]
  }
}
table "d_user_copy1" {
  schema  = schema.unity
  comment = "data_用户表"
  column "pk" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "id" {
    null    = true
    type    = bigint
    comment = "对外ID"
  }
  column "supperior_pk" {
    null    = true
    type    = bigint
    default = 0
    comment = "上级PK"
  }
  column "path_pk" {
    null    = true
    type    = varchar(200)
    default = ""
    comment = "级联PK"
  }
  column "phone" {
    null    = true
    type    = bigint
    comment = "手机号"
  }
  column "account" {
    null    = true
    type    = varchar(20)
    default = ""
    comment = "账号"
  }
  column "password" {
    null    = false
    type    = varchar(20)
    default = ""
    comment = "密码"
  }
  column "name" {
    null    = true
    type    = varchar(20)
    default = ""
    comment = "名称"
  }
  column "rank" {
    null    = true
    type    = int
    comment = "权益级别"
  }
  column "score" {
    null    = true
    type    = bigint
    comment = "当前积分"
  }
  column "scope" {
    null    = true
    type    = set("x","xx","x21","x22")
    default = "x"
    comment = "x： 正常，\r\nxx： 禁用，\r\nx21： 交易冻结，\r\nx22： 任务冻结"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建时间"
  }
  column "updated_time" {
    null      = true
    type      = timestamp
    comment   = "更新时间"
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "last_login_time" {
    null    = true
    type    = timestamp
    comment = "最后登陆时间"
  }
  column "last_login_ip" {
    null    = true
    type    = varchar(20)
    comment = "最后登陆IP"
  }
  primary_key {
    columns = [column.pk]
  }
  index "id" {
    columns = [column.id]
  }
  index "mix_account_phone" {
    columns = [column.account, column.phone]
  }
  index "mix_up_path_pk" {
    columns = [column.supperior_pk, column.path_pk]
  }
}
table "d_user_oauth_account" {
  schema  = schema.unity
  comment = "data_用户三方账户表"
  column "user_pk" {
    null = false
    type = bigint
  }
  column "wechat" {
    null = true
    type = varchar(100)
  }
  column "qq" {
    null = true
    type = varchar(100)
  }
  column "dingtalk" {
    null = true
    type = varchar(100)
  }
  column "github" {
    null = true
    type = varchar(100)
  }
}
table "d_user_portrait" {
  schema  = schema.unity
  comment = "data_用户基础画像表"
  column "user_pk" {
    null = false
    type = bigint
  }
  column "unit_id" {
    null    = true
    type    = varchar(20)
    comment = "身份证"
  }
  column "borth" {
    null    = true
    type    = varchar(20)
    comment = "生日"
  }
  column "hobbies" {
    null    = true
    type    = varchar(100)
    comment = "兴趣爱好"
  }
  column "email" {
    null    = true
    type    = varchar(20)
    comment = "邮箱"
  }
  primary_key {
    columns = [column.user_pk]
  }
}
table "d_welfare" {
  schema  = schema.unity
  comment = "data_福利信息表"
  column "pk" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "id" {
    null    = true
    type    = bigint
    comment = "对外ID"
  }
  column "platform_pk" {
    null    = true
    type    = int
    comment = "平台PK"
  }
  column "welfare_classify_pk" {
    null    = true
    type    = int
    comment = "福利类型PK"
  }
  column "name" {
    null    = true
    type    = varchar(20)
    default = ""
    comment = "福利名称"
  }
  column "title" {
    null    = true
    type    = varchar(100)
    default = ""
    comment = "福利标题"
  }
  column "description" {
    null    = true
    type    = varchar(200)
    default = ""
    comment = "福利描述"
  }
  column "link" {
    null    = true
    type    = varchar(100)
    default = ""
    comment = "福利链接"
  }
  column "status" {
    null    = true
    type    = tinyint
    comment = "状态; 1:启用； 2：禁用"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建时间"
  }
  column "updated_time" {
    null      = true
    type      = timestamp
    comment   = "更新时间"
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "created_user" {
    null    = true
    type    = varchar(20)
    default = ""
    comment = "创建人"
  }
  primary_key {
    columns = [column.pk]
  }
  index "classify_pk" {
    columns = [column.welfare_classify_pk]
  }
  index "mix_name_title_status" {
    columns = [column.name, column.title, column.status]
  }
}
table "d_welfare_classify" {
  schema  = schema.unity
  comment = "data_福利类目表"
  column "pk" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "platform_pk" {
    null    = true
    type    = int
    comment = "平台PK"
  }
  column "name" {
    null    = true
    type    = varchar(20)
    comment = "福利类目名称"
  }
  column "description" {
    null    = true
    type    = varchar(200)
    default = ""
    comment = "描述"
  }
  column "status" {
    null    = true
    type    = tinyint
    comment = "状态; 1:启用 2：禁用"
  }
  column "created_time" {
    null    = true
    type    = timestamp
    comment = "创建时间"
  }
  column "updated_time" {
    null      = true
    type      = timestamp
    comment   = "删除时间"
    on_update = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.pk]
  }
}
table "meta_asset_suffix" {
  schema = schema.unity
  column "pk" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "name" {
    null    = true
    type    = varchar(60)
    comment = "名称"
  }
  column "description" {
    null    = true
    type    = varchar(100)
    comment = "描述"
  }
  column "scene" {
    null    = true
    type    = varchar(100)
    comment = "场景"
  }
  primary_key {
    columns = [column.pk]
  }
}
schema "unity" {
  charset = "utf8mb4"
  collate = "utf8mb4_general_ci"
}
