package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23970<A, B>(val first: A, val second: B)

fun combine23970(left: Int, right: Int): Pair23970<Int, Int> {
    return Pair23970(left + 29, right + 13)
}
@JvmName("step23971")
fun step23971(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
@JvmName("step23972")
fun step23972(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
@JvmName("step23973")
fun step23973(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
