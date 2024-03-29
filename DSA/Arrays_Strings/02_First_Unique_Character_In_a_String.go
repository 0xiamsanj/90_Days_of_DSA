/* Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.



Example 1:

Input: s = "leetcode"
Output: 0
Example 2:

Input: s = "loveleetcode"
Output: 2
Example 3:

Input: s = "aabb"
Output: -1


Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters. */



func firstUniqChar(s string) int {
	//n := len(s)
    var hash[26] int
	for _,char:=range s{
        hash[char-'a']++
    }
    for i:=range s{
        if(hash[s[i] - 'a'] == 1){
            return i
        }
    }
   return -1
}
