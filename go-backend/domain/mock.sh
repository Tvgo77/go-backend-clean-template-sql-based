# Run following command to automatically generate mock of interface

cd /workspaces/go-backend/domain
mockgen -source=user_repository_domain.go -destination ./mock/user_repository_mock.go