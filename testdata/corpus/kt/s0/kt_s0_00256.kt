package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter2090 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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
sealed class Outcome2091 {
    data class Success(val payload: String) : Outcome2091()
    data object Failure : Outcome2091()
}

fun describe2091(outcome: Outcome2091): String = when (outcome) {
    is Outcome2091.Success -> outcome.payload
    Outcome2091.Failure -> "failed"
}
sealed class Outcome2092 {
    data class Success(val payload: String) : Outcome2092()
    data object Failure : Outcome2092()
}

fun describe2092(outcome: Outcome2092): String = when (outcome) {
    is Outcome2092.Success -> outcome.payload
    Outcome2092.Failure -> "failed"
}
