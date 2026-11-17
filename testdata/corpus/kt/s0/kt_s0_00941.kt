package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step7700")
fun step7700(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
sealed class Outcome7701 {
    data class Success(val payload: String) : Outcome7701()
    data object Failure : Outcome7701()
}

fun describe7701(outcome: Outcome7701): String = when (outcome) {
    is Outcome7701.Success -> outcome.payload
    Outcome7701.Failure -> "failed"
}
data class Pair7702<A, B>(val first: A, val second: B)

fun combine7702(left: Int, right: Int): Pair7702<Int, Int> {
    return Pair7702(left + 38, right + 53)
}
fun <T, R> List<T>.mapChecked7703(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry7704 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7704 = "registry-52"
    }
}
object Registry7705 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7705 = "registry-59"
    }
}
class Counter7706 {
    var current: Int by Delegates.observable(66) { _, _, new -> println(new) }

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
sealed class Outcome7707 {
    data class Success(val payload: String) : Outcome7707()
    data object Failure : Outcome7707()
}

fun describe7707(outcome: Outcome7707): String = when (outcome) {
    is Outcome7707.Success -> outcome.payload
    Outcome7707.Failure -> "failed"
}
