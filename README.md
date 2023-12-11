# Music API
This repository contains a simple API for retrieving and managing music information, including tracks and artists. It leverages the Echo framework for building the API and integrates with a Spotify service to fetch track metadata.

## Endpoints
### 1. Get Music by ISRC
#### Endpoint
```
GET /music/{isrc}
```

##### Description
Retrieve music information by its ISRC (International Standard Recording Code).

##### Parameters
- isrc (path, string, required): ISRC of the music.
  
##### Responses
- 200 OK: Successfully retrieved music information (array of objects).
- 400 Bad Request: ISRC not found or invalid.
- 404 Not Found: Music not found.
- 500 Internal Server Error: Internal server error.
  
### 2. Get Music by Artist Name
#### Endpoint
```
GET /music/artist/{name}
```

##### Description
Retrieve music information by the name of the artist.

##### Parameters
- name (path, string, required): Name of the artist.
  
##### Responses
- 200 OK: Successfully retrieved music information (array of objects).
- 400 Bad Request: Name cannot be empty.
- 404 Not Found: Music not found.
- 500 Internal Server Error: Internal server error.
  
### 3. Insert Music Track Metadata
#### Endpoint
```
POST /tracks/{isrc}
```

#### Description
Retrieve track information from Spotify based on the provided ISRC and insert it into the database.

#### Parameters
- isrc (path, string, required): ISRC of the track.
  
Request Body
```
{
  "isrc": "string"
}
```

#### Responses
- 201 Created: Track inserted successfully.
- 400 Bad Request: ISRC not found or invalid.
- 404 Not Found: Track not found.
- 409 Conflict: The song already exists in the database.
- 500 Internal Server Error: Internal server error.
  
#### Implementation Details
1. Get Music by ISRC
The HandlerGetByISRC function retrieves music information by ISRC. It uses the provided ISRC to query the database through the MusicService contract.

2. Get Music by Artist Name
The HandleGetByArtistName function retrieves music information by artist name. It uses the provided artist name to query the database through the MusicService contract.

3. Insert Music Track Metadata
The HandlerInsertTrack function inserts track metadata into the database. It first validates the request body, fetches track details from Spotify using the SpotifyIntegration contract, and then inserts the data into the database using the MusicService contract.

#### Dependencies
Echo: A fast and minimalist Go web framework.
Spotify API: An external service for integrating with Spotify to fetch track metadata. 
PostgreSQL: SQL Database.

#### Getting Started
Clone this repository. Install dependencies using go get. Configure the database connection and Spotify integration. Build and run the application.

To authenticate and access protected features, obtain a JWT token by making a request to the [authentication service (/authenticate)](https://github.com/paulojunior/authentication_service/) . Include the generated token in the headers of subsequent requests to authorized endpoints.
