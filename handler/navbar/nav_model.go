package navbar

type MenuList struct {
	Name      string   `json:"name"`
	ParentID  int      `json:"parentId"`
	ID        int      `json:"id"`
	Meta      MenuMeta `json:"meta"`
	Component string   `json:"component"`
	Redirect  string   `json:"redirect"`
}

type MenuMeta struct {
	Icon         string `json:"icon"`
	Title        string `json:"title"`
	Show         bool   `json:"show"`
	HideHeader   bool   `json:"hideHeader"`
	HideChildren bool   `json:"hideChildren"`
}
