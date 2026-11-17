package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3940<A, B>(val first: A, val second: B)

fun combine3940(left: Int, right: Int): Pair3940<Int, Int> {
    return Pair3940(left + 83, right + 22)
}
@JvmName("step3941")
fun step3941(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
