// use admin;

// db.createUser(
//     {
//         user: "mongo",
//         pwd: "mongo",
//         roles: [
//             "userAdminAnyDatabase",
//             "readWriteAnyDatabase"
//         ]

//     }
// )

// mongo --help
// mongo -u mongo -p mongo

// db.createUser(
//     {
//         user: "admin",
//         pwd: "12345",
//         roles: [
//             { role: "readWrite", db: "backendusermongo"}
//         ]
//     }
// )

// use backendusermongo;
// mongo -u admin -p 12345 --authenticationDatabase backendusermongo