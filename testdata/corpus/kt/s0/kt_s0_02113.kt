package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter17460 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
data class Pair17461<A, B>(val first: A, val second: B)

fun combine17461(left: Int, right: Int): Pair17461<Int, Int> {
    return Pair17461(left + 12, right + 42)
}
object Registry17462 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17462 = "registry-19"
    }
}
sealed class Outcome17463 {
    data class Success(val payload: String) : Outcome17463()
    data object Failure : Outcome17463()
}

fun describe17463(outcome: Outcome17463): String = when (outcome) {
    is Outcome17463.Success -> outcome.payload
    Outcome17463.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked17464(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair17465<A, B>(val first: A, val second: B)

fun combine17465(left: Int, right: Int): Pair17465<Int, Int> {
    return Pair17465(left + 40, right + 25)
}
data class Pair17466<A, B>(val first: A, val second: B)

fun combine17466(left: Int, right: Int): Pair17466<Int, Int> {
    return Pair17466(left + 47, right + 36)
}
object Registry17467 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17467 = "registry-54"
    }
}
sealed class Outcome17468 {
    data class Success(val payload: String) : Outcome17468()
    data object Failure : Outcome17468()
}

fun describe17468(outcome: Outcome17468): String = when (outcome) {
    is Outcome17468.Success -> outcome.payload
    Outcome17468.Failure -> "failed"
}
