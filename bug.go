func myFunc(a []int) (int, error) {
    if len(a) == 0 {
        return 0, errors.New("empty slice")
    }
    sum := 0
    for _, v := range a {
        sum += v
    }
    return sum, nil
}