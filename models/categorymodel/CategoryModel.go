package categorymodel

import (
	"gocrud/config"
	"gocrud/entities"
)

func Get(id string) []entities.Category {
	// menjalankan query dengan function Query
	rows, err := config.Database.Query("SELECT * FROM categories WHERE id=" + id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category

		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories
}

func GetAll() []entities.Category {
	// menjalankan query dengan function Query
	rows, err := config.Database.Query(`SELECT * FROM categories`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category

		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories
}

func Store(category entities.Category) bool {
	result, err := config.Database.Exec(`INSERT INTO categories (name, created_at, updated_at) VALUE (?, ?, ?)`, category.Name, category.CreatedAt, category.UpdatedAt)

	if err != nil {
		panic(err)
	}

	id, err_ := result.LastInsertId()

	if err_ != nil {
		panic(err_)
	}

	return id > 0
}

func Update(category entities.Category) {
	_, err := config.Database.Exec(`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`, category.Name, category.UpdatedAt, category.Id)

	if err != nil {
		panic(err)
	}
}

func Destroy(id uint) {
	_, err := config.Database.Exec(`DELETE FROM categories WHERE id = ?`, id)

	if err != nil {
		panic(err)
	}
}
