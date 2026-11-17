package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17850 {
    data class Success(val payload: String) : Outcome17850()
    data object Failure : Outcome17850()
}

fun describe17850(outcome: Outcome17850): String = when (outcome) {
    is Outcome17850.Success -> outcome.payload
    Outcome17850.Failure -> "failed"
}
class Counter17851 {
    var current: Int by Delegates.observable(34) { _, _, new -> println(new) }

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
data class Pair17852<A, B>(val first: A, val second: B)

fun combine17852(left: Int, right: Int): Pair17852<Int, Int> {
    return Pair17852(left + 41, right + 21)
}
@JvmName("step17853")
fun step17853(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
sealed class Outcome17854 {
    data class Success(val payload: String) : Outcome17854()
    data object Failure : Outcome17854()
}

fun describe17854(outcome: Outcome17854): String = when (outcome) {
    is Outcome17854.Success -> outcome.payload
    Outcome17854.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked17855(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
