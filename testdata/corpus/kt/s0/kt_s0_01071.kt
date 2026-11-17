package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8800<A, B>(val first: A, val second: B)

fun combine8800(left: Int, right: Int): Pair8800<Int, Int> {
    return Pair8800(left + 96, right + 69)
}
@JvmName("step8801")
fun step8801(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
