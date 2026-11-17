package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked6970(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked6971(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6972<A, B>(val first: A, val second: B)

fun combine6972(left: Int, right: Int): Pair6972<Int, Int> {
    return Pair6972(left + 59, right + 30)
}
sealed class Outcome6973 {
    data class Success(val payload: String) : Outcome6973()
    data object Failure : Outcome6973()
}

fun describe6973(outcome: Outcome6973): String = when (outcome) {
    is Outcome6973.Success -> outcome.payload
    Outcome6973.Failure -> "failed"
}
