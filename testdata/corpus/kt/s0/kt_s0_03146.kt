package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25840")
fun step25840(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
data class Pair25841<A, B>(val first: A, val second: B)

fun combine25841(left: Int, right: Int): Pair25841<Int, Int> {
    return Pair25841(left + 42, right + 52)
}
