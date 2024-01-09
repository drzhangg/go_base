package db

import (
	"gorm.io/gorm"
)

type Heroes struct {
	Id        int
	Name      string
	AliasName string
	Attribute int
	Position  int
}

type Against struct {
	Id          int
	OtherSide   string
	OurSide     string
	Probability int
}

type Db struct {
	*gorm.DB
}

var (
	NameToId = make(map[string]int)
	IdToName = make(map[int]string)
)

func (db *Db) GetHeroByName(name string) (Heroes, error) {
	hero := Heroes{}
	err := db.Table("heros").First(&hero, "name = ?", name).Error
	if err != nil {
		return Heroes{}, err
	}
	return hero, nil
}

func (db *Db) GetHeroById(id int) (Heroes, error) {
	hero := Heroes{}
	err := db.Table("heros").First(&hero, "id = ?", id).Error
	if err != nil {
		return Heroes{}, err
	}
	return hero, nil
}

func (db *Db) GetAllHeroes() ([]Heroes, error) {
	heroes := []Heroes{}

	err := db.Table("heros").Find(&heroes).Error
	if err != nil {
		return nil, err
	}

	return heroes, nil
}

func (db *Db) CreateHero(name, aliasName string, attribute, position int) error {
	hero := Heroes{
		Name:      name,
		AliasName: aliasName,
		Attribute: attribute,
		Position:  position,
	}
	err := db.Table("heros").Select("name", "alias_name", "attribute", "position").Create(&hero).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateSideData 创建对阵数据
func (db *Db) CreateSideData(otherSide, ourSide string,probability int) error {
	against := Against{
		OtherSide:   otherSide,
		OurSide:     ourSide,
		Probability: probability,
	}
	err := db.Table("probability").Select("other_side", "our_side", "probability").Create(&against).Error
	if err != nil {
		return err
	}

	return nil
}

// GetSideData 查询对阵数据
func (db *Db) GetSideData(otherSide string) (Against, error) {

	against := Against{}

	err := db.Table("against").First(&against, "other_side = ? ", otherSide).Error
	if err != nil {
		return Against{}, err
	}

	return against, nil
}

func (db *Db) initData() error {
	heroes, err := db.GetAllHeroes()
	if err != nil {
		return err
	}
	for _, val := range heroes {
		NameToId[val.Name] = val.Id
		IdToName[val.Id] = val.Name
	}

	return nil
}

