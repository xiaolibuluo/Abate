package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter10110 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

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
object Registry10111 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10111 = "registry-60"
    }
}
sealed class Outcome10112 {
    data class Success(val payload: String) : Outcome10112()
    data object Failure : Outcome10112()
}

fun describe10112(outcome: Outcome10112): String = when (outcome) {
    is Outcome10112.Success -> outcome.payload
    Outcome10112.Failure -> "failed"
}
@JvmName("step10113")
fun step10113(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
@JvmName("step10114")
fun step10114(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
object Registry10115 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10115 = "registry-88"
    }
}
data class Pair10116<A, B>(val first: A, val second: B)

fun combine10116(left: Int, right: Int): Pair10116<Int, Int> {
    return Pair10116(left + 95, right + 15)
}
@JvmName("step10117")
fun step10117(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
