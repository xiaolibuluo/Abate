package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14910 {
    data class Success(val payload: String) : Outcome14910()
    data object Failure : Outcome14910()
}

fun describe14910(outcome: Outcome14910): String = when (outcome) {
    is Outcome14910.Success -> outcome.payload
    Outcome14910.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked14911(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked14912(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter14913 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
data class Pair14914<A, B>(val first: A, val second: B)

fun combine14914(left: Int, right: Int): Pair14914<Int, Int> {
    return Pair14914(left + 16, right + 59)
}
class Counter14915 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
