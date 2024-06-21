# Хост и порт базы данных
HOST="postgres"
PORT="5432"

# Функция для ожидания доступности сервиса
echo "Waiting for PostgreSQL at $HOST:$PORT..."
while ! nc -z $HOST $PORT; do
  sleep 1
done

echo "PostgreSQL is up - executing command"
exec "$@"