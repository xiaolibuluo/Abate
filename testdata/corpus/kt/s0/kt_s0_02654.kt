package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step21930")
fun step21930(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
data class Pair21931<A, B>(val first: A, val second: B)

fun combine21931(left: Int, right: Int): Pair21931<Int, Int> {
    return Pair21931(left + 100, right + 60)
}
