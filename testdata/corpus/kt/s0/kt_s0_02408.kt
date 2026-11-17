package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19780 {
    data class Success(val payload: String) : Outcome19780()
    data object Failure : Outcome19780()
}

fun describe19780(outcome: Outcome19780): String = when (outcome) {
    is Outcome19780.Success -> outcome.payload
    Outcome19780.Failure -> "failed"
}
@JvmName("step19781")
fun step19781(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
@JvmName("step19782")
fun step19782(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
class Counter19783 {
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
sealed class Outcome19784 {
    data class Success(val payload: String) : Outcome19784()
    data object Failure : Outcome19784()
}

fun describe19784(outcome: Outcome19784): String = when (outcome) {
    is Outcome19784.Success -> outcome.payload
    Outcome19784.Failure -> "failed"
}
object Registry19785 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19785 = "registry-49"
    }
}
