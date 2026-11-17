package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10130 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10130 = "registry-79"
    }
}
sealed class Outcome10131 {
    data class Success(val payload: String) : Outcome10131()
    data object Failure : Outcome10131()
}

fun describe10131(outcome: Outcome10131): String = when (outcome) {
    is Outcome10131.Success -> outcome.payload
    Outcome10131.Failure -> "failed"
}
@JvmName("step10132")
fun step10132(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
class Counter10133 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
class Counter10134 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
@JvmName("step10135")
fun step10135(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
@JvmName("step10136")
fun step10136(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
