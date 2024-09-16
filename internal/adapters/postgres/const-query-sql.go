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
    								COALESCE((SELECT SUM(points) FROM team_text_answer WHERE team_id = $1), 0) +
       								COALESCE((SELECT SUM(points) FROM team_media_answer WHERE team_id = $1), 0) + team.delta_points AS total_points
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
       							   COALESCE((SELECT SUM(points) FROM team_text_answer WHERE team_id = $1), 0) +
       							   COALESCE((SELECT SUM(points) FROM team_media_answer WHERE team_id = $1), 0) + delta_points AS total_points 
								   FROM team WHERE id IN (SELECT team_id FROM "user" WHERE team_id IS NOT NULL
                                							GROUP BY team_id HAVING COUNT(*) = $1);`
	getAllTeamQuery = `SELECT id , name,
    								COALESCE((SELECT SUM(points) FROM team_text_answer WHERE team_id = $1), 0) +
       								COALESCE((SELECT SUM(points) FROM team_media_answer WHERE team_id = $1), 0) + team.delta_points AS total_points
    								FROM team`
	setVideoFlagInTeam = `UPDATE team SET final_video = TRUE WHERE id = $1`
)

const (
	getTextTaskQuery       = `SELECT id, title, description, answer, points FROM text_task WHERE id NOT IN (SELECT text_task_id FROM team_text_answer WHERE team_id = $1) ORDER BY RANDOM() LIMIT 1`
	createAnswerOnTextTask = `INSERT INTO team_text_answer (team_id, text_task_id, date) VALUES ($1, $2, NOW())`
	setAnswerOnTextTask    = `UPDATE team_text_answer SET answer = $1, points = $2, status = $3, date = NOW() WHERE text_task_id = $4 AND team_id = $5`
	getLastTextTask        = `SELECT id, title, description, answer, points FROM text_task WHERE id = (SELECT text_task_id FROM team_text_answer WHERE team_id = $1 ORDER BY date DESC LIMIT 1)`
	getStatusLastTextTask  = `SELECT team_text_answer.status FROM team_text_answer WHERE team_id = $1 ORDER BY date DESC LIMIT 1`
	checkDayLastTextTask   = `SELECT COALESCE( (SELECT DATE_TRUNC('day', (SELECT team_text_answer.date FROM team_text_answer WHERE team_id = $1 ORDER BY team_text_answer.date DESC LIMIT 1)) = CURRENT_DATE), FALSE)`
)

const (
	getNewMediaTask         = `SELECT point_task.id, title, description, media_id, uuid_media, points FROM point_task JOIN public.media_obj mo ON mo.id = point_task.media_id WHERE public.point_task.id NOT IN (SELECT point_task_id FROM team_media_answer WHERE team_id = $1) ORDER BY RANDOM() LIMIT 1`
	getLastMediaTask        = `SELECT point_task.id, title, description, media_id, uuid_media, points FROM point_task JOIN public.media_obj mo ON point_task.media_id = mo.id WHERE public.point_task.id = (SELECT point_task_id FROM team_media_answer WHERE team_id = $1 ORDER BY date DESC LIMIT 1)`
	getStatusLastMediaTask  = `SELECT status FROM team_media_answer WHERE team_id = $1 ORDER BY date DESC LIMIT 1`
	createAnswerOnMediaTask = `INSERT INTO team_media_answer (team_id, point_task_id, status, date) VALUES ($1, $2, $3, NOW())`
	updateAnswerOnMediaTask = `UPDATE team_media_answer SET media_id = $1, status = $2, date = NOW() WHERE id = $3`
	getAllMediaTask         = `SELECT tma.id,
       									pt.title, 
       									pt.description, 
       									pt.media_id, 
       									mov.uuid_media,
       									tma.team_id, 
       									tma.points, 
       									tma.comment, 
       									tma.status
								FROM team_media_answer tma JOIN point_task pt ON tma.point_task_id = pt.id 
								    JOIN public.media_obj mov ON mov.id = tma.media_id
								    WHERE status != 'empty'
								    ORDER BY date DESC`

	getAnswerOnMediaTaskByStatus = `SELECT tma.id,
       									pt.title, 
       									pt.description, 
       									pt.media_id, 
       									mov.uuid_media,
       									tma.team_id, 
       									tma.points, 
       									tma.comment, 
       									tma.status
								FROM team_media_answer tma JOIN point_task pt ON tma.point_task_id = pt.id 
								    JOIN public.media_obj mov ON mov.id = tma.media_id
								    WHERE status = $1
								    ORDER BY date DESC`

	getMediaTaskById = `SELECT tma.id,
       									pt.title, 
       									pt.description, 
       									pt.media_id, 
       									mov.uuid_media,
       									tma.team_id, 
       									tma.points, 
       									tma.comment, 
       									tma.status
								FROM team_media_answer tma JOIN point_task pt ON tma.point_task_id = pt.id 
								    JOIN public.media_obj mov ON mov.id = tma.media_id 
									
								WHERE tma.id = $1`

	setPointsOnMediaTask    = `UPDATE team_media_answer SET points = $1, status = $2, comment = $3 WHERE id = $4`
	getPointsOnMediaTask    = `SELECT point_task.points FROM team_media_answer JOIN point_task ON team_media_answer.point_task_id = point_task.id WHERE team_media_answer.id = $1`
	getUpdateTimeMediaTask  = `SELECT date FROM team_media_answer WHERE id = $1`
	checkAnswerIsExistQuery = `SELECT EXISTS (SELECT id FROM team_media_answer WHERE team_id = $1 AND id = $2)`
	getAllMediaTaskByTeam   = `SELECT 
       								point_task.id, 
       								point_task.title, 
       								point_task.description,
       								team_media_answer.id, 
       								team_media_answer.points, 
       								team_media_answer.comment, 
       								team_media_answer.status
								FROM team_media_answer JOIN point_task ON team_media_answer.point_task_id = point_task.id WHERE team_id = $1`
	getMediaTaskByTeamById = `SELECT point_task.id, 
       								point_task.title, 
       								point_task.description, 
       								point_task.media_id, 
       								team_media_answer.id, 
       								team_media_answer.points, 
       								team_media_answer.comment, 
       								team_media_answer.media_id, 
       								team_media_answer.status
								FROM team_media_answer JOIN point_task ON team_media_answer.point_task_id = point_task.id WHERE team_id = $1 AND team_media_answer.id = $2`
)

const (
	createMediaObjectQuery = `INSERT INTO media_obj (uuid_media, type) VALUES ($1, $2) RETURNING id`
	getMediaObjectQuery    = `SELECT uuid_media FROM media_obj WHERE id = $1`
)

const (
	getAllUsersQuery          = `SELECT "user".id, R.type, "user".name, "group", email, telegram, vk, phone_number, T.name FROM "user" JOIN public.role R ON R.id = "user".role_id JOIN public.team T ON T.id = "user".team_id WHERE role_id IN (1,3)`
	getUserWithoutTeam        = `SELECT "user".id, R.type, "user".name, "group", email, telegram, vk, phone_number, T.name FROM "user" JOIN public.team T ON T.id = "user".team_id JOIN public.role R ON R.id = "user".role_id WHERE team_id IS NULL AND role_id IN (1, 3)`
	getUserWithCountTeamQuery = `SELECT "user".id, R.type, "user".name, "group", email, telegram, vk, phone_number, T.name FROM "user"  JOIN public.team T ON T.id = "user".team_id JOIN public.role R ON R.id = "user".role_id WHERE team_id IN (SELECT "user".team_id FROM "user" GROUP BY team_id HAVING COUNT(*) = $1) AND role_id IN (1, 3)`
	getUserByIdQuery          = `SELECT "user".id, R.type, "user".name, "group", email, telegram, vk, phone_number, T.name FROM "user"  JOIN public.team T ON T.id = "user".team_id JOIN public.role R ON R.id = "user".role_id WHERE id = $1 AND role_id IN (1, 3)`
)
