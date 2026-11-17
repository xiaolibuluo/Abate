package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step5530")
fun step5530(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
sealed class Outcome5531 {
    data class Success(val payload: String) : Outcome5531()
    data object Failure : Outcome5531()
}

fun describe5531(outcome: Outcome5531): String = when (outcome) {
    is Outcome5531.Success -> outcome.payload
    Outcome5531.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked5532(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter5533 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

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
