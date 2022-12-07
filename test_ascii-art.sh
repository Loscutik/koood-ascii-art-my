cd ascii-art
while read -r str 
do
echo "======work with: ======"
echo "$str"
echo "=================="
go run . "$str" | cat -e
echo "=========end==========="
echo
done < ./tests_ascii-art.txt

echo -e 