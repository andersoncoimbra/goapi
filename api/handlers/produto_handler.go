package handlers

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"andersoncoimbra/goapi/utils"
	"andersoncoimbra/goapi/models"
	"strconv"
)


func GetProdutos(c echo.Context) error {
	
	db := utils.DB

	rows, err := db.Query("SELECT id, name, price FROM products ORDER BY id DESC")
	if err != nil {
		fmt.Println("Erro ao buscar produtos no banco de dados")
	}

	defer rows.Close()

	products := []models.Product{}

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			fmt.Println("Erro no Scan dos produtos do banco de dados", err.Error())
		}

		products = append(products, product)
	}

	return c.JSON(http.StatusOK, products)
}

func GetProduto(c echo.Context) error {
	
	db := utils.DB

	id := c.Param("id")

	rows, err := db.Query("SELECT id, name, price FROM products WHERE id = ?", id)
	if err != nil {
		fmt.Println("Erro ao buscar produto no banco de dados")
	}

	defer rows.Close()

	var product models.Product

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			fmt.Println("Erro no Scan do produto do banco de dados", err.Error())
		}
	}

	return c.JSON(http.StatusOK, product)
}

func CreateProduto(c echo.Context) error {
	
	db := utils.DB

	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return err
	}

	result, err := db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		fmt.Println("Erro ao inserir produto no banco de dados")
	}

	product.ID, _ = result.LastInsertId()

	return c.JSON(http.StatusCreated, product)
}

func UpdateProduto(c echo.Context) error {
	
	db := utils.DB

	product := new(models.Product)	
	if err := c.Bind(product); err != nil {
		return err
	}
	product.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)

	_, err := db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.ID)
	if err != nil {
		fmt.Println("Erro ao atualizar produto no banco de dados")
	}



	return c.JSON(http.StatusOK, product)
}

func DeleteProduto(c echo.Context) error {
	
	db := utils.DB

	id := c.Param("id")

	_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		fmt.Println("Erro ao deletar produto no banco de dados")
	}

	return c.NoContent(http.StatusNoContent)
}