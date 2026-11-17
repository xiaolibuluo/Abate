package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23690 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked23691(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter23692 {
    var current: Int by Delegates.observable(67) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked23693(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome23694 {
    data class Success(val payload: String) : Outcome23694()
    data object Failure : Outcome23694()
}

fun describe23694(outcome: Outcome23694): String = when (outcome) {
    is Outcome23694.Success -> outcome.payload
    Outcome23694.Failure -> "failed"
}
sealed class Outcome23695 {
    data class Success(val payload: String) : Outcome23695()
    data object Failure : Outcome23695()
}

fun describe23695(outcome: Outcome23695): String = when (outcome) {
    is Outcome23695.Success -> outcome.payload
    Outcome23695.Failure -> "failed"
}
