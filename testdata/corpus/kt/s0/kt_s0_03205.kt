package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair26340<A, B>(val first: A, val second: B)

fun combine26340(left: Int, right: Int): Pair26340<Int, Int> {
    return Pair26340(left + 6, right + 25)
}
data class Pair26341<A, B>(val first: A, val second: B)

fun combine26341(left: Int, right: Int): Pair26341<Int, Int> {
    return Pair26341(left + 13, right + 36)
}
@JvmName("step26342")
fun step26342(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
