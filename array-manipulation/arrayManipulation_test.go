package hrk

import "testing"

/*
 *Devendra在9号云上看到了他的教练朝他微笑。 每次教授选出Devendra单独问他一个问题，Devendra朦胧的头脑里全是他的教练和她的微笑，以至于他无法专注于其他事情。帮助他解决这个问题：
 *给你一个长度为N的列表，列表的初始值全是0。对此列表，你要进行M次查询，输出列表种最终N个值的最大值。对每次查询，给你的是3个整数——a,b和k，你要对列表中从位置a到位置b范围内的（包含a和b)的全部元素加上k。
 */

func TestArrayManipulation(t *testing.T) {
	td := []struct {
		n       int32
		queries [][]int32
		want    int64
	}{
		{5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}, 200},
	}

	for k, d := range td {
		get := arrayManipulation(d.n, d.queries)
		if get != d.want {
			t.Errorf("Testcase %v failed. Want: %v, but get: %v.", k, d.want, get)
		}
	}
}

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */
func arrayManipulation(n int32, queries [][]int32) int64 {
	dc := make(map[int32]int32)

	for _, q := range queries {
		for i := q[0]; i <= q[1]; i++ {
			dc[i] = dc[i] + int32(q[2])
		}
	}

	max := int64(0)
	for _, v := range dc {
		if int64(v) > max {
			max = int64(v)
		}
	}
	return max
}
