package ent

type Interactable interface {
}

type Entity struct {
	Interactable
	Drawable
	Name        string
	Description string
	/*Location    *Location
	Inventory   *Inventory
	Stats       *Stats
	Actions     *Actions
	Effects     *Effects
	Equipment   *Equipment
	Skills      *Skills
	Items       *Items
	Quests      *Quests
	Spells      *Spells
	Class       *Class*/
}
