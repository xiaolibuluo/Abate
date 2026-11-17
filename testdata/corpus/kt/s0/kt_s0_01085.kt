package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter8900 {
    var current: Int by Delegates.observable(32) { _, _, new -> println(new) }

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
sealed class Outcome8901 {
    data class Success(val payload: String) : Outcome8901()
    data object Failure : Outcome8901()
}

fun describe8901(outcome: Outcome8901): String = when (outcome) {
    is Outcome8901.Success -> outcome.payload
    Outcome8901.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked8902(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
