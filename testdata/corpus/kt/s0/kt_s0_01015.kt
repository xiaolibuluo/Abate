package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8330<A, B>(val first: A, val second: B)

fun combine8330(left: Int, right: Int): Pair8330<Int, Int> {
    return Pair8330(left + 67, right + 45)
}
fun <T, R> List<T>.mapChecked8331(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
