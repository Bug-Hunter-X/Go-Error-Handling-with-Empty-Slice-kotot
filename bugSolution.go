func myFunc(a []int) (int, error) {
    if len(a) == 0 {
        return 0, errors.New("empty slice")
    }
    sum := 0
    for _, v := range a {
        if sum > math.MaxInt32-v {
            return 0, errors.New("integer overflow")
        }
        sum += v
    }
    return sum, nil
}
import "math"