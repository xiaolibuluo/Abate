package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked7360(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair7361<A, B>(val first: A, val second: B)

fun combine7361(left: Int, right: Int): Pair7361<Int, Int> {
    return Pair7361(left + 74, right + 48)
}
