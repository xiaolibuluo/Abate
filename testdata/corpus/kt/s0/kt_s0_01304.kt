package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10760<A, B>(val first: A, val second: B)

fun combine10760(left: Int, right: Int): Pair10760<Int, Int> {
    return Pair10760(left + 25, right + 38)
}
@JvmName("step10761")
fun step10761(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
