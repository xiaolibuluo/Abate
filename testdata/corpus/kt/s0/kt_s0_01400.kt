package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome11520 {
    data class Success(val payload: String) : Outcome11520()
    data object Failure : Outcome11520()
}

fun describe11520(outcome: Outcome11520): String = when (outcome) {
    is Outcome11520.Success -> outcome.payload
    Outcome11520.Failure -> "failed"
}
sealed class Outcome11521 {
    data class Success(val payload: String) : Outcome11521()
    data object Failure : Outcome11521()
}

fun describe11521(outcome: Outcome11521): String = when (outcome) {
    is Outcome11521.Success -> outcome.payload
    Outcome11521.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11522(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter11523 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T, R> List<T>.mapChecked11524(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11525<A, B>(val first: A, val second: B)

fun combine11525(left: Int, right: Int): Pair11525<Int, Int> {
    return Pair11525(left + 78, right + 15)
}
