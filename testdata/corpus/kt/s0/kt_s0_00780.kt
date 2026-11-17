package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6450")
fun step6450(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
@JvmName("step6451")
fun step6451(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
data class Pair6452<A, B>(val first: A, val second: B)

fun combine6452(left: Int, right: Int): Pair6452<Int, Int> {
    return Pair6452(left + 62, right + 32)
}
@JvmName("step6453")
fun step6453(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
