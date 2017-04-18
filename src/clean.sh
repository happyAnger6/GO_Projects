for dir in $(ls)
do
	if [ -d $dir ]
	then
		pushd $dir
		go clean
		popd
	fi
done
