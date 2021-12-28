let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            accumulated
        } else {
            iter(rest(arr), push(accumulated, f(first(arr))));
        }
    };
    iter(arr, []);
};

let a = [1, 2, 3, 4];
let double = fn(x) { x * 2 };
let ans = map(a, double);
print(ans)