package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15800(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome15801 {
    data class Success(val payload: String) : Outcome15801()
    data object Failure : Outcome15801()
}

fun describe15801(outcome: Outcome15801): String = when (outcome) {
    is Outcome15801.Success -> outcome.payload
    Outcome15801.Failure -> "failed"
}
class Counter15802 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
data class Pair15803<A, B>(val first: A, val second: B)

fun combine15803(left: Int, right: Int): Pair15803<Int, Int> {
    return Pair15803(left + 99, right + 61)
}
