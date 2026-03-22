L2.13  
Утилита cut  

Пример запуска:  
echo -e "a\tb\tc\td" | go run ./main.go -f 1,3  
echo "a:b:c:d" | go run ./main.go -f 2,4 -d ":"  
echo "a:b:c:d:e" | go run ./main.go -f 2-4 -d ":"  
echo "a:b:c:d:e:f" | go run ./main.go -f 1,3-4,6 -d ":"  