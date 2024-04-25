package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type BuildingInfo struct {
	id              string      // 唯一id;
	specId          int32       // 策划配表id;
	groupId         int32       // 建筑组;
	level           int32       // 建筑等级;
	x               int32       // 坐标x;
	y               int32       // 坐标y;
	status          int32       // 建筑状态 目前 1、建筑中  2、空闲中  3、升级中 4、修复中 5、修建地基;
	updateEndTime   int64       // 当状态为升级中时  这个代表升级的结束时间;
	beginTime       int64       //建筑产生订单或者资源产出总开始时间;
	productNum      int32       //生产产品的数量;
	currentDuration int64       //当前订单消耗的时间(订单持续时间);
	custom          string      //自定义结构(用于存储特殊信息)
	broken          int32       // 破损状态
	productTarId    int32       //订单配置ID(用于升级后的);
	productCfgId    int32       //当前建筑订单配置ID(所属类型);
	resourceNum     float32     //资源产生的数量
	handleType      int32       // 建筑当前正在干什么工作
	allianceHelp    int32       // 1:已发送过联盟帮助，0：未发送过
	workerNum       int32       // 当前幸存者（工人）数量
	mainId          int32       // 主将ID(城墙建筑使用)
	supId           int32       // 副将ID(城墙建筑使用)
	quest           []QuestInfo // 建筑修复任务
	techIds         []int32     // 建筑科技树
	techPoint       int32       // 建筑科技点数

	//ID              string  `json:"id"`
	//SpecId          int64   `json:"specId"`
	//GroupId         int64   `json:"groupId"`
	//Level           int32   `json:"level"`
	//X               int32   `json:"x"`
	//Y               int32   `json:"y"`
	//Status          int32   `json:"status"`
	//UpdateEndTime   string  `json:"updateEndTime"` // Assuming it's a string because of the provided example
	//BeginTime       string  `json:"beginTime"`     // Assuming it's a string because of the provided example
	//ProductNum      int32   `json:"productNum"`
	//CurrentDuration string  `json:"currentDuration"` // Assuming it's a string because of the provided example
	//Broken          int32   `json:"broken"`
	//ProductTarId    int64   `json:"productTarId"`
	//ProductCfgId    int64   `json:"productCfgId"`
	//ResourceNum     float64 `json:"resourceNum"`
	//HandleType      int32   `json:"handleType"`
	//AllianceHelp    int32   `json:"allianceHelp"`
	//WorkerNum       int32   `json:"workerNum"`
	//TechPoint       int32   `json:"techPoint"`
	//TechIds         []int64 `json:"techIds,omitempty"` // This field is optional and may not be present in all buildings
}

type PlayerCity struct {
	//Id string
	//X  int32
	//Y  int32
	////UserBuildings []*BuildingInfo //玩家所有建筑
	//AllianceId string // 联盟id
	//Hp         int32  // 主城城防血量
	//FireTime   int64  // 着火时间戳
	//LastTime   int64  // 上次维修刷新时间
	//Status     int32  // 主城状态
	ID string `json:"id"`
	X  int32  `json:"x"`
	Y  int32  `json:"y"`
	//UserBuildings []BuildingInfo `json:"userBuildings"`
	AllianceId string `json:"allianceId"`
	Hp         int32  `json:"hp"`
	FireTime   string `json:"fireTime"` // Assuming it's a string because of the provided example
	LastTime   string `json:"lastTime"` // Assuming it's a string because of the provided example
	Status     int32  `json:"status"`
}

type CustomVar struct {
	key   string //客户端自定义参数key
	value string //客户端自定义参数value
}

type ItemInfo struct {
	id         string // 唯一id
	specId     int32  // 策划配表id
	num        int32  // 数量
	expireTime string // 过期时间 毫秒
}

type ResourceInfo struct {
	Type int32   //类型
	Num  float32 //数量
}

type ArmyInfo struct {
	id         string // 士兵id;
	metaId     int32  // 策划配表id;
	hurtNum    int32  // 待治疗
	cureingNum int32  // 治理中的数量
	normalNum  int32  // 正常数量
}

type AreaUnlockInfo struct {
	areaId   int32 // 地块id
	unlocked int32 // 已解锁的关卡
}

