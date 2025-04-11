package helperz

import "gorm.io/gorm"

func CommitAndRollback(tx *gorm.DB, err *error) func() {
	return func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}

		if *err != nil {
			tx.Rollback()
		} else {
			*err = tx.Commit().Error
			if err != nil {
				tx.Rollback()
			}
		}
	}
}
