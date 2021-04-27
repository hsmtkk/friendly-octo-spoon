pub fn sum_multiples(upper: i64) -> i64 {
    let mut s = 0;
    for i in 1..upper {
        if i % 3 == 0 {
            s += i;
            continue;
        }
        if i % 5 == 0 {
            s += i;
            continue;
        }
    }
    return s;
}

#[test]
fn test_sum_multiples(){
    assert_eq!(78, sum_multiples(20));
}