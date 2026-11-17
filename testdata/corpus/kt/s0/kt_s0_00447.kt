package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3640 {
    data class Success(val payload: String) : Outcome3640()
    data object Failure : Outcome3640()
}

fun describe3640(outcome: Outcome3640): String = when (outcome) {
    is Outcome3640.Success -> outcome.payload
    Outcome3640.Failure -> "failed"
}
sealed class Outcome3641 {
    data class Success(val payload: String) : Outcome3641()
    data object Failure : Outcome3641()
}

fun describe3641(outcome: Outcome3641): String = when (outcome) {
    is Outcome3641.Success -> outcome.payload
    Outcome3641.Failure -> "failed"
}
class Counter3642 {
    var current: Int by Delegates.observable(95) { _, _, new -> println(new) }

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
data class Pair3643<A, B>(val first: A, val second: B)

fun combine3643(left: Int, right: Int): Pair3643<Int, Int> {
    return Pair3643(left + 5, right + 28)
}
