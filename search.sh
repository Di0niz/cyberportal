for i in {1..255}; do 
    echo $i;
    ssh -i ~/.ssh/id_rsa student34@192.168.1.$i -J student34@95.163.251.121;
done;

