package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6850<A, B>(val first: A, val second: B)

fun combine6850(left: Int, right: Int): Pair6850<Int, Int> {
    return Pair6850(left + 83, right + 46)
}
@JvmName("step6851")
fun step6851(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
