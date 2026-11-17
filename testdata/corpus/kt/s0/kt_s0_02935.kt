package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24200 {
    data class Success(val payload: String) : Outcome24200()
    data object Failure : Outcome24200()
}

fun describe24200(outcome: Outcome24200): String = when (outcome) {
    is Outcome24200.Success -> outcome.payload
    Outcome24200.Failure -> "failed"
}
data class Pair24201<A, B>(val first: A, val second: B)

fun combine24201(left: Int, right: Int): Pair24201<Int, Int> {
    return Pair24201(left + 44, right + 63)
}
class Counter24202 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
class Counter24203 {
    var current: Int by Delegates.observable(58) { _, _, new -> println(new) }

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
sealed class Outcome24204 {
    data class Success(val payload: String) : Outcome24204()
    data object Failure : Outcome24204()
}

fun describe24204(outcome: Outcome24204): String = when (outcome) {
    is Outcome24204.Success -> outcome.payload
    Outcome24204.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked24205(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
