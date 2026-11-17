package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter660 {
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
sealed class Outcome661 {
    data class Success(val payload: String) : Outcome661()
    data object Failure : Outcome661()
}

fun describe661(outcome: Outcome661): String = when (outcome) {
    is Outcome661.Success -> outcome.payload
    Outcome661.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked662(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
