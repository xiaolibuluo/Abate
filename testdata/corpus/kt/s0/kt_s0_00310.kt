package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2590 {
    data class Success(val payload: String) : Outcome2590()
    data object Failure : Outcome2590()
}

fun describe2590(outcome: Outcome2590): String = when (outcome) {
    is Outcome2590.Success -> outcome.payload
    Outcome2590.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked2591(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter2592 {
    var current: Int by Delegates.observable(88) { _, _, new -> println(new) }

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
