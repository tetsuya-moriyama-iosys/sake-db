db = db.getSiblingDB(process.env.MONGODB_DBNAME);

db.createUser({
  user: process.env.MONGODB_USER,
  pwd: process.env.MONGODB_PASSWORD,
  roles: [{ role: 'readWrite', db: process.env.MONGODB_DBNAME }]
});
