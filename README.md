# cyberus-ais-cpap-service
#
# ├── cmd/                  # Entry point(s)
# │   └── server/           # Main server startup
# │       └── main.go
# │
# ├── internal/             # Internal app logic (not exported)
# │   ├── config/           # App config (env, flags, etc.)
# │   ├── db/               # DB connection (Postgres, etc.)
# │   ├── models/           # Data models (structs)
# │   ├── repositories/     # DB access (CRUD, queries)
# │   ├── services/         # Business logic
# │   ├── controllers/      # HTTP handlers
# │   ├── utils/            # Response , etc.
# │   └── routes/           # Route registration
# │
# ├── pkg/                  # Shared utilities/helpers
# │   └── logger/           # Custom logging setup
# │
# ├── go.mod
# ├── go.sum
# └── .env                  # Environment variables