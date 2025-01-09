package server

import (
	"github.com/NodiumHosting/VaultMapperSyncServer/models"
	pb "github.com/NodiumHosting/VaultMapperSyncServer/proto"
)

func getActiveVaults() int {
	var i int
	HUB.Vaults.Range(func(k, v interface{}) bool {
		i++
		return true
	})
	return i
}

func getActiveConnections() int {
	var i int
	HUB.Vaults.Range(func(k, v interface{}) bool {
		vault := v.(*Vault)
		vault.Connections.Range(func(k, v interface{}) bool {
			i++
			return true
		})
		return true
	})
	return i
}

func getActiveCells() int {
	var i int
	HUB.Vaults.Range(func(k, v interface{}) bool {
		vault := v.(*Vault)
		vault.Cells.Range(func(k, v interface{}) bool {
			i++
			return true
		})
		return true
	})
	return i
}

func getActiveRooms() int {
	var i int
	HUB.Vaults.Range(func(k, v interface{}) bool {
		vault := v.(*Vault)
		vault.Cells.Range(func(k, v interface{}) bool {
			cell := v.(*pb.VaultCell)
			if cell.CellType == pb.CellType_CELLTYPE_ROOM {
				i++
			}
			return true
		})
		return true
	})
	return i
}

func GetTotalPlayerCount() (int64, error) {
	var count int64
	err := DB.Model(&models.Player{}).Count(&count).Error
	return count, err
}

func GetPlayerCountInVault(vaultID string) (int64, error) {
	var count int64
	err := DB.Model(&models.PlayerVault{}).Where("vault_id = ?", vaultID).Count(&count).Error
	return count, err
}

func GetTotalVaults() (int64, error) {
	var count int64
	err := DB.Model(&models.Vault{}).Count(&count).Error
	return count, err
}

func GetTotalRooms() (int64, error) {
	var count int64
	err := DB.Model(&models.VaultCell{}).Where("cell_type = ?", pb.CellType_CELLTYPE_ROOM).Count(&count).Error
	return count, err
}

// GetLargestVault returns the vault with the most cells in it
/*func GetLargestVault() (int64, error) {
	var count int64
	err := DB.Model(&models.VaultCell{}).Select("vault_id").Group("vault_id").Count(&count).Error
	return count, err
}*/

func GetLargestVault() (int64, error) {
	var result struct {
		VaultID string
		Count   int64
	}
	err := DB.Model(&models.VaultCell{}).
		Select("vault_id, COUNT(*) as count").
		Group("vault_id").
		Order("count DESC").
		Limit(1).
		Scan(&result).Error
	if err != nil {
		return 0, err
	}
	return result.Count, nil
}

/*func GetBiggestParty() (int64, error) {
	var count int64
	err := DB.Model(&models.PlayerVault{}).Select("player_uuid").Group("vault_id").Count(&count).Error
	return count, err
}*/

func GetBiggestParty() (int64, error) {
	var result struct {
		VaultID string
		Count   int64
	}
	err := DB.Model(&models.PlayerVault{}).
		Select("vault_id, COUNT(player_uuid) as count").
		Group("vault_id").
		Order("count DESC").
		Limit(1).
		Scan(&result).Error
	if err != nil {
		return 0, err
	}
	return result.Count, nil
}

// GetTotalRoomsBasic returns the total number of basic rooms in db
func GetTotalRoomsBasic() (int64, error) {
	var count int64
	err := DB.Model(&models.VaultCell{}).Where("room_type = ? & cell_type = ?", pb.RoomType_ROOMTYPE_BASIC, pb.CellType_CELLTYPE_ROOM).Count(&count).Error
	return count, err
}

func GetTotalRoomsOre() (int64, error) {
	var count int64
	err := DB.Model(&models.VaultCell{}).Where("room_type = ?", pb.RoomType_ROOMTYPE_ORE).Count(&count).Error
	return count, err
}

func GetTotalRoomsChallenge() (int64, error) {
	var count int64
	err := DB.Model(&models.VaultCell{}).Where("room_type = ?", pb.RoomType_ROOMTYPE_CHALLENGE).Count(&count).Error
	return count, err
}

func GetTotalRoomsOmega() (int64, error) {
	var count int64
	err := DB.Model(&models.VaultCell{}).Where("room_type = ?", pb.RoomType_ROOMTYPE_OMEGA).Count(&count).Error
	return count, err
}

// GetActivity returns a list of vaults with the players inside them
func GetActivity() map[string][]string {
	vaults := make(map[string][]string)
	HUB.Vaults.Range(func(k, v interface{}) bool {
		vaultID := k.(string)
		vault := v.(*Vault)
		var players []string
		vault.Connections.Range(func(k, v interface{}) bool {
			players = append(players, k.(string))
			return true
		})
		vaults[vaultID] = players
		return true
	})

	return vaults
}
