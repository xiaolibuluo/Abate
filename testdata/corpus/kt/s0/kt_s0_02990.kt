package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked24590(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked24591(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair24592<A, B>(val first: A, val second: B)

fun combine24592(left: Int, right: Int): Pair24592<Int, Int> {
    return Pair24592(left + 73, right + 42)
}
sealed class Outcome24593 {
    data class Success(val payload: String) : Outcome24593()
    data object Failure : Outcome24593()
}

fun describe24593(outcome: Outcome24593): String = when (outcome) {
    is Outcome24593.Success -> outcome.payload
    Outcome24593.Failure -> "failed"
}
class Counter24594 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
