package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter10490 {
    var current: Int by Delegates.observable(62) { _, _, new -> println(new) }

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
@JvmName("step10491")
fun step10491(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
@JvmName("step10492")
fun step10492(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
object Registry10493 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10493 = "registry-83"
    }
}
sealed class Outcome10494 {
    data class Success(val payload: String) : Outcome10494()
    data object Failure : Outcome10494()
}

fun describe10494(outcome: Outcome10494): String = when (outcome) {
    is Outcome10494.Success -> outcome.payload
    Outcome10494.Failure -> "failed"
}
