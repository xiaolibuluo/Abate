package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked19690(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair19691<A, B>(val first: A, val second: B)

fun combine19691(left: Int, right: Int): Pair19691<Int, Int> {
    return Pair19691(left + 98, right + 17)
}
class Counter19692 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked19693(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome19694 {
    data class Success(val payload: String) : Outcome19694()
    data object Failure : Outcome19694()
}

fun describe19694(outcome: Outcome19694): String = when (outcome) {
    is Outcome19694.Success -> outcome.payload
    Outcome19694.Failure -> "failed"
}
