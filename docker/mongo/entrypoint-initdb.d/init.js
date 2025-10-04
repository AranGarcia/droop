// Initialization script.
db.createUser(
    {
        user: "droopadmin",
        pwd: "droopadmin",
        roles: [
            {role: "readWrite", db:"droop"}
        ]
    }
);