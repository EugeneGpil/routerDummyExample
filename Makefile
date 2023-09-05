dev:
	docker compose up --build --remove-orphans --detach --force-recreate;
	code --folder-uri vscode-remote://attached-container+$$(printf "rde-golang-1" | xxd -p)/var/www/back

exec:
	docker compose exec golang bash

exec-root:
	docker compose exec --user=root golang bash

stop:
	docker compose stop
