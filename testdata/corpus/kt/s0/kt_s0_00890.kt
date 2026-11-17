package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7290 {
    data class Success(val payload: String) : Outcome7290()
    data object Failure : Outcome7290()
}

fun describe7290(outcome: Outcome7290): String = when (outcome) {
    is Outcome7290.Success -> outcome.payload
    Outcome7290.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked7291(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair7292<A, B>(val first: A, val second: B)

fun combine7292(left: Int, right: Int): Pair7292<Int, Int> {
    return Pair7292(left + 87, right + 10)
}
class Counter7293 {
    var current: Int by Delegates.observable(94) { _, _, new -> println(new) }

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
data class Pair7294<A, B>(val first: A, val second: B)

fun combine7294(left: Int, right: Int): Pair7294<Int, Int> {
    return Pair7294(left + 101, right + 32)
}
object Registry7295 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7295 = "registry-11"
    }
}
class Counter7296 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked7297(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair7298<A, B>(val first: A, val second: B)

fun combine7298(left: Int, right: Int): Pair7298<Int, Int> {
    return Pair7298(left + 32, right + 15)
}
