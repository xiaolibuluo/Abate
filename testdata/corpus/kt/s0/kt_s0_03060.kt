package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25100 {
    data class Success(val payload: String) : Outcome25100()
    data object Failure : Outcome25100()
}

fun describe25100(outcome: Outcome25100): String = when (outcome) {
    is Outcome25100.Success -> outcome.payload
    Outcome25100.Failure -> "failed"
}
class Counter25101 {
    var current: Int by Delegates.observable(50) { _, _, new -> println(new) }

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
sealed class Outcome25102 {
    data class Success(val payload: String) : Outcome25102()
    data object Failure : Outcome25102()
}

fun describe25102(outcome: Outcome25102): String = when (outcome) {
    is Outcome25102.Success -> outcome.payload
    Outcome25102.Failure -> "failed"
}
class Counter25103 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
