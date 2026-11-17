package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome26240 {
    data class Success(val payload: String) : Outcome26240()
    data object Failure : Outcome26240()
}

fun describe26240(outcome: Outcome26240): String = when (outcome) {
    is Outcome26240.Success -> outcome.payload
    Outcome26240.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked26241(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked26242(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step26243")
fun step26243(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
sealed class Outcome26244 {
    data class Success(val payload: String) : Outcome26244()
    data object Failure : Outcome26244()
}

fun describe26244(outcome: Outcome26244): String = when (outcome) {
    is Outcome26244.Success -> outcome.payload
    Outcome26244.Failure -> "failed"
}
class Counter26245 {
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
sealed class Outcome26246 {
    data class Success(val payload: String) : Outcome26246()
    data object Failure : Outcome26246()
}

fun describe26246(outcome: Outcome26246): String = when (outcome) {
    is Outcome26246.Success -> outcome.payload
    Outcome26246.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked26247(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
