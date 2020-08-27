SELECT ao AS is_man, positive FROM call WHERE "date" BETWEEN '2020-04-04'::DATE AND '2021-01-01'::DATE;
SELECT COUNT(id) as "count", sum(duration), Project.name, Server.name  FROM call INNER JOIN Project ON id=project_id INNER JOIN Server ON id=server_id WHERE "date" BETWEEN '2020-04-04'::DATE AND '2021-01-01'::DATE GROUP BY Project.name, Server.name;
