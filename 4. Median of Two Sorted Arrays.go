func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    num:=[] int {}
    num=append(nums1,nums2...)
    sort.Ints(num)
    var median float64
    l:= len(num)
    mid:= l/2
    if l % 2==0 {
        median = (float64(num[mid]) +float64(num[mid-1]))/2
    }else{
        median = float64(num[mid])
    }
    return median
}
