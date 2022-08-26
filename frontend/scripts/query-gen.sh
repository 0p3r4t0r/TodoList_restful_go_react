# Generate swagger from api: https://github.com/swaggo/swag#swag-cli
cd ../backend
swag init --parseDependency --parseInternal # --dir ./backend --output ./backend/docs
cd -

# Copy files over to the frontend
cp ../backend/docs/swagger.json ../backend/docs/swagger.yaml ./specs/

# Generate queries
pnpm exec orval