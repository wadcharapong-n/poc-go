package dto

// SavePackageRequest is the top-level structure representing the JSON body of a save package request
type SavePackageRequest struct {
	Name           string       `json:"name" example:"Example Name"`
	RequestID      string       `json:"requestID" example:"REQ123456"`
	PkgImgKey      string       `json:"pkg_img_key" example:"ImgKey123"`
	BranchID       string       `json:"branch_id" example:"Branch001"`
	PackageID      string       `json:"package_id" example:"Package123"`
	Description    string       `json:"description" example:"Description of the package"`
	OwnerID        string       `json:"owner_id" example:"Owner123"`
	SubPkgMenuList []SubPkgMenu `json:"sub_pkg_menu_list" `
	MenuList       []Menu       `json:"menu_list"`
}

// SubPkgMenu represents an item in the sub package menu list
type SubPkgMenu struct {
	MenuID int    `json:"menu_id" example:"123"`
	Name   string `json:"name" example:"abc"`
	Price  int    `json:"price" example:"100"`
}

// Menu represents an item in the menu list
type Menu struct {
	MenuID     int  `json:"menu_id" example:"123"`
	PkgMenuID  int  `json:"pkg_menu_id" example:"456"`
	DeleteFlag bool `json:"delete_flag" example:"false"`
}
