package repository

const (
	createNews = `INSERT INTO priceGroup (author_id, title, content, image_url, category, created_at) 
					VALUES ($1, $2, $3, NULLIF($4, ''), NULLIF($4, ''), now()) 
					RETURNING *`

	createPriceGroup = `INSERT INTO price_group_settings
	("name", description, currency_type, active, created_by)
	VALUES($1, $2, $3, $4, $5) RETURNING *`

	updateNews = `UPDATE priceGroup 
					SET title = COALESCE(NULLIF($1, ''), title),
						content = COALESCE(NULLIF($2, ''), content), 
					    image_url = COALESCE(NULLIF($3, ''), image_url), 
					    category = COALESCE(NULLIF($4, ''), category), 
					    updated_at = now() 
					WHERE priceGroup_id = $5
					RETURNING *`

	updatePriceGroup = `UPDATE price_group_settings
					SET name = $1,
						description = $2, 
						currency_type = $3,
						active = $4,
						updated_by = $5
						WHERE id = $6 RETURNING *`

	getNewsByID = `SELECT *
FROM price_group_settings
WHERE id = $1`

	getAllPriceGroup = `SELECT *
FROM price_group_settings`

	deleteNews = `DELETE FROM priceGroup WHERE priceGroup_id = $1`

	deletePriceGroup = `UPDATE price_group_settings 
	SET 
		deleted_at = now() 
	WHERE id = $1
	RETURNING *`

	getTotalCount = `SELECT COUNT(id) FROM price_group_settings`

	getNews = `SELECT priceGroup_id, author_id, title, content, image_url, category, updated_at, created_at 
				FROM priceGroup 
				ORDER BY created_at, updated_at OFFSET $1 LIMIT $2`
	getPriceGroup = `SELECT * 
				FROM price_group_settings 
				ORDER BY name, updated_at OFFSET $1 LIMIT $2`

	findByTitleCount = `SELECT COUNT(*)
					FROM priceGroup
					WHERE title ILIKE '%' || $1 || '%'`

	findByTitle = `SELECT priceGroup_id, author_id, title, content, image_url, category, updated_at, created_at
					FROM priceGroup
					WHERE title ILIKE '%' || $1 || '%'
					ORDER BY title, created_at, updated_at
					OFFSET $2 LIMIT $3`
)
