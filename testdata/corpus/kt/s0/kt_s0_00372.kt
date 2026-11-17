package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3080<A, B>(val first: A, val second: B)

fun combine3080(left: Int, right: Int): Pair3080<Int, Int> {
    return Pair3080(left + 32, right + 30)
}
fun <T, R> List<T>.mapChecked3081(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked3082(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair3083<A, B>(val first: A, val second: B)

fun combine3083(left: Int, right: Int): Pair3083<Int, Int> {
    return Pair3083(left + 53, right + 63)
}