type ScienceInfo struct {
	groupId int32 // 科技类型
	level   int32 // 等级
	//  required int32 specId		= 1;	// 配置id
	endTime int64 // 升级结束时间
}

type QuestProgress struct {
	id  int32   // id,序号
	num float32 // 进度 double
}

type QuestInfo struct {
	questId  int32           // id
	progress []QuestProgress // 进度 double
	status   int32           // 状态 进行中/已完成/已领奖
}
type PBECounterType int

const (
	LAST_DAILY_UPDATE_TIME        PBECounterType = iota //每日更新 最后一次更新的时间 这个很重要，不要修改
	ACTIVITY_VALUE                                      //每日活跃度
	ACTIVITY_REWARD                                     //每日活跃度奖励领取记录 按位处理
	TROOP_RECON_NUM                                     //玩家侦查斥候队伍计数器
	GOLD_COLLECTION_NUM                                 //每日金币收割机免费次数
	GOLD_OUTPUT_NUM                                     //收割机6小时收益
	VIDEO_NUM                                           //观看视频的次数
	NO_HAVE_VIDEO_NUM                                   //没有视频可拉去，每日白嫖次数
	PVE_CRASH_MAX_LEVEL                                 //对冲战斗通关的最大关卡数
	FIRST_CREATE_ARMY                                   //首次造兵(使用固定的时间)
	VIDEO_REWARD_NUM                                    //观看视频次数的奖励
	DRAW_CARDS_NEW_NUM                                  //白银宝箱抽卡次数
	DRAW_CARDS_NORMAL_FREE_NUM                          //普通抽卡免费抽卡次数
	DRAW_CARDS_NORMAL_FREE_TIME                         //普通每日免费抽卡间隔时间
	DRAW_CARDS_NORMAL_DAILY_NUM                         //黄金宝箱抽卡次数
	DRAW_CARDS_NORMAL_NUM                               //普通抽奖次数
	DRAW_CARDS_NORMAL_ORANGE_NUM                        //普通抽奖固定次数
	DRAW_CARDS_HIGH_DAILY_NUM                           //高级每日抽奖次数
	DRAW_CARDS_HIGH_NUM                                 //高级抽卡次数
	DRAW_CARDS_HIGH_ORANGE_NUM                          //高级抽卡固定次数
	MYSTERY_SHOP_NUM                                    //每日神秘商店出现次数
	MYSTERY_SHOP_REFRESH_NUM                            //神秘商店刷新次数
	PLAYER_AP                                           //玩家体力计数器 actionPoints
	PLAYER_AP_DAILY_FREE                                //玩家每天免费领取体力次数 从0开始加到上限停止 隔天清零
	PLAYER_AP_DAILY_FREE_CD                             //玩家免费领取体力CD时间戳
	BUILDING_POWER                                      //建筑战力
	SCIENCE_POWER                                       //科技战力
	TROOP_POWER                                         //部队战力
	HERO_POWER                                          //统帅战力
	HISTORY_HIGH_POWER                                  //历史最高战力
	FIGHT_WIN_NUM                                       //战斗胜利
	FIGHT_FAIL_NUM                                      //战斗失败
	DIE_NUM                                             //阵亡
	SCOUT_NUM                                           //侦查次数
	RESOURCE_GATHER                                     //资源采集
	RESOURCE_HELP                                       //资源援助
	ALLIANCE_HELP_NUM                                   //联盟帮助次数
	ALLIANCE_DIAMOND_DONATE_NUM                         //当日联盟钻石捐献次数
	HELP_SCORE                                          //当日联盟帮助获得的积分
	HELP_CURE_COUNT                                     //当日治疗被帮助次数
	HELP_CURE_CD_END                                    //当日发送治疗的CD结束时间
	PAYMENT_DIAMOND                                     //付费钻石
	FREE_DIAMOND                                        //免费钻石
	PAYMENT_MONEY                                       //购买礼包消费金额
	ALLIANCE_SCORE_BUILDING_DAILY                       //当日玩家联盟建筑建造获取联盟积分累积
	ALLIANCE_GET_RES_BEGIN_TIME                         //联盟领地个人领取资源时间戳(开始时间)
	ALLIANCE_PER_RES_FOOD                               //记录个人联盟资源粮食(记录下一生产速度之前的总值)
	ALLIANCE_PER_RES_WOOD                               //记录个人联盟资源木材(记录下一生产速度之前的总值)
	ALLIANCE_PER_RES_STONE                              //记录个人联盟资源石头(记录下一生产速度之前的总值)
	ALLIANCE_PER_RES_COIN                               //记录个人联盟资源金币(记录下一生产速度之前的总值)
	ALLIANCE_PER_RES_CHANGE_TIME                        //记录下一次速度变化的时间戳
	WORKER_TOTAL_COUNT                                  //幸存者（工人）数量获得累计
	MONTHLY_CARD_END_TIME                               //记录月卡结束时间戳
	WEEKLY_CARD_END_TIME                                //记录周卡结束时间戳
	DRAW_CARD_SILVER_NUM                                //白银宝箱次数
	DRAW_CARD_SILVER_FREE_NUM                           //白银宝箱每日免费次数
	DRAW_CARD_SILVER_FREE_TIME                          //白银宝箱每日免费间隔时间
	DRAW_CARD_SILVER_DAILY_NUM                          //白银宝箱每日抽奖次数
	DRAW_CARD_GOLD_NUM                                  //黄金宝箱次数
	DRAW_CARD_GOLD_DAILY_NUM                            //黄金宝箱每日抽奖次数
	DRAW_CARD_GOLD_FREE_TIME                            //黄金宝箱免费间隔时间
	DRAW_CARD_EQUIP_NUM                                 //装备宝箱次数
	DRAW_CARD_EQUIP_DAILY_NUM                           //装备宝箱每日抽奖次数
	DRAW_CARD_EQUIP_FREE_TIME                           //装备宝箱免费间隔时间
	PLAYER_KEY                                          //玩家钥匙计数器(关卡消耗资源)
	NPC_CAN_ATTACK_LEVEL                                //可攻击npc最高等级
	ARMY_T1_INFANTRY_NUM                                //玩家T1步兵士兵总数
	ARMY_T1_CAVALRY_NUM                                 //玩家T1骑兵士兵总数
	ARMY_T1_RANGED_NUM                                  //玩家T1弓兵士兵总数
	ARMY_T2_INFANTRY_NUM                                //玩家T2步兵士兵总数
	ARMY_T2_CAVALRY_NUM                                 //玩家T2骑兵士兵总数
	ARMY_T2_RANGED_NUM                                  //玩家T2弓兵士兵总数
	ARMY_T3_INFANTRY_NUM                                //玩家T3步兵士兵总数
	ARMY_T3_CAVALRY_NUM                                 //玩家T3骑兵士兵总数
	ARMY_T3_RANGED_NUM                                  //玩家T3弓兵士兵总数
	ARMY_T4_INFANTRY_NUM                                //玩家T4步兵士兵总数
	ARMY_T4_CAVALRY_NUM                                 //玩家T4骑兵士兵总数
	ARMY_T4_RANGED_NUM                                  //玩家T4弓兵士兵总数
	ARMY_T5_INFANTRY_NUM                                //玩家T5步兵士兵总数
	ARMY_T5_CAVALRY_NUM                                 //玩家T5骑兵士兵总数
	ARMY_T5_RANGED_NUM                                  //玩家T5弓兵士兵总数
	PLAYER_LOGIN_DAY_UPDATE                             // 玩家累计登录的更新时间
	PLAYER_DURATION_LOGIN_DAY                           // 连续登录天数
	ONLINE_REWARD_TIME                                  // 下次领取持续在线奖励的时间
	ONLINE_REWARD_NUMBER                                // 持续在线奖励上次是第几次领取
	LAST_WEEK_UPDATE_TIME                               //每日更新 最后一次更新的时间 这个很重要，不要修改
)

