package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked8400(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8401 {
    data class Success(val payload: String) : Outcome8401()
    data object Failure : Outcome8401()
}

fun describe8401(outcome: Outcome8401): String = when (outcome) {
    is Outcome8401.Success -> outcome.payload
    Outcome8401.Failure -> "failed"
}
data class Pair8402<A, B>(val first: A, val second: B)

fun combine8402(left: Int, right: Int): Pair8402<Int, Int> {
    return Pair8402(left + 75, right + 55)
}
