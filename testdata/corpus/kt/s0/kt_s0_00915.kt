package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter7510 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
@JvmName("step7511")
fun step7511(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
object Registry7512 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7512 = "registry-82"
    }
}
@JvmName("step7513")
fun step7513(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
sealed class Outcome7514 {
    data class Success(val payload: String) : Outcome7514()
    data object Failure : Outcome7514()
}

fun describe7514(outcome: Outcome7514): String = when (outcome) {
    is Outcome7514.Success -> outcome.payload
    Outcome7514.Failure -> "failed"
}
