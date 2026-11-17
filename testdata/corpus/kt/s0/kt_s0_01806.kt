package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14850 {
    data class Success(val payload: String) : Outcome14850()
    data object Failure : Outcome14850()
}

fun describe14850(outcome: Outcome14850): String = when (outcome) {
    is Outcome14850.Success -> outcome.payload
    Outcome14850.Failure -> "failed"
}
@JvmName("step14851")
fun step14851(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
object Registry14852 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14852 = "registry-21"
    }
}
sealed class Outcome14853 {
    data class Success(val payload: String) : Outcome14853()
    data object Failure : Outcome14853()
}

fun describe14853(outcome: Outcome14853): String = when (outcome) {
    is Outcome14853.Success -> outcome.payload
    Outcome14853.Failure -> "failed"
}
class Counter14854 {
    var current: Int by Delegates.observable(35) { _, _, new -> println(new) }

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
