package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18520<A, B>(val first: A, val second: B)

fun combine18520(left: Int, right: Int): Pair18520<Int, Int> {
    return Pair18520(left + 25, right + 41)
}
fun <T, R> List<T>.mapChecked18521(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
