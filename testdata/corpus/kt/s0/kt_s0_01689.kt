package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13920 {
    data class Success(val payload: String) : Outcome13920()
    data object Failure : Outcome13920()
}

fun describe13920(outcome: Outcome13920): String = when (outcome) {
    is Outcome13920.Success -> outcome.payload
    Outcome13920.Failure -> "failed"
}
@JvmName("step13921")
fun step13921(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
sealed class Outcome13922 {
    data class Success(val payload: String) : Outcome13922()
    data object Failure : Outcome13922()
}

fun describe13922(outcome: Outcome13922): String = when (outcome) {
    is Outcome13922.Success -> outcome.payload
    Outcome13922.Failure -> "failed"
}
sealed class Outcome13923 {
    data class Success(val payload: String) : Outcome13923()
    data object Failure : Outcome13923()
}

fun describe13923(outcome: Outcome13923): String = when (outcome) {
    is Outcome13923.Success -> outcome.payload
    Outcome13923.Failure -> "failed"
}
class Counter13924 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
