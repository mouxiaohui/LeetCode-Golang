package insert_delete_getRandom

import "math/rand"

/*
第 380 题

实现RandomizedSet 类：

RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

示例：
输入
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出
[null, true, false, true, 2, true, false, 2]

解释
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。

提示：
-231 <= val <= 231 - 1
最多调用 insert、remove 和 getRandom 函数 2 * 105 次
在调用 getRandom 方法时，数据结构中 至少存在一个 元素。
*/

type RandomizedSet struct {
	indices map[int]int
	val     []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{indices: map[int]int{}, val: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; !ok {
		this.indices[val] = len(this.val)
		this.val = append(this.val, val)
		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.indices[val]; ok {
		lastId := len(this.val) - 1

		// 把最后一个值移到需要移除的值上
		this.val[index] = this.val[lastId]
		this.indices[this.val[index]] = index
		this.val = this.val[:lastId]
		delete(this.indices, val)

		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	// 因为题目提示: 在调用 getRandom 方法时，数据结构中 至少存在一个 元素。
	return this.val[rand.Intn(len(this.val))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
