package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25630 {
    data class Success(val payload: String) : Outcome25630()
    data object Failure : Outcome25630()
}

fun describe25630(outcome: Outcome25630): String = when (outcome) {
    is Outcome25630.Success -> outcome.payload
    Outcome25630.Failure -> "failed"
}
sealed class Outcome25631 {
    data class Success(val payload: String) : Outcome25631()
    data object Failure : Outcome25631()
}

fun describe25631(outcome: Outcome25631): String = when (outcome) {
    is Outcome25631.Success -> outcome.payload
    Outcome25631.Failure -> "failed"
}
sealed class Outcome25632 {
    data class Success(val payload: String) : Outcome25632()
    data object Failure : Outcome25632()
}

fun describe25632(outcome: Outcome25632): String = when (outcome) {
    is Outcome25632.Success -> outcome.payload
    Outcome25632.Failure -> "failed"
}
class Counter25633 {
    var current: Int by Delegates.observable(74) { _, _, new -> println(new) }

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
class Counter25634 {
    var current: Int by Delegates.observable(81) { _, _, new -> println(new) }

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
