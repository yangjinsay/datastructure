func largestNumber(nums []int) string {

    le := len(nums)
    if le == 0 {
        return ""
    } else if le == 1 {
        return strconv.Itoa(nums[0])
    }
    //sort
    for i:=0;i<le-1;i++{
        for j:=i+1;j<le;j++{
            a,b := nums[i], nums[j]
            if nums[i]%10 ==0 {
                a = nums[i]%10
            }
            if nums[j]%10 == 0{
                b = nums[j]%10
            }
            if a > b  {
               nums[i],nums[j] = nums[j],nums[i] 
            }
        }
    }
    re := ""
    for _, num := range nums {
        re = re + strconv.Itoa(num)
    }

    return re
}
