package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15720 {
    data class Success(val payload: String) : Outcome15720()
    data object Failure : Outcome15720()
}

fun describe15720(outcome: Outcome15720): String = when (outcome) {
    is Outcome15720.Success -> outcome.payload
    Outcome15720.Failure -> "failed"
}
sealed class Outcome15721 {
    data class Success(val payload: String) : Outcome15721()
    data object Failure : Outcome15721()
}

fun describe15721(outcome: Outcome15721): String = when (outcome) {
    is Outcome15721.Success -> outcome.payload
    Outcome15721.Failure -> "failed"
}
class Counter15722 {
    var current: Int by Delegates.observable(85) { _, _, new -> println(new) }

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
