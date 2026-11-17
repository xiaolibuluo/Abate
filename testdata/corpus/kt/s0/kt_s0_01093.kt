package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8970 {
    data class Success(val payload: String) : Outcome8970()
    data object Failure : Outcome8970()
}

fun describe8970(outcome: Outcome8970): String = when (outcome) {
    is Outcome8970.Success -> outcome.payload
    Outcome8970.Failure -> "failed"
}
data class Pair8971<A, B>(val first: A, val second: B)

fun combine8971(left: Int, right: Int): Pair8971<Int, Int> {
    return Pair8971(left + 33, right + 16)
}
sealed class Outcome8972 {
    data class Success(val payload: String) : Outcome8972()
    data object Failure : Outcome8972()
}

fun describe8972(outcome: Outcome8972): String = when (outcome) {
    is Outcome8972.Success -> outcome.payload
    Outcome8972.Failure -> "failed"
}
data class Pair8973<A, B>(val first: A, val second: B)

fun combine8973(left: Int, right: Int): Pair8973<Int, Int> {
    return Pair8973(left + 47, right + 38)
}
fun <T, R> List<T>.mapChecked8974(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
