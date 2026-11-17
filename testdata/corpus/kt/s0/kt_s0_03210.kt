package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step26370")
fun step26370(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
fun <T, R> List<T>.mapChecked26371(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair26372<A, B>(val first: A, val second: B)

fun combine26372(left: Int, right: Int): Pair26372<Int, Int> {
    return Pair26372(left + 59, right + 68)
}
