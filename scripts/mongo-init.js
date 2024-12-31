db = db.getSiblingDB("admin");
db = db.getSiblingDB("service-media");
db.createUser({
    user: "user",
    pwd: "pvs1909~",
    roles: [
      {
        role: "readWrite",
        db: "service-media"
      }
    ]
});
  
