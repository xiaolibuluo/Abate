package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21940<A, B>(val first: A, val second: B)

fun combine21940(left: Int, right: Int): Pair21940<Int, Int> {
    return Pair21940(left + 9, right + 56)
}
@JvmName("step21941")
fun step21941(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
