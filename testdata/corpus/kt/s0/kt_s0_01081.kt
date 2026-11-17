package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8870 {
    data class Success(val payload: String) : Outcome8870()
    data object Failure : Outcome8870()
}

fun describe8870(outcome: Outcome8870): String = when (outcome) {
    is Outcome8870.Success -> outcome.payload
    Outcome8870.Failure -> "failed"
}
@JvmName("step8871")
fun step8871(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
class Counter8872 {
    var current: Int by Delegates.observable(7) { _, _, new -> println(new) }

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
