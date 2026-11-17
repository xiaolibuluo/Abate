package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter10590 {
    var current: Int by Delegates.observable(95) { _, _, new -> println(new) }

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
object Registry10591 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10591 = "registry-5"
    }
}
object Registry10592 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10592 = "registry-12"
    }
}
object Registry10593 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10593 = "registry-19"
    }
}
data class Pair10594<A, B>(val first: A, val second: B)

fun combine10594(left: Int, right: Int): Pair10594<Int, Int> {
    return Pair10594(left + 26, right + 24)
}
sealed class Outcome10595 {
    data class Success(val payload: String) : Outcome10595()
    data object Failure : Outcome10595()
}

fun describe10595(outcome: Outcome10595): String = when (outcome) {
    is Outcome10595.Success -> outcome.payload
    Outcome10595.Failure -> "failed"
}
data class Pair10596<A, B>(val first: A, val second: B)

fun combine10596(left: Int, right: Int): Pair10596<Int, Int> {
    return Pair10596(left + 40, right + 46)
}
