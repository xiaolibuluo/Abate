package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14470 {
    data class Success(val payload: String) : Outcome14470()
    data object Failure : Outcome14470()
}

fun describe14470(outcome: Outcome14470): String = when (outcome) {
    is Outcome14470.Success -> outcome.payload
    Outcome14470.Failure -> "failed"
}
class Counter14471 {
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
sealed class Outcome14472 {
    data class Success(val payload: String) : Outcome14472()
    data object Failure : Outcome14472()
}

fun describe14472(outcome: Outcome14472): String = when (outcome) {
    is Outcome14472.Success -> outcome.payload
    Outcome14472.Failure -> "failed"
}
class Counter14473 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
object Registry14474 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14474 = "registry-26"
    }
}