type CounterInfo struct {
	Type PBECounterType //类型
	num  string         //数量 long
}

// PBETroopType 是一个基于 int 的自定义类型，用于表示部队类型
type PBETroopType int

// 定义 PBETroopType 的可能值
const (
	ATTACK_NPC           PBETroopType = iota // 打野怪
	ATTACK_PLAYER                            // 打玩家世界地图上走动的部队
	GATHER                                   // 占领资源
	SCOUT                                    // 侦查
	BACK_TO_MY_CITY                          // 返回主城
	MOVE_TO_PLACE                            // 移动到某点
	START_RALLY                              // 发起集结
	JOIN_RALLY                               // 加入集结
	NPC_TROOP                                // 野怪行军
	ATTACK_CITY                              // 打玩家主城
	RALLY_NULL                               // 非集结类型
	ENTER_ABUILDING                          // 进入联盟建筑
	ATTACK_ABUILDING                         // 攻击联盟建筑
	ATTACK_NEUTRAL_BUILD                     // 攻击中立建筑
	ENTER_NEUTRAL_BUILD                      // 进入中立建筑
	BACK_MY_UNREAL_CITY                      // 返回假主城
)

// PBETroopStatus 是一个基于 int 的自定义类型，用于表示部队状态
type PBETroopStatus int

