package initialize

import (
	"context"

	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/consts"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/utils/crypto"
)


func InitSeedData() {
	query := `INSERT INTO role (role_name, role_slug, role_description) VALUE (?, ?, ?)`
	_, err := global.Mdb.ExecContext(context.Background(), query, consts.ADMIN, crypto.GetHash(consts.ADMIN), consts.RoleDesc[consts.ADMIN])
	if err != nil {
		panic(err)
	}

	_, err = global.Mdb.ExecContext(context.Background(), query, consts.MANAGER, crypto.GetHash(consts.MANAGER), consts.RoleDesc[consts.MANAGER])
	if err != nil {
		panic(err)
	}

	_, err = global.Mdb.ExecContext(context.Background(), query, consts.EDITOR, crypto.GetHash(consts.EDITOR), consts.RoleDesc[consts.EDITOR])
	if err != nil {
		panic(err)
	}
	_, err = global.Mdb.ExecContext(context.Background(), query, consts.VIEWER, crypto.GetHash(consts.VIEWER), consts.RoleDesc[consts.VIEWER])
	if err != nil {
		panic(err)
	}

	_, err = global.Mdb.ExecContext(context.Background(), query, consts.SUPPORT, crypto.GetHash(consts.SUPPORT), consts.RoleDesc[consts.SUPPORT])
	if err != nil {
		panic(err)
	}
}
