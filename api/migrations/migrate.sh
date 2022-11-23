echo `date '+%y/%m/%d %H:%M:%S'`
echo "一括SQLを実行します。"
psql -h 127.0.0.1 -p 5432 -U ren -d todo_orjt -f ./migrate.sql > ./result.log
echo `date '+%y/%m/%d %H:%M:%S'`
echo "一括SQLの実行が終了しました。"