//.envの値が読み込めないようなので暫定処理 デプロイ時にどうするか考える
const MONGODB_DBNAME="sakedb"
const MONGODB_USER="root"
const MONGODB_PASSWORD="root"

db = db.getSiblingDB(MONGODB_DBNAME);

db.createUser({
  user: "admin",
  pwd: "admin",
  roles: [{ role: 'root', db: "admin" }]
});

db.createUser({
  user: MONGODB_USER,
  pwd: MONGODB_PASSWORD,
  roles: [{ role: 'readWrite', db: MONGODB_DBNAME }]
});
