namespace go pkg

struct User {
    1:required i32 id;
    2:required string name;
    3:required i32 Age;
    4:required string address;
    5:required string mobile;
}

service UserService {
    User GetUser(
        1:required User user
    )
}
