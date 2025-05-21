docker exec rabbit-2 rabbitmqctl stop_app
docker exec rabbit-2 rabbitmqctl reset
docker exec rabbit-2 rabbitmqctl join_cluster rabbit@rabbit-1
docker exec rabbit-2 rabbitmqctl start_app

docker exec rabbit-3 rabbitmqctl stop_app
docker exec rabbit-3 rabbitmqctl reset
docker exec rabbit-3 rabbitmqctl join_cluster rabbit@rabbit-1
docker exec rabbit-3 rabbitmqctl start_app

docker exec rabbit-4 rabbitmqctl stop_app
docker exec rabbit-4 rabbitmqctl reset
docker exec rabbit-4 rabbitmqctl join_cluster rabbit@rabbit-1
docker exec rabbit-4 rabbitmqctl start_app

docker exec rabbit-5 rabbitmqctl stop_app
docker exec rabbit-5 rabbitmqctl reset
docker exec rabbit-5 rabbitmqctl join_cluster rabbit@rabbit-1
docker exec rabbit-5 rabbitmqctl start_app

docker exec rabbit-6 rabbitmqctl stop_app
docker exec rabbit-6 rabbitmqctl reset
docker exec rabbit-6 rabbitmqctl join_cluster rabbit@rabbit-1
docker exec rabbit-6 rabbitmqctl start_app