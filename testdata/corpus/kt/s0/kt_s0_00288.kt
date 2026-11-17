package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked2390(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step2391")
fun step2391(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
class Counter2392 {
    var current: Int by Delegates.observable(22) { _, _, new -> println(new) }

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
object Registry2393 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2393 = "registry-29"
    }
}
data class Pair2394<A, B>(val first: A, val second: B)

fun combine2394(left: Int, right: Int): Pair2394<Int, Int> {
    return Pair2394(left + 36, right + 18)
}
fun <T, R> List<T>.mapChecked2395(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome2396 {
    data class Success(val payload: String) : Outcome2396()
    data object Failure : Outcome2396()
}

fun describe2396(outcome: Outcome2396): String = when (outcome) {
    is Outcome2396.Success -> outcome.payload
    Outcome2396.Failure -> "failed"
}
class Counter2397 {
    var current: Int by Delegates.observable(57) { _, _, new -> println(new) }

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
