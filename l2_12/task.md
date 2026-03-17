L2.12  
Утилита grep  

Пример запуска:  
echo -e "one\ntwo\nthree\ntwo" | go run main.go two  
echo -e "one\ntwo\nthree\ntwo" | go run main.go -n two    
echo -e "one\ntwo\nthree\nfour" | go run main.go -A 1 two  
echo -e "one\ntwo\nthree\nfour" | go run main.go -C 1 two  
