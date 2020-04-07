package hrk

import "testing"

/**
*You're researching friendships between groups of n new college students where each student is distinctly numbered from 1 to n. At the beginning of the semester, no student knew any other student; instead, they met and formed individual friendships as the semester went on. The friendships between students are:
*
*Bidirectional: If student a is friends with student b, then student b is also friends with student a.
*Transitive: If student a is friends with student b and student b is friends with student c, then student a is friends with student c. In other words, two students are considered to be friends even if they are only indirectly linked through a network of mutual (i.e., directly connected) friends.
*The purpose of your research is to find the maximum total value of a group's friendships, denoted by total. Each time a direct friendship forms between two students, you sum the number of friends that each of the n students has and add the sum to total.
*
*You are given q queries, where each query is in the form of an unordered list of  distinct direct friendships between  students. For each query, find the maximum value of  among all possible orderings of formed friendships and print it on a new line.
**/

func TestValueOfFriendship(t *testing.T) {
	td := []struct {
		n      int32
		fs     [][]int32
		expect int32
	}{
		{5, [][]int32{
			[]int32{1, 2},
			[]int32{3, 2},
			[]int32{4, 2},
			[]int32{4, 3},
		}, 32},
	}

	for _, v := range td {
		if get := valueOfFriendsship(v.n, v.fs); get != v.expect {
			t.Errorf("Given %v, expect %v, but %v.", v, v.expect, get)
		}
	}
}

/**
* Input Format:
* The first line contains an integer q, denoting the number of queries. The subsequent lines describe each query in the following format:
* The first line contains two space-separated integers describing the respective values of n(the number of students) and m(the number of distinct direct friendships).
* Each of the m subsequent lines contains two space-separated integers describing the respective values of x and y (where x != y) describing a friendship between student x and student y.
 */
func valueOfFriendsship(n int32, friendships [][]int32) int32 {

	students := students{}

	for i := int32(1); i <= n; i++ {
		students = append(students, &student{id: i})
	}

	total := int32(0)
	for _, m := range friendships {
		students.getStudent(m[0]).makeFriend(students.getStudent(m[1]))
		total += students.getTotal()
	}

	return 0
}

type students []*student

func (ss students) getTotal() int32 {
	total := int32(0)
	for _, s := range ss {
		total += s.fc
	}
	return total
}

func (ss students) getStudent(id int32) *student {
	for _, v := range ss {
		if v.id == id {
			return v
		}
	}
	return nil
}

type student struct {
	id int32
	fc int32   // friends count
	fl []int32 // list of friends' ids
}

func (s *student) makeFriend(target *student) {

}
