Welcome to gallus gator feed aggregator.

REQUIREMENTS:
- Postgres v15 or later
- Go

HOW TO INSTALL:
On the command line you use, type "go install" (without the quotation mark)

SETTING UP CONFIG FILE:
1. Go into your root directory.
2. Create a file named ".gatorconfig.json"
3. Inside the file create the following content
    `{"db_url": "postgres://example"}`
4. Replace the "postgres://example" with connection string on your computer

GALLUS GATOR COMMAND LIST:
1. login\
Used to login to certain username that already created in the database\
It expect 1 username string argument\
Example: login gallus

2. register\
Used to create a certain username that has not been in the database\
It expect 1 username string argument\
Example: register gallus

3. users\
Used to list the users in the database\
No argument needed\
Example: users

4. addfeed\
Used to add certain feed to the database for the current user and automatically follow it\
It needs 2 strings argument, the name of the feed and the url of the feed\
Example: addfeed "Gamespot" "https://www.gamespot.com/feeds/anime-news"

5. feeds\
Used to list the detailed feed struct that the current user follow\
Example: feeds

6. follow\
Used to follow a certain feed that already in the database for the current user\
It needs 1 string argument, the url of the feed\
Example: follow "https://www.gamespot.com/feeds/anime-news"

7. unfollow\
Used to unfollow a certain feed that already in the database for the current user\
It needs 1 string argument, the url of the feed\
Example: unfollow "https://www.gamespot.com/feeds/anime-news"

8. following\
Used to list the name of the feed that the current user follow\
Example: following

9. agg\
Used to scrape and get the post list on the feeds that current user follow every specified duration\
It needs 1 duration string argument\
Example: agg 30s

10. browse\
Used to browse the posts that the current user have in the database\
It has optional number limit of posts shown (Default is 2)\
Example: browse 30

11. reset\
Used to reset the database\
(WARNING: IT WILL WIPE ALL THE DATABASE ALREADY SAVED)\
Example: reset