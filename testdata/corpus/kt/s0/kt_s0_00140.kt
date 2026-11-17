package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked1130(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked1131(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome1132 {
    data class Success(val payload: String) : Outcome1132()
    data object Failure : Outcome1132()
}

fun describe1132(outcome: Outcome1132): String = when (outcome) {
    is Outcome1132.Success -> outcome.payload
    Outcome1132.Failure -> "failed"
}
data class Pair1133<A, B>(val first: A, val second: B)

fun combine1133(left: Int, right: Int): Pair1133<Int, Int> {
    return Pair1133(left + 40, right + 40)
}
