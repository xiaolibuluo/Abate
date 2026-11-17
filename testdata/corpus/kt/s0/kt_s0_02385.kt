package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19610 {
    data class Success(val payload: String) : Outcome19610()
    data object Failure : Outcome19610()
}

fun describe19610(outcome: Outcome19610): String = when (outcome) {
    is Outcome19610.Success -> outcome.payload
    Outcome19610.Failure -> "failed"
}
class Counter19611 {
    var current: Int by Delegates.observable(91) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked19612(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
