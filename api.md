# ArticlesFeedAPI Documentation

ArticlesFeedAPI provides endpoints to retrieve the latest articles and sources. Below are the available endpoints.

## Endpoints

### 1. Get Latest Articles

#### **Endpoint:**
GET /feed

#### **Description:**
Retrieves the latest articles.

#### **Query Parameters:**
| Name    | Type  | Required | Description                                |
|---------|-------|----------|--------------------------------------------|
| sources | []int | Optional | Sources IDs (Default all sources)          |
| lang    | int   | Optional | Post language (Default 1)                   |
| page    | int   | Optional | Current page (Default 1)                   |
| size    | int   | Optional | Page size (Default 10, max 50)             |

#### **Example Request:**
http://localhost:8084/feed?sources=0,1,2,4,5&lang=1&page=2&size=10

#### **Response:**
A JSON array of objects, each containing details of an article including the article's ID, source ID, source name, URL, headline, summary, image url, normalized tags, and the scrapped date.

#### **Example Response:**
[
    {
        "ID": 1,
        "SourceID": 1,
        "SourceName": "Diario AS",
        "URL": "https://example.com/article1",
        "Headline": "Headline 1",
        "Summary": "Summary for article 1",
        "OGImage": "https://example.com/image1.jpg",
        "NormTags": ["tag1", "tag2"],
        "ScrappedDate": "2023-09-28T15:04:05Z"
    },
    // ... more articles
]

### 2. Get Avalible Sources

#### **Endpoint:**
GET /sources

#### **Description:**
Retrieves a list of available sources. To be used to filter results on /feed

#### **Parameters:**
None

#### **Response:**
A JSON array of objects, each containing the ID and name of the available sources.

#### **Example Response:**
[
    {
        "ID": 0,
        "Name": "Mundo Deportivo"
    },
    {
        "ID": 1,
        "Name": "Diario AS"
    },
    // ... more sources
]

## Server Address
The API is hosted on http://localhost:8084.