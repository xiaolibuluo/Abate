package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21080 {
    data class Success(val payload: String) : Outcome21080()
    data object Failure : Outcome21080()
}

fun describe21080(outcome: Outcome21080): String = when (outcome) {
    is Outcome21080.Success -> outcome.payload
    Outcome21080.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked21081(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter21082 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
