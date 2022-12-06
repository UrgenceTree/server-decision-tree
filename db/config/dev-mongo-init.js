db.createUser({
    user: "db-usr-test",
    pwd: "db-usr-test",
    roles: [
        {
            role: "readWrite",
            db: "decision-tree",
        },
    ],
});
