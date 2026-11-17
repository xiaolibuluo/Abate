package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6580 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6580 = "registry-23"
    }
}
object Registry6581 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6581 = "registry-30"
    }
}
@JvmName("step6582")
fun step6582(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
sealed class Outcome6583 {
    data class Success(val payload: String) : Outcome6583()
    data object Failure : Outcome6583()
}

fun describe6583(outcome: Outcome6583): String = when (outcome) {
    is Outcome6583.Success -> outcome.payload
    Outcome6583.Failure -> "failed"
}
class Counter6584 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
