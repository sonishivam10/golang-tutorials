A simple RESTApi with Go and Gin

- This repo contains code to an API that provides access to a store selling grocery items.

These are the endpoints listed for the introduction purpose:

- /items
 - GET : Get a list of all the grocery items available in the store
 - POST : Add a new item from request data sent as json to the grocery items.

- /items/:id
 - GET : Get a grocery items by its ID, returning the items details as JSON.