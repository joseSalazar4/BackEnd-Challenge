# LTV BackEnd Challenge 
##### By: Jose Pablo Salazar

All of the tasks were completed properly. On this document you will find various elements that will give a better understanding of how the work was done, errors that appeared in the middle and some suggestions as requested.

#### Table of Contents
- Main Tasks 
- Extra Credit Tasks
- Dependencies
- Set Up 
- Installation 
- Suggestions 

## Main Tasks

##### Connection to database

Since the database is SQLite the file had to be retrieved only once from the S3 bucket provided. Once obtained, little inspection was needed to understand the schema. A SELECT * FROM songs gave a better understanding of the data, showing some null values which would be important later on the development. 


> DB Browser was the tool used to visually see the data inside the db
> This allowed for queries to be executed there without having to set them on the app

After the database seemed to be the right one then the coding started!

##### After one falls, the others will too...
Since the structure of the API is very similar, as soon as we could have one endpoint working the others would not take long to become available. Always thinking into the details that will arrive later on.

The best start would be just to retrieve all the songs without filtering anything since the DB connection could be tested out and the response too. The approach was to have a fast but reliable start, which is what generating that test in the main.go file was intended for. After some tweaks the data was finally being retrieved and connection established properly.

##### Order is your ally in any scenario, use it!
This is what followed up. Organization will give me an advantage and the earlier a sketch is converted into a proper drawing the better. Next commit will be creating the model and controller folders to distribute the current code into proper files and their proper locations.

##### Set up done, start doing GETs
Now, the first route was created and after testing its functionality it was only sending "problem retriving songs data"  which was not a surpirse since there was no other c.JSON() apart from the one with the error. That is why copy paste can sometimes be you enemy. Or maybe just being careful could be enough. After noticing that, the route started to have a better shape. That did not took long so the model was incorporated into the route to make it much more down to earth. It worked!

##### And now they started falling! Or maybe arriving(?)
Now that this route was fully functioning the others would be pretty similar, it would just be a matter of focus on which ones to create first. Of course the main tasks were the ideal path, which is what came along apart from a constant of the SQL statement which was more than obvious to be needed.

## Extra Credit Tasks
So far, the API endpoints were pretty direct and simple. These two other tasks seemed a little more complicated but nothing that a coffee and a small break wouldn't solve. 

The following steps were to test some SQL queries inside the GUI mentioned before to test out some syntax and make sure it would work. Then the last two tasks got done, one would only need to be inside of a range with WHERE clause and the other would need to be ORDERED BY and then use some basic functions such as COUNT() and SUM() on the SELECT section at the top.

This were the elements needed to make the DB to the hardwork for us. 
##### Better late than Never!
Everything was working by then, but the code was prone to SQL Injection which seemed a bad thing not to have into consideration since the beginning. Some refactoring was done but nothing that would take too much time or testing.

## Dependencies:
- [Gin-Gonic] - HTTP web framework, It features a Martini-like API with much better performance 
- [Testify] - A  toolkit with common assertions and mocks that plays nicely with the standard library
- [Go-SQLite3] - sqlite3 driver for go using database/sql

## Setup
To setup the environment to get the application working these are the following steps:
1. Download GO: https://go.dev/doc/install and install a GCC compiler. Here you can find TDM-GCC: https://jmeubank.github.io/tdm-gcc/ which is what I used in Windows 10. Linux already comes with GCC, to check if you have it properly installed open a terminal and type: gcc --version and it should display some information.
3. Clone the following Repo: https://github.com/joseSalazar4/BackEnd-Challenge
4. Once cloned you will notice the DB is not be available. Please ask your manager to get you a copy of it and change its name to this exactly: "songsDatabaseLTV.db"
5. The project is now almost ready to be executed!

## Installation
Now open a terminal inside the folder and do the following commands inside of it.

```sh
go build
```

The build command will start compiling the code and setting up the exec. If there are no errors found then continue with the following command. If you encounter an issue with the gcc compiler make sure to add it to the path on windows so it can be reachable from anywhere. 

```sh
go run main.go
```
This _run_ command will start the app and a message could pop saying the file wants access to the network, click on accept. 

Congratulations! You have the API Running! Oh where? Important detail... 
Link: http://localhost:8080/songs there you go.

Available routes can be seen in the code, but for ease:
- Retrieve ALL songs (not requested in the challenge): http://localhost:8080/songs
- http://localhost:8080/songsByGenre/:genre
- http://localhost:8080/songsByArtist/:artist
- http://localhost:8080/songsBySongName/:songName
- http://localhost:8080/songsByLength?minLen=1&maxLen=100
- http://localhost:8080/songsSortedByGenre
Names are pretty self explanatory.


## Suggestions

| Suggestion | Description |
| ------ | ------ |
| Naming  | The song table should have songName o just name instead of song since it can create confusion in the queries. |
| Quantity | The artist name and song name should not be so long since space can be wasted. Here the _song_  field can generate confusion as if it is the lyrics of the song due to its length. |
| Artist field | The artist field could be an id so the artist can have a table of their own and have other fields that would allow for escalability. |
| Index fields | The field to index by would be the genres since they are a small dataset and there won't be a compromise in storage. The other fields if indexed will make faster queries but with a enormous impact in storage and in the long run is a bad decision since it is expensive to maintain. |
| Naming length | The length is an integer which could be in minutes, seconds or any other unit. A better name could be lengthMinutes or even better lengthSeconds which could allow for a more precise value |

Thanks for this opportunity!

   [Gin-Gonic]: <https://github.com/gin-gonic/gin>
   [Go-SQLite3]: <https://github.com/mattn/go-sqlite3>
   [Testify]: <https://github.com/stretchr/testify>
