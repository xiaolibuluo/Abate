package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12140 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12140 = "registry-73"
    }
}
fun <T, R> List<T>.mapChecked12141(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12142 {
    data class Success(val payload: String) : Outcome12142()
    data object Failure : Outcome12142()
}

fun describe12142(outcome: Outcome12142): String = when (outcome) {
    is Outcome12142.Success -> outcome.payload
    Outcome12142.Failure -> "failed"
}
data class Pair12143<A, B>(val first: A, val second: B)

fun combine12143(left: Int, right: Int): Pair12143<Int, Int> {
    return Pair12143(left + 94, right + 61)
}
data class Pair12144<A, B>(val first: A, val second: B)

fun combine12144(left: Int, right: Int): Pair12144<Int, Int> {
    return Pair12144(left + 101, right + 11)
}
sealed class Outcome12145 {
    data class Success(val payload: String) : Outcome12145()
    data object Failure : Outcome12145()
}

fun describe12145(outcome: Outcome12145): String = when (outcome) {
    is Outcome12145.Success -> outcome.payload
    Outcome12145.Failure -> "failed"
}
class Counter12146 {
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
data class Pair12147<A, B>(val first: A, val second: B)

fun combine12147(left: Int, right: Int): Pair12147<Int, Int> {
    return Pair12147(left + 25, right + 44)
}
object Registry12148 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12148 = "registry-32"
    }
}
