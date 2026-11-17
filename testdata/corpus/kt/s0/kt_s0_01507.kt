package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12470 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12470 = "registry-17"
    }
}
class Counter12471 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
sealed class Outcome12472 {
    data class Success(val payload: String) : Outcome12472()
    data object Failure : Outcome12472()
}

fun describe12472(outcome: Outcome12472): String = when (outcome) {
    is Outcome12472.Success -> outcome.payload
    Outcome12472.Failure -> "failed"
}
sealed class Outcome12473 {
    data class Success(val payload: String) : Outcome12473()
    data object Failure : Outcome12473()
}

fun describe12473(outcome: Outcome12473): String = when (outcome) {
    is Outcome12473.Success -> outcome.payload
    Outcome12473.Failure -> "failed"
}
sealed class Outcome12474 {
    data class Success(val payload: String) : Outcome12474()
    data object Failure : Outcome12474()
}

fun describe12474(outcome: Outcome12474): String = when (outcome) {
    is Outcome12474.Success -> outcome.payload
    Outcome12474.Failure -> "failed"
}
