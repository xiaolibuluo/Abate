package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13760")
fun step13760(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
class Counter13761 {
    var current: Int by Delegates.observable(52) { _, _, new -> println(new) }

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
sealed class Outcome13762 {
    data class Success(val payload: String) : Outcome13762()
    data object Failure : Outcome13762()
}

fun describe13762(outcome: Outcome13762): String = when (outcome) {
    is Outcome13762.Success -> outcome.payload
    Outcome13762.Failure -> "failed"
}