// 定义 PBETroopStatus 的可能值
const (
	PBETS_DEFAULT     PBETroopStatus = iota // 默认值，无意义
	IDLE                                    // 待机
	MOVE                                    // 移动
	WORK                                    // 采集
	FOLLOW                                  // 集结队列中的跟随
	WAIT                                    // 集结等待状态
	BATTLE                                  // 战斗中状态
	BATTLE_FAIL                             // 战斗失败 溃败状态
	BACK_CITY                               // 回城
	BATTLE_MOVE                             // 战斗中主动移动，不会出招那种
	BACK_NPC_BORN                           // 野怪返回出生点位置
	BATTLE_FOLLOW                           // 战斗中追击状态，可以出招攻击
	JOIN_RALLY_FOLLOW                       // 参与集结部队的跟随状态
	STATIONED                               // 入驻状态 联盟建筑、盟友主城、联盟关卡、神庙等
	ALLIANCE_WORK                           // 联盟矿采集
	CITY_DEFENSE                            // 在主城内
	UP_TROOP                                // 部队信息变更
	STATIONED_BATTLE                        // 驻防战斗状态
)

type PBVectorI struct {
	x   int32
	y   int32
	sId int32 //服ID
}

type ArmyKVInfo struct {
	id         int32 //兵种ID
	aliveCount int32 //该兵种数量(存活)
	minorCount int32 // 轻伤数量
	majorCount int32 // 重伤数量
	deadCount  int32 // 死亡数量
	healCount  int32 // 治疗量
}

type ComingTroop struct {
	troopId string //部队ID
	x       int32  //初始坐标
	y       int32  //初始坐标
}

type AoiTroopBuff struct {
	uuid     string  // buff的uuid
	specId   int32   // buff的配置id
	value    float32 // buff效果数值
	duration int32   // buff持续回合数
	endTime  int64   // buff结束时间戳
	stack    int32   // buff层数
}

type AoiTroop struct {
	totalArmyNum int64          //部队总兵数
	armyNum      int64          //部队当前兵数
	mainHeroId   int32          //英雄主将
	supHeroId    int32          //英雄副将
	rage         int32          //怒气
	rageMax      int32          //最大怒气
	buffs        []AoiTroopBuff //部队buff信息
	buffsEmpty   bool           //部队buff信息清空标记（标记清空true 一般情况为false）
	mainTarget   string         //战斗主目标
	heroId       []int32        //上阵英雄ID列表
}

type UserNameInfo struct {
	id           string // 玩家id
	name         string // 玩家名称
	headUrl      string // 头像
	frame        int32  // 头像框
	allianceTag  string // 联盟tag
	allianceName string // 联盟名称
	camp         string // 阵营
	power        string // 总战力
	serverid     int32  //服务器id
	castleLv     int32  //主堡等级
}

type TroopTargetInfo struct {
	metaId       int32
	userNameInfo UserNameInfo //玩家信息
}

