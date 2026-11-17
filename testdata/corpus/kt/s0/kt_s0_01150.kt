package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step9460")
fun step9460(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
fun <T, R> List<T>.mapChecked9461(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome9462 {
    data class Success(val payload: String) : Outcome9462()
    data object Failure : Outcome9462()
}

fun describe9462(outcome: Outcome9462): String = when (outcome) {
    is Outcome9462.Success -> outcome.payload
    Outcome9462.Failure -> "failed"
}
class Counter9463 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
