package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked16670(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair16671<A, B>(val first: A, val second: B)

fun combine16671(left: Int, right: Int): Pair16671<Int, Int> {
    return Pair16671(left + 52, right + 38)
}
sealed class Outcome16672 {
    data class Success(val payload: String) : Outcome16672()
    data object Failure : Outcome16672()
}

fun describe16672(outcome: Outcome16672): String = when (outcome) {
    is Outcome16672.Success -> outcome.payload
    Outcome16672.Failure -> "failed"
}
