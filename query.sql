-- name: GetCurrentVersion :one
select version
from versions
where id = ?;

-- name: UpdateVersion :exec
update versions
set version = ?
where id = ?;

-- name: InsertVersion :exec
insert into versions (
    id, version
) values (
    ?, ?
);