type WorldTroopInfo struct {
	id                string          //部队id
	updateTime        int64           //上次更新时间
	troopType         PBETroopType    //部队类型
	status            PBETroopStatus  //队伍状态
	troopPos          PBVectorI       //当前队伍坐标信息
	soldiers          []ArmyKVInfo    //队列中的士兵信息
	currExploitAmount int64           //当前采集量
	maxExploitAmount  int64           //最大采集量
	speed             float64         //行军速度
	power             int64           //部队战力
	keyPoints         []PBVectorI     //行军路线关键点坐标 起点 拐点 终点
	keyPointTimes     int64           //每个关键点时间戳 起点 拐点 终点
	rallyType         PBETroopType    //部队集结类型 发起集结START_RALLY 加入集结JOIN_RALLY 其他值为非集结
	rallyId           string          //部队如果参加集结的话，这里是参与集结的ID 也是车头troop的id
	allianceId        string          //部队玩家所属联盟ID
	targetId          string          //部队行军目标ID
	comingTroops      []ComingTroop   //即将到来部队信息
	comingEmpty       bool            //是否comingTroop有改变（标记删除true 一般情况为false）
	aoiTroops         AoiTroop        //战斗相关客户端表现数据
	userNameInfo      UserNameInfo    //玩家信息
	targetInfo        TroopTargetInfo // 行军目标信息
}

type AreaGridInfo struct {
	id   string // 地格ID
	x    int32  // 坐标x
	y    int32  // 坐标y
	info string // 地格信息
}

type EffectInfo struct {
	effectId  int32 //
	value     int32
	startTime int64
	endTime   int64 // 时间格式的可以使用 int64
	itemId    int32 //使用道具的策划配置ID
	source    int32 //来源类型 buffAttribute.proto BuffAttributeSource
}

type KillPoint struct {
	level int32 // 等级
	num   int32 // 数量
	point int32 // 积分
}

type UserData struct {
	registerTime       int64  // 注册时间 毫秒
	playerLevel        int32  // 玩家等级
	todayEndTime       int64  // 当天结束的时间(跨天时间)
	allianceId         string // 联盟id
	allianceRank       int32  // 联盟级别，联盟id > 0 时有效
	appliedAllianceIds string // 已申请过的联盟，联盟id = 0 时有效
	//  optional  int64    dismissTime      // 解散时间(只对盟主有意义,>0表示正在解散倒计时)
	power        int64
	killPoints   []KillPoint  //击杀积分
	userNameInfo UserNameInfo //玩家的名字数据
	anonymous    bool         // 联盟礼包是否匿名发送
}

// SearchType 是一个基于 int 的自定义类型，用于表示查询类型
type SearchType int

// 定义 SearchType 的可能值
const (
	GOLD    SearchType = iota // 金
	FOOD                      // 粮
	WOOD                      // 木
	STONE                     // 石
	MONSTER                   // NPC
)

type SearchRank struct {
	minRank    int32      // 最低等级
	maxRank    int32      // 最高等级
	searchType SearchType // 类型
}

type GiftInfo struct {
	id       string  //唯一ID
	metaId   int32   //礼包ID
	disTime  int64   //消失时间
	multiple float32 //倍数
}

type VideoReward struct {
	giftInfo []GiftInfo //礼包信息
	free     bool       //是否免观看
}

type Skills struct {
	id    int32 //技能ID
	level int32 //技能等级
}

type Talents struct {
	id    int32 //天赋ID
	level int32 //天赋等级
}

type HeroPower struct {
	totalPower  int32 // 总战力
	basePower   int32 // 基础战力
	levelPower  int32 // 等级战力
	skillPower  int32 // 技能战力
	talentPower int32 // 天赋战力
	equipPower  int32 // 装备战力
	starPower   int32 // 装备战力
}

// HeroStatus 是一个基于 int 的自定义类型，用于表示英雄状态
type HeroStatus int

// 定义 HeroStatus 的可能值
const (
	Default HeroStatus = 1 // 默认
	Battle  HeroStatus = 2 // 出征
	Defense HeroStatus = 3 // 驻防
)

type SkillDepletes struct {
	metaId int32 // 道具配置ID
	num    int32 // 数量
}
type StarDepletes struct {
	metaId int32 // 道具配置ID
	num    int32 // 数量
}

type HeroAttribute struct {
	attack           int32 // 攻击力
	attackAttribute  int32 // 攻击力属性
	defense          int32 // 防御
	defenseAttribute int32 // 防御属性
	hp               int32 // 血量
	hpAttribute      int32 // 血量属性
}

type EquipPosition struct {
	position int32  // 装备位置
	id       string // 装备唯一ID
}

type HeroAttr struct {
	attr []Attr
}

type Attr struct {
	id    int32
	value float32
}

