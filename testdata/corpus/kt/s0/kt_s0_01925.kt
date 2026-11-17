package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15880 {
    data class Success(val payload: String) : Outcome15880()
    data object Failure : Outcome15880()
}

fun describe15880(outcome: Outcome15880): String = when (outcome) {
    is Outcome15880.Success -> outcome.payload
    Outcome15880.Failure -> "failed"
}
@JvmName("step15881")
fun step15881(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
object Registry15882 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15882 = "registry-99"
    }
}
sealed class Outcome15883 {
    data class Success(val payload: String) : Outcome15883()
    data object Failure : Outcome15883()
}

fun describe15883(outcome: Outcome15883): String = when (outcome) {
    is Outcome15883.Success -> outcome.payload
    Outcome15883.Failure -> "failed"
}
@JvmName("step15884")
fun step15884(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
@JvmName("step15885")
fun step15885(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
class Counter15886 {
    var current: Int by Delegates.observable(30) { _, _, new -> println(new) }

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
