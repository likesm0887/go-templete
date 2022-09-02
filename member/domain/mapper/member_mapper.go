package mapper

import (
	"awesomeProject1/member/adapter/contract"
	"awesomeProject1/member/adapter/repository/data"
)


// region to repository data


func ToDataInfo(information contract.Information) data.Information {
	return data.Information{
		UserID: information.UserID,
		UserName: data.UserName{
			Name: data.Name{
				FirstName: information.UserName.Name.FirstName,
				LastName:  information.UserName.Name.LastName,
			},
			NickName: information.UserName.NickName,
		},
		Phone:    information.Phone,
		Photo:    information.Photo,
		Birthday: information.Birthday,
		AddressObject: data.AddressObject{
			Address:    information.AddressObject.Address,
			PostalCode: information.AddressObject.PostalCode,
		},
	}
}


// endregion

// region to contract

func ToContractInfo(information data.Information) contract.Information {
	return contract.Information{
		UserID: information.UserID,
		UserName: contract.UserName{
			Name: contract.Name{
				FirstName: information.UserName.Name.FirstName,
				LastName:  information.UserName.Name.LastName,
			},
			NickName: information.UserName.NickName,
		},
		Photo:    information.Photo,
		Phone:    information.Phone,
		Birthday: information.Birthday,
		AddressObject: contract.AddressObject{
			Address:    information.AddressObject.Address,
			PostalCode: information.AddressObject.PostalCode,
		},
	}
}



// endregion
