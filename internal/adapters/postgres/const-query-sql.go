package postgres

const (
	checkTeamQuery               = `SELECT EXISTS (SELECT 1 FROM team WHERE name = $1)`
	checkUserRoleQuery           = `SELECT EXISTS ( SELECT 1 FROM public."user" WHERE id = $1 AND role_id = $2) `
	checkUserIsExistByEmailQuery = `SELECT COALESCE((SELECT id FROM "user" WHERE email = $1), 0);`
	createTeamQuery              = `INSERT INTO team (name) VALUES ($1) RETURNING id`
	setRoleByUserIdQuery         = `UPDATE "user" SET role_id = $1 WHERE id = $2`
	checkUserHasTeamByIdQuery    = `SELECT EXISTS (SELECT 1 FROM "user" WHERE (id = $1) AND (team_id IS NOT NULL));`
	checkUserHasTeamByEmailQuery = `SELECT EXISTS (SELECT 1 FROM "user" WHERE (email = $1) AND (team_id IS NOT NULL));`
	setTeamIdQuery               = `UPDATE "user" SET team_id = $1 WHERE id = $2`
	updateTeamNameQuery          = `UPDATE team SET name = $1 WHERE id = (SELECT team_id FROM "user" WHERE id = $2)`
	deleteMemberFromTeamQuery    = `UPDATE "user" SET team_id = NULL WHERE id = $1`
	addMemberToTeamQuery         = `UPDATE "user" SET team_id = $1 WHERE id = $2`
	getTeamByIdQuery             = `SELECT id , name,
    								COALESCE((SELECT SUM(points) FROM team_text_answer WHERE team_id = 1), 0) +
       								COALESCE((SELECT SUM(points) FROM team_media_answer WHERE team_id = 1), 0) + team.delta_points AS total_points
    								FROM team
									WHERE id = (SELECT team_id FROM "user" WHERE id = $1)`
	getMembersTeamQuery         = `SELECT id, role_id, name, "group", email FROM "user" WHERE team_id =$1;`
	deleteTeamQuery             = `DELETE FROM team WHERE team.id = $1`
	getTotalPointsByTeamIdQuery = `SELECT
    							   COALESCE((SELECT SUM(points) FROM team_text_answer WHERE team_id = $1), 0) +
    							   COALESCE((SELECT SUM(points) FROM team_media_answer WHERE team_id = $1), 0) AS total_points`
	setDeltaPointsByTeamIdQuery = `UPDATE team SET delta_points = delta_points + $1 WHERE id = $2`
	getCountUserInTeamQuery     = `SELECT COUNT(*) FROM "user" WHERE team_id = $1`
	getTeamByCountUsersQuery    = `SELECT  id , name,
       							   COALESCE((SELECT SUM(points) FROM team_text_answer WHERE team_id = 1), 0) +
       							   COALESCE((SELECT SUM(points) FROM team_media_answer WHERE team_id = 1), 0) + delta_points AS total_points 
								   FROM team WHERE id IN (SELECT team_id FROM "user" WHERE team_id IS NOT NULL
                                							GROUP BY team_id HAVING COUNT(*) = $1);`
	getAllTeamQuery = `SELECT id , name,
    								COALESCE((SELECT SUM(points) FROM team_text_answer WHERE team_id = 1), 0) +
       								COALESCE((SELECT SUM(points) FROM team_media_answer WHERE team_id = 1), 0) + team.delta_points AS total_points
    								FROM team`
	setVideoFlagInTeam = `UPDATE team SET final_video = TRUE WHERE id = $1`
)