type HeroInfo struct {
	metaId        int32           // 英雄metaId
	level         int32           // 英雄等级
	exp           string          // 英雄经验
	star          int32           // 英雄星级
	starExp       string          // 英雄星级经验
	skills        []Skills        // 技能
	talents       []Talents       // 天赋
	getTime       int64           // 获取时间
	kills         string          // 击杀
	heroPower     HeroPower       // 战力
	status        HeroStatus      // 英雄状态
	skillDepletes []SkillDepletes // 技能消耗
	starDepletes  []StarDepletes  // 星级消耗
	heroAttribute HeroAttribute   // 英雄属性
	equipPosition []EquipPosition // 装备位置信息(记录装备ID，位置)
	heroAttr      HeroAttr        // 英雄属性
	traingSeat    int32           // 特训位置
	lastResetTime int64           // 上次重置时间
	slgAttr       HeroAttr        // slg属性
}

type BuildingQueueInfo struct {
	queueId    string // 队列id
	createTime int64
	expireTime int64  //队列过期时间，0表示永久队列，时间类型可以使用int64
	buildingId string // 工作的建筑唯一id,0表示空闲
	isTry      bool   //  是否试用过
	buyTimes   int32  // 临时队列购买次数
}

type BuffAttributeInfo struct {
	id     int32
	source int32  // 来源
	value  string // 值 long
}

type LotteryInfo struct {
	metaId   int32 // 抽卡类型
	freeTime int64 // 上次免费时间
	freeNum  int32 // 免费次数
	dayNum   int32 // 当天剩余的次数
	num      int32 // 抽卡次数
}

type ShopSimpleInfo struct {
	shopId          int32  //
	beginTime       int64  // 商店生存期开始时间
	endTime         int64  // 商店生存期结束时间
	nextOpenTime    int64  // 商店最近一次的开启时间
	nextCloseTime   int64  // 商店最近一次的关闭时间
	opening         bool   // 是否开启中
	ver             string // 商店当前版本标识()
	nextRefreshTime int64  // 定时刷新商品的，下次刷新时间
}

type PlayerActivitySimpleInfo struct {
	activityId   int32
	beginTime    int64
	endTime      int64
	isEnd        int32
	extra        string // 活动的额外数据 json格式的字符串
	serial       int32  // 活动开启的次数
	retainTimeMs int64  //剩余时间戳（活动展现的结束时间）
}

type ActivityTargetInfo struct {
	targetId int32
	progress QuestProgress // 完成进度
	status   int32         // 完成状态
	status2  int32         // 备用的第二状态
}

type PlayerActivityInfo struct {
	simpleInfo PlayerActivitySimpleInfo // 基本信息
	targetInfo ActivityTargetInfo       // 任务进度
}

type MysteryShopInfo struct {
	shopItemId int32   // 商品配置id
	discount   float32 // 折扣
	buyCount   int32   // 已购数量
}

type MysteryShop struct {
	endTime   int64           // 结束时间
	shopInfos MysteryShopInfo // 所有神秘商品
	ver       string          // 商店当前版本标识()
}

type BuyStatus int

const (
	BUY_STATUS_0 BuyStatus = iota // 可购买
	BUY_STATUS_1                  // 不可购买
)

type RechargeInfo struct {
	id          string    // 礼包唯一ID
	metaId      int32     // 礼包ID(策划配表recharge)
	beginTime   int64     // 开始时间
	endTime     int64     // 结束时间
	realEndTime int64     // 有效结束时间
	resetTime   int64     // 重置周期时间
	limitCount  int32     // 限购数量
	count       int32     // 已购数量
	isEnd       int32     // 是否结束
	buyStatus   BuyStatus // 礼包购买状态
	appKey      string    // 平台配置的礼包key
	subscribe   int32     // 是否订阅的
}

type VipInfo struct {
	vipLevel    int32 //vip等级
	vipPoint    int32 //vip点数
	loginDayNum int32 //连续登陆时间
}

type QuestChapter struct {
	chapter    int32
	isFinished int32 // 是否完成当前章节，1：完成，可领取奖励
	maxChapter int32 // 最大章节，客户端判断是否完成所有章节
}

// 探索小队状态
type ExplorerTeamStatus int

const (
	ETS_IDLE      = 0 // 休闲
	ETS_EXPLORING = 1 // 工作中
	ETS_IN_CD     = 2 // CD中
	ETS_LOCK      = 3 // 锁定中 客户端使用，服务器不使用
)

