package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4990<A, B>(val first: A, val second: B)

fun combine4990(left: Int, right: Int): Pair4990<Int, Int> {
    return Pair4990(left + 90, right + 25)
}
sealed class Outcome4991 {
    data class Success(val payload: String) : Outcome4991()
    data object Failure : Outcome4991()
}

fun describe4991(outcome: Outcome4991): String = when (outcome) {
    is Outcome4991.Success -> outcome.payload
    Outcome4991.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked4992(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
