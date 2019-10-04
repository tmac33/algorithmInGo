package replacespace04

func replaceSpace(str []byte, length int) {
	count := 0
	//traverse string, count the number of space
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			count++
		}
	}
	newlength := length + 2*count

	//two index,one point to (length-1),the other point to (newlength-1),traverse string once to finish replace
	for l, nl := length-1, newlength-1; l >= 0 && nl >= 0; {
		if str[l] == ' ' {
			str[nl] = '0'
			nl--
			str[nl] = '2'
			nl--
			str[nl] = '%'
			nl--
			l--
		} else {
			//fill the P1 value
			str[nl] = str[l]
			nl--
			l--
		}
	}
}
