package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked6380(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6381<A, B>(val first: A, val second: B)

fun combine6381(left: Int, right: Int): Pair6381<Int, Int> {
    return Pair6381(left + 61, right + 33)
}
