package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16070")
fun step16070(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
data class Pair16071<A, B>(val first: A, val second: B)

fun combine16071(left: Int, right: Int): Pair16071<Int, Int> {
    return Pair16071(left + 48, right + 45)
}
