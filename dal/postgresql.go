package dal

import (
	"articlesfeedapi/domain"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func connectToPostgreSQL() *sql.DB {

	connectionString := os.Getenv("CONNECTIONSTRING")
	if connectionString == "" {
		log.Fatal("CONNECTIONSTRING is not set")
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GetArticles(sourceIDs []int, lang, pageSize, currentPage int) ([]domain.Article, error) {
	db := connectToPostgreSQL()
	var sourceIDsCondition string
	if len(sourceIDs) > 0 {
		// Convert sourceIDs slice to a comma-separated string
		sourceIDsStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sourceIDs)), ","), "[]")
		sourceIDsCondition = fmt.Sprintf(" AND a.SourceID IN (%s)", sourceIDsStr)
	}

	// Calculate the offset for pagination
	offset := (currentPage - 1) * pageSize

	query := fmt.Sprintf(`
		SELECT a.ID AS ArticleId, 
		a.SourceID, 
		CASE a.SourceID WHEN 0 THEN 'Mundo Deportivo' WHEN 1 THEN 'Diario AS' WHEN 2 THEN 'Marca' WHEN 3 THEN 'Oficial'
		WHEN 4 THEN '90min' WHEN 5 THEN 'Sport' WHEN 6 THEN 'Sky Sports' WHEN 7 THEN 'El Desmarque' ELSE 'unknown' END as SourceName,	
		a.URL,  
		a.OGImage,  		
		p.headline as PostHeadline,
		p.summary as PostSummary,
		STRING_AGG(DISTINCT nt.Name, ', ') AS NormTags,
		a.scrappeddate
		FROM Article a 
		JOIN ArticleNormTag AS art_nt ON a.ID = art_nt.ArticleID
		JOIN NormTag AS nt ON nt.ID = art_nt.NormTagID
		JOIN Post p ON a.ID = p.articleid  
		WHERE p.languageid = %d 
		%s
		GROUP BY a.id, p.headline, p.summary
		ORDER BY a.id DESC
		LIMIT %d OFFSET %d`, lang, sourceIDsCondition, pageSize, offset)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []domain.Article

	for rows.Next() {
		var a domain.Article
		var normTags string
		if err := rows.Scan(&a.ID, &a.SourceID, &a.SourceName, &a.URL, &a.ImageURL, &a.Headline, &a.Summary, &normTags, &a.ScrappedDate); err != nil {
			return nil, err
		}
		a.NormTags = strings.Split(normTags, ", ")
		if a.ImageURL == "" {
			switch a.SourceID {
			case 0:
				a.ImageURL = "https://upload.wikimedia.org/wikipedia/commons/thumb/1/11/MundoDeportivo.svg/640px-MundoDeportivo.svg.png"
			case 1:
				a.ImageURL = "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d8/Diario_AS.svg/640px-Diario_AS.svg.png"
			case 2:
				a.ImageURL = "https://upload.wikimedia.org/wikipedia/commons/thumb/1/16/Marca.svg/640px-Marca.svg.png"
			case 3:
				a.ImageURL = "https://en.wikipedia.org/wiki/Atl%C3%A9tico_Madrid#/media/File:Atletico_Madrid_2017_logo.svg"
			case 4:
				a.ImageURL = "https://upload.wikimedia.org/wikipedia/commons/thumb/4/48/Logo_Sport.svg/640px-Logo_Sport.svg.png"
			case 5:
				a.ImageURL = "https://upload.wikimedia.org/wikipedia/commons/6/6b/90min_website_logo.png"
			case 6:
				a.ImageURL = "https://upload.wikimedia.org/wikipedia/en/thumb/b/b7/Sky_Sports_logo_2020.svg/640px-Sky_Sports_logo_2020.svg.png"
			case 7:
				a.ImageURL = "https://piks.eldesmarque.com/thumbs/680/bin/2023/06/30/mediaset_espana_unifica_bajo____el_desmarque____todas_sus_ventanas_de_informacion_depo.jpg"
			}
		}
		articles = append(articles, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
