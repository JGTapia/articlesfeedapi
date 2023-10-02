# ArticlesFeedAPI Documentation

ArticlesFeedAPI provides endpoints to retrieve the latest articles and sources. Below are the available endpoints.

# Endpoints

# `/feed` Endpoint

This endpoint retrieves a feed of articles based on optional query parameters.

## Request

- **Method:** GET
- **Endpoint:** `/feed`

### Query Parameters 

- `sources` (String, Optional, Default: all sourdes): A comma-separated list of source IDs to filter articles from multiple sources.
- `page` (Number, Optional, Default: 1): The page number for paginated results.
- `size` (Number, Optional, Default: 10, Max:50): The number of articles to include per page.
- `lang` (Number, Optional, Default: 1): The language identifier for filtering articles by language.

**Example Request:**

```http
GET http://localhost:8084/feed
GET http://localhost:8084/feed?sources=0,1,2,4,5&page=1&size=100&lang=1
```

## Response:
- `ID` (Number): The unique identifier of the article.
- `SourceID` (Number): The ID of the source where the article originated.
- `SourceName` (String): The name of the source.
- `URL` (String): The URL of the article.
- `Headline` (String): The headline or title of the article.
- `Summary` (String): A brief summary or description of the article.
- `OGImage` (String): The URL of the Open Graph image associated with the article.
- `NormTags` (Array of Strings): An array of normalized tags associated with the article.
- `ScrappedDate (String, ISO 8601 format): The date and time when the article was scrapped or retrieved.

**Example response:**

```json
[
    {
        "ID": 1,
        "SourceID": 1,
        "SourceName": "Source Name",
        "URL": "https://example.com/article1",
        "Headline": "Headline 1",
        "Summary": "Summary for article 1",
        "OGImage": "https://example.com/image1.jpg",
        "NormTags": ["tag1", "tag2"],
        "ScrappedDate": "2023-09-28T15:04:05Z"
    },
    // ... more articles
]
```

# `/article/{article_id}` Endpoint

Retrieves the an article's data given an ID

## Request

- **Method:** GET
- **Endpoint:** `/article/{article_id}`

### **Path Parameters:**
| Name    | Type  | Required | Description                                |
|---------|-------|----------|--------------------------------------------|
| id      | int   | Required | Article id to retrieve                     |

**Example Request:**
http://localhost:8084/article/23174

## Response

The response is a JSON array containing information about the article with ID 1111. The article object has the following properties:

- `ID` (Number): The unique identifier of the article.
- `SourceID` (Number): The ID of the source where the article originated.
- `SourceName` (String): The name of the source.
- `URL` (String): The URL of the article.
- `Headline` (String): The headline or title of the article.
- `Summary` (String): A brief summary or description of the article.
- `OGImage` (String): The URL of the Open Graph image associated with the article.
- `NormTags` (Array of Strings): An array of normalized tags associated with the article.
- `ScrappedDate` (String, ISO 8601 format): The date and time when the article was scrapped or retrieved.

**Example response:**

```json
[
    {
        "ID": 23174,
        "SourceID": 1,
        "SourceName": "Source Name",
        "URL": "https://example.com/article1",
        "Headline": "Headline 1",
        "Summary": "Summary for article 1",
        "OGImage": "https://example.com/image1.jpg",
        "NormTags": ["tag1", "tag2"],
        "ScrappedDate": "2023-09-28T15:04:05Z"
    }
]
```

# `/briefing` Endpoint

This endpoint provides a briefing of news topics related to Atlético de Madrid and Valencia.

## Request

- **Method:** GET
- **Endpoint:** `/briefing`

**Example Request:**
http://localhost:8084/briefing

## Response

The response is a JSON object containing an array of news topics with the following structure:

- `Topics` (Array): An array of news topics.

Each news topic object within the `Topics` array has the following properties:

- `Headline` (String): The headline or title of the news topic.
- `Summary` (String): A brief summary or description of the news topic.
- `ArticleIDs` (Array of Numbers): An array of source IDs associated with the news topic.

Example response:

```json
{
    "Topics": [
        {
            "Headline": "Preparativos y ausencias para el partido contra el Valencia",
            "Summary": "El Atlético de Madrid se prepara para un enfrentamiento crucial contra el Valencia. Un punto destacado es la ausencia de Rodrigo de Paul, quien no podrá enfrentarse a su exequipo en su regreso a Mestalla. Además, hay incertidumbre sobre las tácticas y la alineación, ya que Simeone no ha dado indicios concretos sobre su estrategia para el partido.",
            "ArticleIDs": [
                22713,
                22712,
                22709
            ]
        },
        {
            (Next News Topic)
        }
    ]
}
```


#  `/sources` Endpoint
Retrieves a list of available sources. To be used to filter results on /feed

## Request

- **Method:** GET
- **Endpoint:** `/sources`

**Example Request:**
http://localhost:8084/sources

## **Response:**
A JSON array of objects, each containing the ID and name of the available sources.


**Example Response:**

```json
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
```

# Server Address
The API is hosted on http://localhost:8084