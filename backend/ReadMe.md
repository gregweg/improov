# Suggest Task
#curl "http://localhost:8080/api/tasks/suggest?category=Fitness"

# Complete Task
#curl -X POST http://localhost:8080/api/tasks/complete \
#  -H "Content-Type: application/json" \
#  -d '{"taskId": "1", "category": "Fitness"}'
