package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step3920")
fun step3920(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
data class Pair3921<A, B>(val first: A, val second: B)

fun combine3921(left: Int, right: Int): Pair3921<Int, Int> {
    return Pair3921(left + 64, right + 19)
}
