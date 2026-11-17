package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome1660 {
    data class Success(val payload: String) : Outcome1660()
    data object Failure : Outcome1660()
}

fun describe1660(outcome: Outcome1660): String = when (outcome) {
    is Outcome1660.Success -> outcome.payload
    Outcome1660.Failure -> "failed"
}
data class Pair1661<A, B>(val first: A, val second: B)

fun combine1661(left: Int, right: Int): Pair1661<Int, Int> {
    return Pair1661(left + 36, right + 23)
}
data class Pair1662<A, B>(val first: A, val second: B)

fun combine1662(left: Int, right: Int): Pair1662<Int, Int> {
    return Pair1662(left + 43, right + 34)
}
fun <T, R> List<T>.mapChecked1663(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked1664(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