// ExplorerTeamInfo 小队数据
type ExplorerTeamInfo struct {
	teamId    int32
	workerNum int32
	endTime   int64
	status    ExplorerTeamStatus
}

type ExploreQuestInfo struct {
	questId    string // 唯一id
	specId     int32  // 配置表id
	expireTime int64  // 消失时间
	Type       int32  // 任务类型
	param      string // 任务参数， EQT_MONSTER：怪物唯一id
	xPos       int32  // 坐标，世界坐标
	yPos       int32  //
	status     int32  // 任务完成状态
	costAP     int32  // 体力扣除
	trigerType int32  // 触发类型-对应 探索任务触发类型
}

type EquipInfo struct {
	id         string // 装备唯一ID
	metaId     int32  // 装备配置ID
	heroId     int32  // 所属英雄ID
	level      int32  // 等级
	forgeLevel int32  // 锻造等级
}

type PvePassInfo struct {
	passId      int32 // 关卡ID
	star        int32 // 关卡星级
	isFirstPass bool  // 是否首次通关
	pass        bool  // 是否已通关
	reward      bool  // 是否已领奖
}

type ChapterPassInfo struct {
	pvePassInfo    PvePassInfo // 关卡信息
	unlockChapters []int32     // 已解锁章节
}

type FunctionUnlockData struct {
	functionId int32 // 功能id
	Switch     bool  // switch 开关 1：开启状态
	abtest     bool  // abtest 开关 1：开启状态
	unlock     bool  // 功能开关 1：开启状态
}

type kvkProtoStatus int

const (
	NONE          = 0 // 未定义的
	OPENING       = 1 // 开启中
	FINISHED      = 2 // 结束
	VOTING        = 3 // 投票中(投下一赛季的)
	VOTE_FINISHED = 4 // 投票结束，待匹配
	MATCH         = 5 // 匹配中(GM已操作，但没有发布)
	PRE_OPEN      = 6 // 预热(匹配发布，vkv开启，但未到开始时间，不可进入k服)
)

type KvkEnterGameInfo struct {
	status        kvkProtoStatus
	currentSeason int32 // 当前赛季
	kvkServerId   int32 // K服ID
}

type StagePassInfo struct {
	passId     int32 // 关卡ID
	enemyOrder int32 // 敌人顺序
	isPass     bool  // 是否通关
	levelId    int32 // 关卡配置ID
}

type PveSkillInfo struct {
	skillId int32
	value   int32
}

type PveHeroInfo struct {
	heroId int32
	atk    float32
}

type PveCrashSave struct {
	id       int32 // passId
	time     int32
	level    int32
	exp      int32
	dropStr  string
	skill    []PveSkillInfo // 技能信息
	progress int32
	strId    string        // 唯一Id
	hero     []PveHeroInfo // pve 副本英雄信息
	hp       int32
	maxHp    int32
}

type DungeonPassInfo struct {
	pvePassInfo      PvePassInfo  // 关卡信息
	dungeonRecord    int32        // 最高关卡记录
	dungeonMaxRecord int32        // 最高解锁关卡记录(屏蔽)
	saveInfo         PveCrashSave // 大地图副本暂存信息
}

type PosHero struct {
	pos    int32
	heroId int32
}

type FormationArmyInfo struct {
	armyId int32
	num    int32
}

type FormationInfo struct {
	metaId   int32               // 编队配置ID
	hero     []PosHero           // 英雄站位信息
	armyInfo []FormationArmyInfo // 士兵数量信息
}

type VillagerEventInfo struct {
	id         string // id
	eventId    int32  // villager_event_task metaId
	villagerId int32
	questInfo  QuestInfo
}

type PlayerVillagerInfo struct {
	villagerId []int32 // 拥有的npcId
	eventInfo  []VillagerEventInfo
}

type ExploreQuestBrief struct {
	eventLv      int32 // 当前悬赏任务等级
	exp          int32 // 当前等级内的经验值
	nextTickTime int64 // 下次tick刷新时间
	eventLvUp    bool  // 本次是否悬赏任务等级提升
}

type RelicInfo struct {
	id    int32
	level int32
}

type SuitInfo struct {
	id    int32
	level int32
}

