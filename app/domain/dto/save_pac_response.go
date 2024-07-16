package dto

// SavePacResponse Define a struct for your response
type SavePacResponse struct {
	Success     string  `json:"success" example:"1"`
	RequestID   string  `json:"requestID" example:"REQ-001"`
	PackageList Package `json:"package_list" `
}

type Package struct {
	Name           string       `json:"name"`
	RequestID      string       `json:"requestID"`
	PkgImgKey      string       `json:"pkg_img_key"`
	BranchID       string       `json:"branch_id"`
	PackageID      string       `json:"package_id"`
	Description    string       `json:"description"`
	OwnerID        string       `json:"owner_id"`
	SubPkgMenuList []SubPkgMenu `json:"sub_pkg_menu_list"`
	MenuList       []Menu       `json:"menu_list"`
}
