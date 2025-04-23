import "sort"

func minimumCost(cost []int) int {
    if len(cost) == 1 {
        return cost[0]
    }

    if len(cost) == 2 {
        return cost[0] + cost[1]
    }

    sort.Sort(sort.Reverse(sort.IntSlice(cost)))

    totalCost := 0
    counter := 0
    
    for i := 0; i < len(cost); i++ {
        if counter == 2 {
            counter = 0
            continue
        }

        totalCost += cost[i]
        counter++
    }

    return totalCost
}