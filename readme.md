1. App will have entity user 
2. User will have following parameters
    a. user_Id
    b. user_Name
    c. user_Email
    d. user_Password
3. App will have following features
    a. POST :- /user to add a new user to DB
    b. GET :- /user/{id} to read detail of user
    c. GET :- /user to read details of all users
    d. PATCH :- /user/{id} to patch update the password with given id
    e. DELETE :- /user/{id} to delete the user with given id
    f. DELETE :- /user to delete every user
4. App will support JWT Authorization//Optional for now