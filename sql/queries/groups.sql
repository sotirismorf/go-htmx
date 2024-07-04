-- name: CreateGroup :exec
INSERT INTO groups (name, location) VALUES ($1, $2);

-- name: SelectGroups :many
select groups.id, groups.name, places.name as city
from groups
left join places on places.id = groups.location
order by groups.name;
