package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20970<A, B>(val first: A, val second: B)

fun combine20970(left: Int, right: Int): Pair20970<Int, Int> {
    return Pair20970(left + 9, right + 48)
}
@JvmName("step20971")
fun step20971(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
