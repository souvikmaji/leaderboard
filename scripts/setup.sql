
COPY players(id,name,playing_style_desc,score) FROM '/Users/souvik/code/leaderboard/scripts/players.csv' DELIMITER ',' CSV HEADER;
COPY teams(team_id,team_name,user_id,match_id,captain_id,v_captain_id,total_score) FROM '/Users/souvik/code/leaderboard/scripts/teams.csv' DELIMITER ',' CSV HEADER;
COPY team_players(team_team_id,player_id) FROM '/Users/souvik/code/leaderboard/scripts/team_players.csv' DELIMITER ',' CSV HEADER;
