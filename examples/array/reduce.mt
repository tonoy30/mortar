let reduce = fn(arr, initial, f) {
    let iter = fn(arr, result) {
        if (len(arr) == 0) {
            result
        } else {
            iter(rest(arr), f(result, first(arr)));
        }
    };
    iter(arr, initial);
};
let sum = fn(arr) {
    reduce(arr, 0, fn(initial, el) { initial + el });
};
let ans = sum([1, 2, 3, 4, 5]);
print(ans);