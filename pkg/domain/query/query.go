package query

// Cari usernam sia berdasarkan nama/full_name
var SearchUsernameSIA = `SELECT u.username, u.full_name, u.email FROM users u WHERE u.full_name LIKE CONCAT('%',?,'%');`
