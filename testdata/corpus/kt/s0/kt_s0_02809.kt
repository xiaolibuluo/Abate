package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23180 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
object Registry23181 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23181 = "registry-76"
    }
}
@JvmName("step23182")
fun step23182(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
@JvmName("step23183")
fun step23183(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
sealed class Outcome23184 {
    data class Success(val payload: String) : Outcome23184()
    data object Failure : Outcome23184()
}

fun describe23184(outcome: Outcome23184): String = when (outcome) {
    is Outcome23184.Success -> outcome.payload
    Outcome23184.Failure -> "failed"
}
class Counter23185 {
    var current: Int by Delegates.observable(7) { _, _, new -> println(new) }

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
