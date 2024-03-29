# course-completion-strategy

There are a total of n courses you have to take, labeled from 0 to n - 1. You are given an array - prereqs where prereqs[i] = [a, b] indicates that you must take course b first if you want to take course a. For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return the ordering of courses you have to take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

Constraints:

- 1 <= n <= 2000
- 0 <= prereqs.length <= n * (n - 1)
- prereqs[i].length == 2
- 0 <= a, b < n
- a != b
- All the pairs [a, b] are distinct.

Example 1:

Input: n = 2, prereqs = [[1,0]]

Output: [0,1]

Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].


Example 2 :

Input: n = 4, prereqs = [[1,0],[2,0],[3,1],[3,2]]

Output: [0,2,1,3]

Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finish course 0.

Two correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].


Example 3 :

Input: n = 1, prereqs = []

Output: [0]


## Solution

This can be solved using Topological sort using Kahn's algo. 
- Time complexity: O(V + E), where V is the number of vertices (courses) and E is the number of edges (dependencies/prerequisites).
- Space complexity: O(V + E), where V is the number of vertices and E is the number of edges.

<img width="1512" alt="Screenshot 2024-01-21 at 6 32 24 AM" src="https://github.com/tarunngusain08/course-completion-strategy/assets/36428256/b8b8c3e7-93f7-41c5-90bd-f50450aad849">

