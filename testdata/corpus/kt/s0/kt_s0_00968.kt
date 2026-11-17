package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7950 {
    data class Success(val payload: String) : Outcome7950()
    data object Failure : Outcome7950()
}

fun describe7950(outcome: Outcome7950): String = when (outcome) {
    is Outcome7950.Success -> outcome.payload
    Outcome7950.Failure -> "failed"
}
class Counter7951 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
sealed class Outcome7952 {
    data class Success(val payload: String) : Outcome7952()
    data object Failure : Outcome7952()
}

fun describe7952(outcome: Outcome7952): String = when (outcome) {
    is Outcome7952.Success -> outcome.payload
    Outcome7952.Failure -> "failed"
}
