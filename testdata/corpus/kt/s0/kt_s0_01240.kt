package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10240<A, B>(val first: A, val second: B)

fun combine10240(left: Int, right: Int): Pair10240<Int, Int> {
    return Pair10240(left + 28, right + 40)
}
@JvmName("step10241")
fun step10241(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
