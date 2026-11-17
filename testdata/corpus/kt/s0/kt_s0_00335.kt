package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked2780(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step2781")
fun step2781(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
fun <T, R> List<T>.mapChecked2782(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome2783 {
    data class Success(val payload: String) : Outcome2783()
    data object Failure : Outcome2783()
}

fun describe2783(outcome: Outcome2783): String = when (outcome) {
    is Outcome2783.Success -> outcome.payload
    Outcome2783.Failure -> "failed"
}
sealed class Outcome2784 {
    data class Success(val payload: String) : Outcome2784()
    data object Failure : Outcome2784()
}

fun describe2784(outcome: Outcome2784): String = when (outcome) {
    is Outcome2784.Success -> outcome.payload
    Outcome2784.Failure -> "failed"
}
class Counter2785 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