type PlayerCollectionBookInfo struct {
	relicInfos []RelicInfo // 遗物信息
	suitInfo   []SuitInfo  // 套装信息
}

type FixedTimeApInfo struct {
	ap         int32 // 可以领取的体力
	expireTime int64 // 过期时间
	take       bool  // 是否已领取
}

type OnlineRewardInfo struct {
	onlineRewardTime int64 // 下次领取的时间
	onlineRewardId   int32 // 上次领取的time_reward配置表id
}

type User struct {
	PlayerGuid string     // 玩家guid;
	playerCity PlayerCity // 玩家所有建筑
	//customs       []CustomVar      // 客户端自定义变量
	//playerItems   []ItemInfo       // 所有物品
	//resourceInfos []ResourceInfo   // 所有资源
	//armyInfos     []ArmyInfo       // 玩家所有部队
	//areas         []int32          // 已解锁地块
	//unlockAreas   []AreaUnlockInfo // 正解锁的地块
	//sciences      []ScienceInfo    // 科技
	//quests        []QuestInfo      // 章节
	//dailyQuests   []QuestInfo      // 日常
	////  repeated AttributeInfo  attibute = 13;  // 属性
	//counters    []CounterInfo // 计数
	//millisecond int64         // 服务器时间戳 只有时间使用int64,非时间的使用string
	////  repeated int32      abtests = 16;  // ab测试
	////  repeated int32      switchs = 17;  // 功能开关
	//troops              []WorldTroopInfo     // 玩家自己的出征队伍信息
	//areaGridInfos       []AreaGridInfo       // 地格信息
	//effects             []EffectInfo         // 状态效果
	//userData            UserData             // ！！！玩家信息
	//searchRanks         []SearchRank         // 世界资源查询等级信息
	//videoRewards        []VideoReward        // 视频激励
	//heroInfos           []HeroInfo           // 英雄信息
	//buildingQueues      []BuildingQueueInfo  // 建造队列
	//completedQuestId    []int32              // 已领完奖励的任务id
	//buffAttributes      []BuffAttributeInfo  // buff属性
	//lotteryInfos        []LotteryInfo        // 抽卡信息
	//shopInfos           []ShopSimpleInfo     // 商店基本信息
	//playerActivityInfo  []PlayerActivityInfo // 玩家活动(不是全局的)信息
	//mysteryShop         MysteryShop          // 神秘商店(如果有商店才会推送)
	//rechargeInfos       RechargeInfo         // 玩家礼包信息
	//vipInfo             VipInfo              // 玩家VIP信息
	//chapterInfo         QuestChapter         // 任务章节信息
	//explorerTeams       []ExplorerTeamInfo   // 探索小队数据
	//exploreQuest        []ExploreQuestInfo   // 探索任务数据
	//equipInfos          []EquipInfo          // 玩家装备信息
	//mainQuest           []QuestInfo          // 主线任务
	//chapterPassInfo     ChapterPassInfo      // pve章节关卡信息
	//functionUnlocks     []FunctionUnlockData // 功能解锁
	//fogs                []int32              // 已解锁迷雾(也就是Grid表中的ID)
	//kvkSimpleInfo       KvkEnterGameInfo     //kvk 基本信息
	//currentServerType   int32                // 当前服务器类型 普通服/K服/副本服 等
	//sceneMetaId         int32                //玩家进入游戏服关联的场景配置ID
	//stagePassInfo       StagePassInfo        // pve阶段关卡信息
	//dungeonPassInfo     DungeonPassInfo      // 地牢副本关卡信息
	//formationInfos      []FormationInfo      // 编队信息
	//villagerInfo        PlayerVillagerInfo   // 村民信息
	//exploreBrief        ExploreQuestBrief
	//collectionBook      PlayerCollectionBookInfo // 图鉴信息
	//apInfo              FixedTimeApInfo          // 固定时间刷新的体力信息
	//onlineRewardInfo    OnlineRewardInfo         // 持续在线奖励信息
	//frozenGameIds       []int32                  // 冰雪小游戏通关Id
	//frozenGameUnlockIds []int32                  // 冰雪小游戏已解锁Id
	//unlockMarchIds      []int32                  // 解锁的行军队列id
}

func init() {
	orm.RegisterModel(new(User))
}
