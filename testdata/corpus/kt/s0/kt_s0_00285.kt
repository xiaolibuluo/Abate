package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair2360<A, B>(val first: A, val second: B)

fun combine2360(left: Int, right: Int): Pair2360<Int, Int> {
    return Pair2360(left + 66, right + 14)
}
sealed class Outcome2361 {
    data class Success(val payload: String) : Outcome2361()
    data object Failure : Outcome2361()
}

fun describe2361(outcome: Outcome2361): String = when (outcome) {
    is Outcome2361.Success -> outcome.payload
    Outcome2361.Failure -> "failed"
}
object Registry2362 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2362 = "registry-80"
    }
}
object Registry2363 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2363 = "registry-87"
    }
}
class Counter2364 {
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
fun <T, R> List<T>.mapChecked2365(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome2366 {
    data class Success(val payload: String) : Outcome2366()
    data object Failure : Outcome2366()
}

fun describe2366(outcome: Outcome2366): String = when (outcome) {
    is Outcome2366.Success -> outcome.payload
    Outcome2366.Failure -> "failed"
}
data class Pair2367<A, B>(val first: A, val second: B)

fun combine2367(left: Int, right: Int): Pair2367<Int, Int> {
    return Pair2367(left + 18, right + 30)
}
