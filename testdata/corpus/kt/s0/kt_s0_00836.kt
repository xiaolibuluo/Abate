package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6870<A, B>(val first: A, val second: B)

fun combine6870(left: Int, right: Int): Pair6870<Int, Int> {
    return Pair6870(left + 12, right + 60)
}
fun <T, R> List<T>.mapChecked6871(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
