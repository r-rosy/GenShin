package tables

type User struct {
	Uid       string `gorm:"column:uid;PRIMARY_KEY;"`
	Password  string `gorm:"column:password;"`
	Username  string `gorm:"column:username;"`
	Account   string `gorm:"column:account;"`
	HeadImage string `gorm:"column:head_image;"`
	Signature string `gorm:"column:signature;"`
	Rough     int    `gorm:"column:rough;"`
	Mora      int    `gorm:"column:mora;"`
	Number1   int    `gorm:"column:number1;"`
	Number2   int    `gorm:"column:number2;"`
}
type Weapon struct {
	Id          int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;"`
	Rank        int    `gorm:"column:rank;"`
	Name        string `gorm:"column:name;"`
	WeaponId    int    `gorm:"column:weapon_id;"`
	RefineRange int    `gorm:"column:refine_range;"`
	Grade       int    `gorm:"column:grade;"`
	Account     string `gorm:"column:account;"`
}
type Character struct {
	Id              int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;"`
	Account         string `gorm:"column:account;"`
	Name            string `gorm:"column:name;"`
	Rank            int    `gorm:"column:rank;"`
	WeaponId        int    `gorm:"column:weapon_id;"`
	Grade           int    `gorm:"column:grade;"`
	GodLeftValueId1 int    `gorm:"column:god_left_value_id1;"`
	GodLeftValueId2 int    `gorm:"column:god_left_value_id2;"`
	GodLeftValueId3 int    `gorm:"column:god_left_value_id3;"`
	GodLeftValueId4 int    `gorm:"column:god_left_value_id4;"`
	GodLeftValueId5 int    `gorm:"column:god_left_value_id5;"`
	RefineRange     int    `gorm:"column:god_left_value_id6;"`
}
type GodLeftValue struct {
	Id      int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;"`
	Account string `gorm:"column:account;"`
	Grade   int    `gorm:"column:grade;"`
	Rank    int    `gorm:"column:rank;"`
	Name    string `gorm:"column:name;"`
	Index   int    `gorm:"column:index;"`
}
type Cookie struct {
	Id      int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;"`
	Name    string `gorm:"column:name;"`
	Value   string `gorm:"column:value;"`
	Account string `gorm:"column:account;"`
}
