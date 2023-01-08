exec 3<./tests_ascii-art.txt
while read -r -u 3 str 
do
echo "======work with: ======"
echo "$str"
echo "=================="
go run . "$str" | cat -e
echo "=========end==========="
echo
read -p "Press enter to continue"
done 

echo -e 