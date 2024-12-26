package repository

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (r *UserRepository) AddAdmin(ctx context.Context, res userModel.AdminResponse) (int, error) {
	var adminID int

	err := r.psqlDB.QueryRow(
		ctx,
		addAdminQuery,
		res.Name,
		res.Surname,
		res.Username,
		res.Password,
	).Scan(&adminID)

	if err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return adminID, nil
}

func (r *UserRepository) ListAdmins(ctx context.Context, paginationData abstract.PaginationData) (
	[]*userModel.AdminData, error,
) {
	var adminDatas []*userModel.AdminData
	offset := (paginationData.CurrentPage - 1) * paginationData.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&adminDatas,
		listAdminsQuery,
		offset,
		paginationData.Limit,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return adminDatas, nil
}

func (r *UserRepository) CountAdmins(ctx context.Context) (int, error) {
	var totalCount int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalCount,
		totalAdminCountQuery,
	)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalCount, nil
}

func (r *UserRepository) GetAdmin(ctx context.Context, adminID int) (*userModel.AdminData, error) {
	var adminData userModel.AdminData

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&adminData,
		getAdminQuery,
		adminID,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &adminData, nil
}

func (r *UserRepository) DeleteAdmin(ctx context.Context, adminID int) error {
	result, err := r.psqlDB.Exec(
		ctx,
		deleteAdminQuery,
		adminID,
	)

	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	if result.RowsAffected() == 0 {
		return errlst.NewNotFoundError("[userRepository][DeleteAdmin]: admin not found")
	}

	return nil
}

func (r *UserRepository) UpdateAdmin(ctx context.Context, updateData userModel.AdminUpdateData) (string, error) {
	var updateRes string

	err := r.psqlDB.QueryRow(
		ctx,
		updateAdminQuery,
		updateData.Name,
		updateData.Surname,
		updateData.UserName,
		updateData.Password,
		updateData.ID,
	).Scan(&updateRes)

	if err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return updateRes, nil
}

func (r *UserRepository) AdminFindStudent(ctx context.Context, filterStudent userModel.FilterStudent,
	paginationQuery abstract.PaginationData,
) (
	[]*userModel.AllStudentData, error,
) {
	var allStudentDataWithFilter []*userModel.AllStudentData
	offset := (paginationQuery.CurrentPage - 1) * paginationQuery.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&allStudentDataWithFilter,
		adminFindStudentQuery,
		filterStudent.StudentName,
		filterStudent.StudentSurname,
		offset,
		paginationQuery.Limit,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return allStudentDataWithFilter, nil
}
