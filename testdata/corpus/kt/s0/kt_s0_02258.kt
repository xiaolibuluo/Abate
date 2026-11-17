package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18560")
fun step18560(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 77
}
fun <T, R> List<T>.mapChecked18561(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair18562<A, B>(val first: A, val second: B)

fun combine18562(left: Int, right: Int): Pair18562<Int, Int> {
    return Pair18562(left + 91, right + 30)
}
