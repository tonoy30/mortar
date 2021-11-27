let fact = fn (x) {
    if (x == 1) {
        return 1
    }
    return x * fact(x-1)
}
let f = fact(5)
print(f)

let adder = fn(x) {
    fn(y) {
        x + y
    }
}
let a = adder(3)
let ans = a(5)
print(ans)