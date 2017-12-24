SELECT
	*
FROM
	SYSTEM_PRIVILEGE sp
WHERE
	sp.`CODE` IN (
		SELECT
			PRIVILEGECODE
		FROM
			SYSTEM_ROLE_PRIVILEGE
		WHERE
			ROLEID = (#找出roleId
				SELECT
					sr.ROLEID
				FROM
					SYSTEM_USER_ROLE sr
				WHERE
					sr.USERID = (
						SELECT
							su.id
						FROM
							SYSTEM_USER su
						WHERE
							su.LOGINNAME = 'login'
					)
			)
	);