package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step20490")
fun step20490(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
data class Pair20491<A, B>(val first: A, val second: B)

fun combine20491(left: Int, right: Int): Pair20491<Int, Int> {
    return Pair20491(left + 71, right + 28)
}
